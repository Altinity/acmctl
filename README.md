# acmctl

Slim CLI for the [Altinity Cloud Manager API](https://acm.altinity.cloud/docs/).

## Design

acmctl exposes a small set of high-touch commands (login, env list/get, cluster
list/get/launch/update/delete/temp-creds) plus a generic `raw` passthrough for
the rest of the API. The full hand-curated mapping of all 269 endpoints is
preserved in `_legacy/` as reference material.

For agent use, the `altinity-cloud` Claude Code skill in this repo's
`.claude/skills/` is the recommended path — it teaches the agent to call the
API directly with curl, with workflow recipes and a per-tag endpoint digest.

## Installation

```bash
go build -o acmctl .
cp acmctl ~/bin/
```

## Authentication

Keep the token in 1Password and resolve at session start via
[`op`](https://developer.1password.com/docs/cli/) — no plaintext on
disk.

```bash
# Resolve from 1Password into the env (set ACM_API_KEY_OP_REF in
# ~/.acm.env to a reference like op://Personal/ACM/api_key first):
export ACMCTL_TOKEN=$(op read --no-newline "$ACM_API_KEY_OP_REF")

# Or pass per-invocation:
acmctl --token "$(op read --no-newline op://...)" cluster list
```

Supported sources, in precedence order:

- `--token <key>` / `--url <url>` flags
- `ACMCTL_TOKEN` / `ACMCTL_URL` env vars
- `ACM_API_KEY` env var (alias for token)

Interactive login (`acmctl login`) is also supported, but writes
the token to `~/.acmctl.yaml`. Avoid it on shared workstations.

## Commands

### Lifecycle

```bash
acmctl env list
acmctl env get 37

acmctl cluster list                       # all clusters
acmctl cluster list --env 37              # filter by environment (client-side)
acmctl cluster get 337
acmctl cluster delete 337                 # keep resources
acmctl cluster delete 337 --terminate     # tear down resources
acmctl cluster temp-creds 337             # → {"password":"..."}
```

### Launch / Update (JSON body on stdin)

```bash
cat <<EOF | acmctl cluster launch 37
{ "name": "my-cluster", "nodeType": "s1", "version": "24.3" }
EOF

cat update.json | acmctl cluster update 337
```

`launch`'s env-id arg is optional — falls back to `ACM_API_ENV_ID` env var.

### Generic passthrough

`acmctl raw <METHOD> <path>` covers everything else. Body is auto-detected:

- **No body** — most GETs and DELETEs:
  ```bash
  acmctl raw GET /cluster/337/status
  acmctl raw POST /cluster/337/backup
  acmctl raw DELETE /cluster/337/0
  ```
- **JSON body** — pipe on stdin:
  ```bash
  echo '{"name":"foo"}' | acmctl raw POST /cluster/337
  ```
- **Form-urlencoded** — `-F key=value` (repeatable, supports `@file` to load):
  ```bash
  acmctl raw POST /cluster/337/query -F query='SELECT 1' -F user=admin
  acmctl raw POST /cluster/337/kafka-configuration -F filename=k.xml -F xml=@./kafka.xml
  ```

Combining stdin JSON with `-F` flags is an error.

## Output

All commands emit JSON to stdout. Pipe to `jq` for filtering / extraction. No
table or YAML formatting — agents and scripts both want JSON.

## Restoring richer commands

The 269-endpoint hand-curated implementation lives in `_legacy/cmd/` and
`_legacy/pkg/`. Go's build toolchain ignores directories starting with `_`, so
those files don't compile against the slim build but stay easy to grep.

To bring back, e.g., billing:

```bash
cp _legacy/cmd/billing.go cmd/
cp _legacy/pkg/api/billing.go pkg/api/
# Re-add helpers if needed: collectFieldFlags, flagsToParams (in _legacy/cmd/helpers.go)
go build ./...
```

See `_legacy/README.md`.

## Shell completion

```bash
source <(acmctl completion bash)
acmctl completion zsh | source
acmctl completion fish | source
```
