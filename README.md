<div align="center">

# Zoltar

Zoltar is a static analyzer that predicts the values in a Go program without
running it.

[![Documentation](https://img.shields.io/badge/go.dev-documentation-007d9c?&style=for-the-badge)](https://pkg.go.dev/github.com/dogmatiq/zoltar)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/zoltar.svg?&style=for-the-badge&label=semver)](https://github.com/dogmatiq/zoltar/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/dogmatiq/zoltar/ci.yml?style=for-the-badge&branch=main)](https://github.com/dogmatiq/zoltar/actions/workflows/ci.yml)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/zoltar/main.svg?style=for-the-badge)](https://codecov.io/github/dogmatiq/zoltar)

</div>

## Overview

Zoltar reads Go code in [SSA form] and predicts the values it would produce at
run time. Given a function and what's known about its arguments, it predicts the
value of any expression at any point in that function, following calls,
branches, pointers and closures along the way.

A prediction is never wrong, only more or less specific.

## Features

- **Never guesses** — Anything Zoltar can't predict is reported as a _blind
  spot_, with the reason and the position in the code where it lost track.

- **Every possibility** — A value that depends on something unknown, such as
  which way an `if` statement goes, is predicted on each _timeline_ the program
  could take.

- **Call by call** — Each call is predicted with its own arguments, so a
  function called from two places can give two different, exact results. Calls
  through interfaces and function values are resolved to the functions they run.

- **Memory-aware** — Pointers, struct fields and slice elements are predicted as
  of a specific point in the code, counting only the writes that can happen
  before it.

- **Configurable assumptions** — Code that isn't followed, such as most of the
  standard library, is described by assumptions that you can add to or
  replace.

## Why _Zoltar_?

Zoltar is the fortune-telling machine from the film _Big_. This one predicts
your program's future, without running it, and tells you when it can't see.

## Getting started

See the [API documentation] for usage and examples.

<!-- references -->

[api documentation]: https://pkg.go.dev/github.com/dogmatiq/zoltar
[ssa form]: https://pkg.go.dev/golang.org/x/tools/go/ssa
