# acmctl

Slim CLI for the [Altinity Cloud Manager API](https://acm.altinity.cloud/docs/),
paired with an `altinity-cloud` Claude Code skill for agent use.

## Who uses what

| Caller | Path |
|---|---|
| Human at a terminal | `acmctl` directly |
| Bash script / CI | `acmctl` — stable interface, JSON output, single binary |
| Agent in Claude Code | `altinity-cloud` skill — curl-direct with workflow recipes |

The skill and the CLI both hit the same API; pick whichever has lower
overhead for your context. Agents save tokens by skipping `--help` and
calling `curl` directly with the patterns the skill provides; humans get
discoverability and tab completion from `acmctl`.

## Design

acmctl exposes a small set of high-touch commands — login, env list/get,
cluster list/get/launch/update/delete/temp-creds — plus a generic `raw`
passthrough for the rest of the API. All output is JSON.

The full hand-curated mapping of all 269 endpoints is preserved in `_legacy/`
as reference material — Go's toolchain ignores `_`-prefixed dirs, so it stays
greppable without participating in the build. Cherry-pick from there if a
specific tier-2 endpoint becomes hot enough to warrant typing.

## The companion skill

The `altinity-cloud` skill lives in this repo at
`skills/altinity-cloud/`. Contents:

- `SKILL.md` — setup, common ops, conventions, per-tag index
- `endpoints/<tag>.md` — 20 per-tag endpoint digests (clusters,
  environments, billing, …)
- `workflows.md` — multi-step recipes (launch-and-wait,
  diagnose-slow-query, …)

Loaded lazily by Claude Code — ~30 tokens of static context until
triggered. When triggered, ~3 KB of guidance loads. Per-tag files
load only when the agent needs an endpoint outside the common-ops
list in SKILL.md.

### Installing the skill

`acmctl skills install` downloads the latest skill bundle (a
`<skill>.zip` published by this repo's
[Skill bundle](.github/workflows/skill-bundle.yml) GitHub Action)
and copies it into the per-agent skills directory:

```bash
acmctl skills install                    # default: --scope global, --agent claude
acmctl skills install --agent codex      # ~/.codex/skills/altinity-cloud/
acmctl skills install --all              # claude + codex
acmctl skills install --scope project    # ./.claude/skills/ (CWD)
acmctl skills install --dry-run          # preview, no writes
acmctl skills install --ref v1.0         # pin to a tagged release

acmctl skills update                     # refresh whatever's installed
```

`install` errors if a local file would be overwritten (rerun with
`--force`); `update` always overwrites — use it when you've made
local edits you want reverted, or just to pull the latest content.

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
