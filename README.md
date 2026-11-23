<div align="center">

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://dogmatiq.io/res/img/dynamic/repo/zoltar/banner/dark.svg">
  <img alt="Dogmatiq zoltar" src="https://dogmatiq.io/res/img/dynamic/repo/zoltar/banner/light.svg" height="44">
</picture>

Zoltar is a static analyzer that predicts the values in a Go program without
running it.

[![API reference](https://dogmatiq.io/res/img/dynamic/repo/zoltar/badge/docs.svg)](https://pkg.go.dev/github.com/dogmatiq/zoltar)
[![Latest version](https://dogmatiq.io/res/img/dynamic/repo/zoltar/badge/semver.svg)](https://github.com/dogmatiq/zoltar/releases)
[![Build status](https://dogmatiq.io/res/img/dynamic/repo/zoltar/badge/build.svg)](https://github.com/dogmatiq/zoltar/actions/workflows/ci.yml)
[![Code coverage](https://dogmatiq.io/res/img/dynamic/repo/zoltar/badge/coverage.svg)](https://codecov.io/github/dogmatiq/zoltar)

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

- **Memory-aware** — Pointers, struct fields and slice elements are predicted at
  a specific point in the code, counting only the writes that can happen before
  it.

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
