# Legacy: full ACM API coverage (~269 endpoints)

Snapshot of the hand-curated CLI mapping every ACM endpoint to its own
Cobra subcommand. Preserved here for reference / cherry-picking.

## Why archived

Slim acmctl (active build) covers ~10 commands plus a generic `raw`
passthrough. For the long tail, agents call ACM directly via curl using
the `altinity-cloud` Claude Code skill, or `acmctl raw <method> <path>`.

This directory holds the previous implementation: ~6000 lines of
hand-written wrapper code mapping 269 endpoints.

## Layout

- `cmd/` — every subcommand from the full version (account, billing,
  console, storage, etc.)
- `pkg/` — typed API client methods + models for those subcommands

## Restoring a specific command

To bring back, e.g., billing support:

    cp _legacy/cmd/billing.go cmd/
    cp _legacy/pkg/api/billing.go pkg/api/
    # add billingCmd to cmd/root.go's init() if needed

You may also need supporting helpers (`flagsToParams`, `collectFieldFlags`,
etc.) — see `_legacy/cmd/helpers.go`.

## Why _underscore prefix

Go's build toolchain ignores directories starting with `_` or `.`, so this
tree never participates in `go build` while staying easy to grep.
