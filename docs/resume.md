# Resuming work on zoltar

Work was paused on 2026-11-02. This file records where everything is and what
comes next. Delete it once the work resumes.

## Where things are

- **`api-design` branch (this repo)**: the public API, locked except for the
  open questions below. Every function body panics with "not implemented", so
  the examples in `example_test.go` fail until the implementation lands.
- **`prototype` branch (this repo)**: the implementation, copied from
  enginekit's `internal/zoltar`. It uses the earlier API, so it doesn't match
  `api-design`. See "Changes from the earlier API" in `docs/design.md`.
- **enginekit `zoltar` branch**: the same implementation, still at
  `internal/zoltar` in enginekit. Kept for reference only; new work happens
  here.
- **enginekit `staticconfig` branch**: the `config/staticconfig` package,
  built on its own evaluator. It's the first consumer and gets ported onto
  zoltar later. See `docs/staticconfig-on-zoltar.md`.

## Documents

- `docs/design.md`: the decisions, and the plan for each remaining step.
- `docs/observations.md`: notes collected during step 1. Most have since
  become decisions in `docs/design.md`.
- `docs/staticconfig-on-zoltar.md`: how staticconfig uses the API.

## Toolchain

Generic methods (such as `Value.As`) need Go 1.27. Run Go commands with
`GOTOOLCHAIN=auto`, and format with that toolchain's gofmt:
`$(GOTOOLCHAIN=auto go env GOROOT)/bin/gofmt`.

## Open questions

1. Should `docs/design.md` ship in this repo, or stay a working document?
2. `ReturnsResultOf`'s doc was trimmed. These details were removed from it
   and might belong in `docs/design.md`: why only slices created for a
   variadic call are computed; that a nil error result becomes the zero value
   and a non-nil one is unknown; that uncomputed results keep their declared
   types; that a standard library `fn` computes the result of the analyzer's
   Go version; and what happens if `fn` calls `runtime.Goexit` or blocks.

## Next steps

1. Answer the open questions, then merge `api-design` into `main`.
2. Bring the prototype in line with the API, applying the renames
   (`Const` to `Known`, `Placeholder` to `Sentinel`, `ComputedBy` to
   `ReturnsResultOf`, `NoReturn` to `DoesNotReturn`, and so on) until the
   examples pass.
3. Work through steps 1b to 5 in `docs/design.md`.
4. Port enginekit's `staticconfig` onto zoltar.
