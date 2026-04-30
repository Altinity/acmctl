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

## Profiles (multi-environment)

A single `~/.acmctl.yaml` can hold multiple ACM environments —
prod, dev, stage, …— and you switch between them per command or
per shell.

```yaml
# ~/.acmctl.yaml
default_profile: prod
profiles:
  prod:
    url: https://acm.altinity.cloud/api
    token: <prod token>
  dev:
    url: https://acm.dev.altinity.cloud/api
    token: <dev token>
  stage:
    url: https://acm.staging.altinity.cloud/api
    token: <stage token>
```

Manage profiles via:

```bash
acmctl config add-profile prod  --url https://acm.altinity.cloud/api
acmctl config add-profile dev   --url https://acm.dev.altinity.cloud/api
acmctl config add-profile stage --url https://acm.staging.altinity.cloud/api

acmctl config list                # show all profiles + which is default
acmctl config get                 # show the active profile's url + token state
acmctl config use-profile dev     # change the default
acmctl config remove-profile dev
```

Select per command (highest precedence wins):

1. `--profile <name>` flag
2. `ACMCTL_PROFILE` env var
3. `default_profile:` in the config
4. fallback: the only profile, if exactly one is defined

```bash
acmctl --profile dev cluster list
ACMCTL_PROFILE=stage acmctl cluster list
```

Tokens for `login` / `oauth` / `logout` apply to the **active**
profile. `acmctl logout --all` clears tokens from every profile.

A legacy flat config (top-level `url`/`token`, no `profiles:`
section) is auto-loaded as a synthetic `default` profile and
rewritten in the new layout on first save.

## Authentication

Three ways to populate a profile's token, in order of preference:

### 1. 1Password-resolved API key (recommended for unattended use)

Mint a long-lived API key once via the ACM web UI (Account → API
tokens), store it in 1Password, then resolve it at shell start:

```bash
# in ~/.acm.env or your shell rc, with ACM_API_KEY_OP_REF set to
# the op:// reference of the item:
export ACM_API_KEY=$(op read --no-newline "$ACM_API_KEY_OP_REF")
```

`ACM_API_KEY` overrides any token saved in the active profile, so
nothing needs to be on disk. This is what `iso-acm` (the sandboxed
acm wrapper) uses.

### 2. `acmctl oauth` — browser-based PKCE (interactive)

```bash
acmctl --profile prod oauth
```

Opens your default browser to the Auth0 sign-in for the configured
Auth0 tenant, listens on a fixed loopback port (49152, fallback
49153/49154) for the callback, and saves the resulting ACM session
token to the active profile.

> **Status: blocked on an ACM backend change.** The full PKCE flow
> works against Auth0 (token exchange returns a valid id_token), but
> `POST /api/singleauth` rejects id_tokens minted by a Native CLI
> client with "Bad credentials" — the endpoint expects state from
> a flow ACM itself initiated server-side. Tracking:
> [issue #1](https://github.com/Altinity/acmctl/issues/1).
> Until that's resolved, `acmctl oauth` exits with a clear error;
> use the API-key path above instead.

Auth0 setup required (one-time, by an admin of the
`altinity.auth0.com` tenant):

- Native application, **Token Endpoint Authentication Method: None**
- Allowed Callback URLs:
  ```
  http://localhost:49152/cb
  http://localhost:49153/cb
  http://localhost:49154/cb
  ```
- Connections enabled: at least the SSO connection users sign in
  with (`google-oauth2` and/or any enterprise SSO)

### 3. `acmctl login` — email / password (legacy)

```bash
acmctl --profile prod login
```

Prompts for email + password. Most ACM tenants are SSO-only and
won't accept this; use `oauth` or the API-key path instead. Writes
the token to the active profile.

`--token <key>` non-interactively writes a known API key without
prompting.

### Sources, in precedence order

For each request, the URL and token are resolved separately:

- **URL**: `--url` flag → `ACMCTL_URL` env → active profile's `url`
- **Token**: `--token` flag → `ACM_API_KEY` env → active profile's `token`

### Logging out

```bash
acmctl logout                    # clear active profile's token
acmctl --profile dev logout      # clear a specific profile's token
acmctl logout --all              # clear every profile's token
```

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
