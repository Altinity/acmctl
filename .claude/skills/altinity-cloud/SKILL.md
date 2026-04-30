---
name: altinity-cloud
description: Use to call the Altinity Cloud Manager (ACM) REST API. Triggers when the user asks about ACM clusters, environments, ClickHouse on Altinity Cloud, Altinity support credentials, or any operation against acm.altinity.cloud or *.altinity.cloud.
---

# Altinity Cloud Manager API

Two ways to hit the API. **Curl is primary** — zero abstraction, agent reads
the spec directly. **`acmctl` is a wrapper** — handles auth/URL plumbing and
covers a few quirky endpoints (e.g. cluster delete's odd path shape).

## Setup (do once per session)

```bash
# Token + URL come from env vars set by the calling environment:
#   - ACM_API_KEY:  set by `iso` from [users.<name>.auth] (1Password),
#                   and by `acm-shell` after acm_resolve_acm_api_key
#   - ACMCTL_TOKEN: alias accepted by acmctl
#   - ACM_API_URL:  set by acm-shell — note this is the HOST ROOT
#                   (e.g. https://acm.altinity.cloud), without /api,
#                   because curl-based callers append paths themselves
export ACM_TOKEN="${ACM_API_KEY:-${ACMCTL_TOKEN:-}}"
ACM_URL="${ACM_API_URL:-${ACMCTL_URL:-https://acm.altinity.cloud}}"

# Normalize: API endpoints live under /api/. Append if the caller's
# URL doesn't already include it. Otherwise GETs hit the SPA HTML root
# and jq chokes on `<!doctype html>`.
case "$ACM_URL" in
    */api|*/api/) ;;
    *) ACM_URL="${ACM_URL%/}/api" ;;
esac
export ACM_URL

[ -n "$ACM_TOKEN" ] || { echo "no ACM token — set ACM_API_KEY or run via iso-acm/acm" >&2; }
```

## Curl helper

```bash
acm() {
  local method=$1 path=$2; shift 2
  curl -sfX "$method" \
    -H "X-Auth-Token: $ACM_TOKEN" \
    "$@" \
    "${ACM_URL%/}$path" | jq '.data // .'
}
```

Unwraps the API's `{data, error}` envelope. To see the raw envelope (e.g. for
debugging), drop the `| jq …` and use `curl -s ...` directly.

## acmctl wrapper

```bash
acmctl env list
acmctl env get 37
acmctl cluster list [--env 37]
acmctl cluster get 337
acmctl cluster launch [37] < cluster.json     # env-id arg optional, falls back to ACM_API_ENV_ID
acmctl cluster update 337 < update.json
acmctl cluster delete 337 [--terminate]
acmctl cluster temp-creds 337                 # → {"password":"..."}

# Catch-all for everything else
acmctl raw GET    /cluster/337/status
acmctl raw POST   /cluster/337/backup
acmctl raw POST   /cluster/337/query -F query='SELECT 1' -F user=admin
acmctl raw POST   /cluster/337/kafka-configuration -F xml=@./kafka.xml
acmctl raw POST   /environment/37/clusters/launch < cluster.json
acmctl raw DELETE /cluster/337/0
```

Body shape for `acmctl raw` is auto-detected:
- **stdin JSON** → sent as `application/json`
- **`-F key=value`** flags → sent as `application/x-www-form-urlencoded`
- **neither** → no body (most GETs/DELETEs)

Combining the two is an error.

## Common operations

```bash
# Listing
acm GET /clusters
acm GET /environments
acm GET /cluster/337/nodes
acm GET /cluster/337/users          # ClickHouse users on a cluster

# Inspecting
acm GET /cluster/337
acm GET /cluster/337/status
acm GET /cluster/337/logs
acm GET /cluster/337/workload-queries
acm GET /cluster/337/errors

# Lifecycle
acm PUT /cluster/337/stop
acm PUT /cluster/337/resume
acm PUT /cluster/337/restart
acm POST /cluster/337/backup

# Querying ClickHouse (form-urlencoded)
acm POST /cluster/337/query \
  --data-urlencode "query=SELECT count() FROM system.tables" \
  --data-urlencode "user=admin"

# Or with JSON body — also works
acm POST /cluster/337/query \
  -H "Content-Type: application/json" \
  --data-binary '{"query":"SELECT 1","user":"admin"}'

# Altinity support workflows
acm GET /cluster/337/support/credentials  # shape varies — see Conventions below
acm POST /cluster/337/support/refresh
acm POST /cluster/337/push                # publish pending config
```

## Conventions you must know

- **Body encoding**: ACM accepts both `application/x-www-form-urlencoded` and
  `application/json`. Form encoding matches the OpenAPI spec; JSON works too
  (the existing `support-team/acm-shell/scripts/api.sh` uses JSON). Pick
  whichever is easier for your input shape.

- **Path params**: spec uses `{id}` / `{cluster}` / `{environment}`
  interchangeably; they're all integer IDs.

- **Delete cluster**: the path includes a trailing terminate flag.
  `DELETE /cluster/337/0` keeps resources, `/cluster/337/1` terminates.
  (`acmctl cluster delete --terminate` handles this.)

- **`temp-creds` response shape is unstable.** The OpenAPI spec leaves it
  undefined (`"schema": []`), and the format has changed at least once:
  - Older API: `{"data": {"login": "...", "password": "..."}}`
  - Some versions: `{"data": "<password-string>"}` (just the password)
  - Current: assume it could be either — handle both shapes in code, e.g.
    ```bash
    resp=$(acm GET /cluster/337/support/credentials)
    user=$(echo "$resp" | jq -r 'if type == "object" then .login // empty else empty end')
    pass=$(echo "$resp" | jq -r 'if type == "object" then .password else . end')
    [ -z "$user" ] && user="${EXPERT_CH_USER:-}"   # fall back to session user
    ```
  `acmctl cluster temp-creds <id>` and `acmctl raw GET /cluster/<id>/support/credentials`
  both pass through `.data` verbatim — no reshaping.

- **Auth failure**: 401 → token expired or wrong. Rerun `acmctl login` or
  reset `ACMCTL_TOKEN`. Don't try to refresh tokens programmatically.

- **Pagination**: a few endpoints support `page` and `limit` query params
  (audit logs, account log, console logs, console tasks). Most don't.

- **Empty response schema**: many endpoints declare `"schema": []` — they
  return arbitrary JSON. Trust what you get back, not the spec.

## When this skill isn't enough

- For an endpoint not in "Common operations" above, see
  `resources/endpoints.md` — a per-tag index of all 270 endpoints. Read the
  index to find the right tag, then read only `resources/endpoints/<tag>.md`.

- For multi-step workflows (launch-and-wait, diagnose-slow-query,
  restore-from-backup), see `resources/workflows.md`.

- For full parameter schemas (which fields a POST accepts), the OpenAPI
  spec is at `/Users/Workspaces/altinity/acmctl/reference_auth.json`.
  Don't read the whole file — grep for the path you care about.

- For richer typed wrappers (e.g. billing, console, storage), the original
  269-endpoint implementation is preserved at
  `/Users/Workspaces/altinity/acmctl/_legacy/`. Cherry-pick if needed.
