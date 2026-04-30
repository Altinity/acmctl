---
name: altinity-cloud
description: Use to call the Altinity Cloud Manager (ACM) REST API. Triggers when the user asks about ACM clusters, environments, ClickHouse on Altinity Cloud, Altinity support credentials, or any operation against acm.altinity.cloud or *.altinity.cloud.
---

# Altinity Cloud Manager API

Use `acmctl` for all API calls. It handles auth, URL/`/api` path
normalization, and the `{data, error}` response envelope. The high-touch
endpoints have dedicated subcommands; `acmctl raw <METHOD> /path` is the
fallback for everything else.

## Setup

`acmctl` reads `ACM_API_KEY` from env (set by iso from
`[users.<name>.auth]` via 1Password, or by acm-shell after
`acm_resolve_acm_api_key`). Nothing to do at session start — just
verify it's set:

```bash
[ -n "$ACM_API_KEY" ] || echo "no ACM token — run via iso-acm/acm or set ACM_API_KEY manually" >&2
```

URL defaults to `https://acm.altinity.cloud/api/`. Override with `--url` or
`ACMCTL_URL` for staging/dev.

## Common operations

```bash
# Environments
acmctl env list
acmctl env get 37

# Clusters
acmctl cluster list                          # all
acmctl cluster list --env 37                 # filter by env
acmctl cluster get 337
acmctl cluster launch [37] < cluster.json    # env-id falls back to ACM_API_ENV_ID
acmctl cluster update 337 < update.json
acmctl cluster delete 337 [--terminate]
acmctl cluster temp-creds 337                # mint Altinity-support creds

# Catch-all for everything else (270+ endpoints)
acmctl raw GET    /cluster/337/status
acmctl raw POST   /cluster/337/backup
acmctl raw POST   /cluster/337/query -F query='SELECT 1' -F user=admin
acmctl raw POST   /cluster/337/kafka-configuration -F xml=@./kafka.xml
acmctl raw DELETE /cluster/337/0             # last segment: 0=keep resources, 1=terminate
```

`acmctl raw` body shape is auto-detected:
- **stdin JSON** → `application/json`
- **`-F key=value`** flags → `application/x-www-form-urlencoded`
- **neither** → no body

Combining stdin JSON with `-F` is an error.

## Conventions you must know

- **Path params** are integer IDs. The spec uses `{id}` / `{cluster}` /
  `{environment}` interchangeably.

- **`temp-creds` response shape is unstable.** The OpenAPI spec leaves it
  undefined, and the format has varied: sometimes `{login, password}`,
  sometimes a bare password string. Handle both:

  ```bash
  resp=$(acmctl cluster temp-creds 337)
  user=$(jq -r 'if type == "object" then .login // empty else empty end' <<<"$resp")
  pass=$(jq -r 'if type == "object" then .password else . end' <<<"$resp")
  [ -z "$user" ] && user="${EXPERT_CH_USER:-}"   # fall back to session user
  ```

- **Auth failure**: 401 → token expired or wrong. Don't try to refresh
  tokens programmatically; tell the user.

- **Pagination**: a handful of endpoints accept `page` / `limit` query
  params (audit logs, account log, console logs, console tasks). Most
  don't — assume not unless the per-tag digest says otherwise.

- **Empty response schema**: many endpoints declare `"schema": []` —
  they return arbitrary JSON. Trust what comes back, not the spec.

## When this skill isn't enough

For an endpoint not listed above, read the per-tag digest. **Open only
the tag you need** — each file is a focused 1–4 KB excerpt.

| Tag | Ops | File |
|---|---|---|
| Accounts | 15 | `endpoints/accounts.md` |
| AltinityNotes | 4 | `endpoints/altinity-notes.md` |
| AuditReports | 3 | `endpoints/audit-reports.md` |
| Auth | 17 | `endpoints/auth.md` |
| Billing | 23 | `endpoints/billing.md` |
| CHAPIEndpoints | 3 | `endpoints/ch-api-endpoints.md` |
| Cloud | 3 | `endpoints/cloud.md` |
| ClusterSettings | 9 | `endpoints/cluster-settings.md` |
| Clusters | 87 | `endpoints/clusters.md` |
| Console | 10 | `endpoints/console.md` |
| DatabaseProfileSettings | 5 | `endpoints/database-profile-settings.md` |
| DatabaseProfiles | 4 | `endpoints/database-profiles.md` |
| DatabaseUsers | 6 | `endpoints/database-users.md` |
| Environments | 40 | `endpoints/environments.md` |
| NodeTypes | 4 | `endpoints/node-types.md` |
| Notifications | 7 | `endpoints/notifications.md` |
| Organizations | 6 | `endpoints/organizations.md` |
| Storage | 10 | `endpoints/storage.md` |
| Utilities | 4 | `endpoints/utilities.md` |
| Zookeepers | 8 | `endpoints/zookeepers.md` |

For multi-step workflows (launch-and-wait, diagnose-slow-query,
restore-from-backup), see `workflows.md`.

For full parameter schemas, the OpenAPI spec is at
`../reference_auth.json` (relative to this skill, i.e. acmctl repo
root). **Don't read the whole file** — grep for the path you care about.
