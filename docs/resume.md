# Resuming work on zoltar

Work was paused on 2026-11-02 and resumed the same day. The public API has
since been reviewed file by file. This file records where everything is and
what comes next. Delete it once the prototype matches the API.

## Where things are

- **`main` (this repo)**: the public API, reviewed and locked except for the
  open question below. Every function body panics with "not implemented". The
  examples in `example_test.go` use `// Prints:` instead of `// Output:`, so
  they compile but don't run.
- **`prototype` branch (this repo)**: the implementation, copied from
  enginekit's `internal/zoltar`. It uses an intermediate API. See "Changes
  from the prototype branch" in `docs/design.md` for every difference.
- **enginekit `zoltar` branch**: the same implementation, still at
  `internal/zoltar` in enginekit. Kept for reference only; new work happens
  here.
- **enginekit `staticconfig` branch**: the `config/staticconfig` package,
  built on its own evaluator. It's the first consumer and gets ported onto
  zoltar later. See `docs/staticconfig-on-zoltar.md`.

## Documents

- `docs/design.md`: the decisions, the changes from earlier APIs, and the plan
  for each remaining step.
- `docs/observations.md`: notes collected during step 1. Most have since
  become decisions in `docs/design.md`.
- `docs/staticconfig-on-zoltar.md`: how staticconfig uses the API.

## Toolchain

Generic methods (such as `Value.As`) need Go 1.27. Run Go commands with
`GOTOOLCHAIN=auto`, and format with that toolchain's gofmt:
`$(GOTOOLCHAIN=auto go env GOROOT)/bin/gofmt`.

## Open question

Should `docs/design.md` ship in this repo, or stay a working document?

## Next steps

1. Bring the prototype in line with the API, using "Changes from the
   prototype branch" in `docs/design.md`, until the examples pass. Then
   rename each `// Prints:` in `example_test.go` back to `// Output:` so that
   `go test` checks them.
2. Work through steps 1b to 5 in `docs/design.md`.
3. Port enginekit's `staticconfig` onto zoltar.
