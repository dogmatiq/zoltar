# Zoltar design

The public API is defined by the package and its doc comments. This document
records the decisions behind it, the measurements that informed them, and the
plan for implementing it. It deliberately doesn't repeat the API.

Requires Go 1.27 (generic methods).

## Decisions

These are the decisions that shape the API, with their reasons. Each is stated
in the relevant doc comments; this is the record of why.

- **Predict, then get values.** A prediction describes a call; values are got
  from it with `ValueOf`, `Returns` and `Targets`. Arguments are variadic and
  made with `Known`, `KnownAs`, `Zero`, `Unknown` or `Sentinel`. There are no
  prediction options: code that runs during initialization is predicted by
  predicting the package's `init` function, which runs every initializer in
  Go's order, so package-level variables are exact at each point.
- **Values carry their "as of" point.** `ValueOf(v, before)` takes the point
  (required, never nil); everything read from a value afterwards is as of that
  point. There is no "after".
- **Timelines.** A value covers all of its timelines at once. Methods returning
  a `Value` combine; methods returning `ok` succeed only if every timeline
  agrees. Blind spots cover every timeline: a value with no blind spots is
  known on every timeline.
- **Never wrong, only less specific.** The current implementation has known
  exceptions; see "Soundness (step 2)".
- **Eager values and control flow.** Making a prediction works out which code
  can run and where the loops are; getting a value works out it and everything
  reachable from it. Everything else is a lookup, so only those calls take a
  context. Measured on five real codebases: eager values cost 0-3.2x (at most
  +285 ms per program); eager control flow costs the same as a consumer that
  asks about every block (0.9-1.3x), and at most about 0.9 s per program more
  than one that asks almost nothing.
- **Cancellation.** A call that does work returns ctx's error if ctx is done;
  nothing it was working on is cached, and the predictor, and everything it
  returned, stays usable.
- **Work limit.** Every prediction and every value has a fixed internal work
  limit, counted in units of work, not time, so results are deterministic. A
  cached result counts the work it originally cost, so whether the limit is
  hit never depends on what was asked before. The heaviest normal question
  measured used about 0.6 million units; the limit is around 5-10 million.
  Recursion depth is also limited (32 calls). Reaching either gives a `Limit`
  blind spot. Termination isn't speed: a call graph that fans out with distinct
  arguments costs one prediction per path.
- **Nil versus ok.** Results that can be absent use `ok` where a zero value
  could be a real answer (`DynamicType`, `As`, `Len`, `Origin`, ...), and nil
  for pointers and interfaces, as `go/types` and `go/ssa` do (`Parent`,
  `CallSite`, `Origin.Prediction`). A pass's `Parent` is the prediction
  containing the loop; its `CallSite` is nil, as it isn't a call.
- **Mixing values between predictions.** Values that refer to memory belong to
  the prediction that produced them. They can be used with any prediction
  related to it by followed calls (which share variables) and passed as
  arguments to any prediction. Using one with an unrelated prediction,
  including a different pass of a loop, panics. Constants, unknowns and
  sentinels belong to none.
- **Sentinels.** A sentinel is a unique value recognized by identity, like a
  sentinel error. It has no blind spots. A sentinel of an interface type has
  no dynamic type (`DynamicType` returns ok=false), which isn't a blind spot.
- **Misuse panics; uncertainty never does.** Breaking a method's
  preconditions (a nil `before`, values from the wrong function, mismatched
  `FollowAll` targets, unsupported methods for a type) panics, and each doc
  says so.
- **Assumptions are claims.** Combining them keeps all their claims, so the
  result is the most specific they allow together. `Opaque` claims nothing, so
  combined with others it has no effect; on its own it replaces any built-in
  assumptions with none. `WithAssumptions` needs at least one assumption (it
  panics otherwise), and a later call for the same function replaces an
  earlier one, including the built-in assumptions. Assumptions about a
  function the program doesn't contain are ignored, so one list works with
  any program. `ReturnsResultOf` claims only results, so a call it doesn't
  compute is treated as for `Opaque` unless other assumptions say otherwise;
  `New` panics if its function's shape doesn't match.

## Changes from the earlier API

| Earlier | Now |
|---|---|
| `Evaluator` | `Predictor` |
| `Options.ReadOnly`, `IsReadOnlyStandardLibrary` | `WithAssumptions(pkgPath, name, ReadOnly)`; built-in assumptions from the census |
| `Invocation` | `Prediction` |
| `Options` struct | functional options; currently just `WithAssumptions` |
| `Options.Skip`, `IsStandardLibrary` | the built-in assumptions from the census; skipping is a per-function classification inside it, not an option |
| `Options.MaxSteps`, `DefaultMaxSteps` | no budget |
| `Options.MaxCallDepth`, `DefaultMaxCallDepth` | internal constant; `Limit` blind spot |
| `Options.Hooks`, `Hook`, `HookedCall`, `Custom` | `WithAssumptions`: the same mechanism, restricted to declarative assumptions and constant→constant functions; `Origin` provides the metadata hooks used to smuggle |
| `Evaluator.Invoke(fn, args, captured)` | `Predictor.Predict(fn, args)`; `Func.Predict` for function values |
| `InvokeOption`, `MayRunDuringInit` | none: predict the package's `init` function and follow calls from it, which gives the exact initialization state at each point |
| `Constant`, `OneOf`, `TypeOnly`, `Unknown`, `Pointer`, `FieldPointer`, `Composite`, `Slice`, `Func`, `Tuple`, `Placeholder` structs | the single `Value` and its methods |
| `Placeholder(t)` | `Sentinel(t)`: a unique value recognized by identity, like a sentinel error |
| `Invocation.ValueOf(v, before)` | `Prediction.ValueOf(v, before)`, with `before` now required |
| `Invocation.ValueAt` | values carry their own "as of" point |
| `Target.Func`, `Target.Args` fields | `Target.Func()`, `Target.Args()` methods |
| plain `error` strings from `Targets`/`Follow` | `BlindSpot`, which implements `error` |
| `Invocation.Iteration(index, i, over)` | `Prediction.Loops()` and `Loop.Pass(i)`: the predictor finds range loops itself |
| `Invocation.Condition` | `ValueOf(ifInstr.Cond, ifInstr).As[bool]()` |
| `AsBool`, `AsInt`, `AsString` | `Value.As[T]()` |
| `Known(t, c constant.Value)` | `Known(v)` for a predeclared type, `KnownAs(t, v)` for one of the program's types, `Zero(t)` for the zero value; callers no longer need `go/constant` |
| `ComputedBy(compute func([]constant.Value) (constant.Value, bool))` | `ReturnsResultOf(fn any)`, given the real Go function; named for the claim it makes (the program's function returns the same results as fn), not the mechanism |
| `NoReturn` | `DoesNotReturn` |
| `ConcreteType` | `DynamicType` |
| `IsKnown`, `Reason` | `HasBlindSpots`, `BlindSpots` |
| `IsNil`, `Equal` | methods on `Value` |
| `Possibilities`, `OneOf` | `HasSingleTimeline`, `Timelines` |
| `Contains`, `Evaluator.ContainsDeep` | `Value.Reachable()` |
| `Map` | methods on `Value` combine timelines themselves |
| `Slice.Elems`, `Element` fields | `Elements()`, `SliceElements`, `Element` methods |
| `Func.Func`, `Func.Captured` | `Func.Declared()`, `Func.Receiver()`, `Func.Predict()` |

## Map tracking (step 4)

Zoltar doesn't track maps today: every map's contents are unknown. Maps work
like the array behind a slice: a variable that several values can refer to,
whose contents are worked out from the writes that can happen before the
value's "as of" point.

- **Creating:** `make(map...)` and map literals create a tracked variable.
  (The ssa package turns a literal into `make` followed by one write per entry.)
- **Writing:** `m[k] = v` adds an entry. `delete(m, k)` and `clear(m)` remove
  entries.
- **Reading:** `m[k]` and `v, ok := m[k]` find the writes and deletes for the
  same key. `len(m)` is known when the entries are complete.
- **Keys:** two keys are the same if they are known to be equal (see Equal),
  and different if they are known to differ. A key that may or may not match,
  such as an unknown string, makes the read uncertain: its result includes
  every entry it might match, and the zero value.
- **Ranging:** `for k, v := range m` visits every entry in no particular
  order, so inside the loop k and v are one of the entries' keys and values.
  A map whose entries aren't complete gives unknown keys and values.
- **Escaping:** as with any tracked variable, passing the map to code that
  zoltar can't see makes its contents unknown from then on.

## Assumptions: hooks, with a verified default set

Why the standard library is skipped, measured on a small sample program with seven
string values built with `strings`, `strconv` and `fmt`:

| | Time | Exact answers |
|---|---|---|
| Skip std (default) | 129ms | 0 |
| Follow std, 100k-step budget | 929ms | 2 (`strconv.Itoa`, `strings.TrimSuffix`) |
| Follow std, no budget | 2.06s | 2, plus an infeasible timeline from `ToUpper` |

Following the standard library is not viable as a default: the useful
functions loop over their data, and loops don't iterate, so most of the work
ends in a blind spot. But every function in that sample is **pure**, and the
analyzer binary has the standard library linked in, so zoltar can *call* them
instead of analyzing them.

### Mechanism

For each function that a prediction reaches, zoltar does exactly one of:

- **Follow**: analyze the body. The default for the program's own code.
- **Compute** (`ReturnsResultOf`): the function is pure (its result depends only on its
  arguments). When every argument is known, as constants or a complete slice
  of constants whose capacity is known (see "`ReturnsResultOf` and slice
  arguments" under step 2), call the analyzer's own copy and return the
  result as a constant. Exact, at zero analysis cost. Covers the common idiom
  `fmt.Sprintf("%s-%d", name, n)`, and `strings.ToUpper`,
  `strconv.Itoa`, `path.Join`, `unicode.IsUpper`, ..., and `strings.Join`
  once slice capacity is tracked (step 2).
- **Assume**: zoltar makes assumptions about the function without following
  it, and produces the most specific value those assumptions allow:
  - it leaves no variable different (`ReadOnly`), so memory stays known
    across the call;
  - it returns a new, non-nil pointer (`ReturnsNew`: `errors.New`,
    constructor-like functions), so `IsNil` is known and `Origin` points at
    the call;
  - it never returns (`DoesNotReturn`: `os.Exit`, `log.Fatal`, `log.Panic`,
    `runtime.Goexit`), so the code after it can't run.

  Anything the assumptions don't say is a `BlindSpot` of kind `NotFollowed`.
- **Opaque**: nothing is known. The result is unknown and the arguments are
  assumed to be modified. The census assigns this with a recorded reason.
  Users can too: if following some third-party package ever proves too slow,
  `WithAssumptions(pkgPath, name, Opaque)` is the escape hatch, at the same
  per-function granularity as everything else.

This is the hook mechanism again: a table keyed by function, consulted before
following a call. What changed from `Options.Hooks` is not the mechanism but
**what an entry may return**. The old `Hook` returned any `Value`, including
an opaque `Custom`, which is where the `ValueAt` problem and the smuggled
metadata came from. An assumption is restricted to things zoltar can fully
reason about: declarative assumptions, or a pure function from constants to
constants, which can't touch the memory model.

Under that restriction users get the same mechanism as the built-in table,
because third-party packages have the same pure helpers and `Fatal`-style
functions as the standard library. For example:


```go
z := zoltar.New(prog,
	zoltar.WithAssumptions("errors", "New", zoltar.ReadOnly, zoltar.ReturnsNew),
	zoltar.WithAssumptions("bytes", "Buffer.String", zoltar.ReadOnly),
	zoltar.WithAssumptions("github.com/sirupsen/logrus", "Fatal", zoltar.DoesNotReturn),
	zoltar.WithAssumptions("example.com/app", "Slugify", zoltar.Pure),
	zoltar.WithAssumptions("strings", "Join", zoltar.ReturnsResultOf(strings.Join), zoltar.Pure),
	zoltar.WithAssumptions("strconv", "Atoi", zoltar.ReturnsResultOf(strconv.Atoi), zoltar.Pure),
	zoltar.WithAssumptions("strconv", "AppendInt", zoltar.ReturnsResultOf(strconv.AppendInt)),
)
```

Assumptions combine: `errors.New` is `ReadOnly` and `ReturnsNew`.
`Pure` implies `ReadOnly`. `ReturnsResultOf` implies nothing about writes: a call
that isn't computed may modify anything reachable from its arguments, unless
another assumption, such as `Pure`, says otherwise. Combined with others,
`Opaque` has no effect.

A non-nil error result is unknown for now. Step 2 makes it a new pointer
whose dynamic type is unknown, as for `ReturnsNew` on an interface result.

Until identity for unknowns exists (see below), `Pure` without `ReturnsResultOf`
behaves exactly like `ReadOnly`. That is correct, just not as specific as it
will become.

Zoltar's default assumptions about the standard library come from the census
below.

Caveat for the docs: a computed result comes from the analyzer's Go version,
not the analyzed program's. For these functions the two haven't differed in
years, but it should be said.

### Interim default assumptions (part of step 1)

The reshape removes the standard-library skip, so it can't ship without
default assumptions: `fmt.Println(h)` would otherwise make everything
reachable from `h` unknown. Until the census exists, the reshape carries a
small hand-written table in the same (package path, name) form:

- `ReadOnly`: exactly today's `IsReadOnlyStandardLibrary` list (`fmt.Sprintf`
  and friends, `errors.New`, `strings.*`, `strconv.*`, ...), with the
  existing special case that `fmt` functions are read-only only when no
  argument can have a `String()` or `Error()` method.
- Everything else in the standard library: `Opaque`.

Together these match today's behavior exactly, which step 1 requires. Two
things deliberately wait for step 5, because they change answers:

- auditing each `ReadOnly` entry against the stricter definition (today's
  list only promises not to modify what's reachable from the arguments;
  `ReadOnly` now also rules out writing package-level variables);
- adding `ReturnsNew` (`errors.New`, `fmt.Errorf`) and `DoesNotReturn`
  (`os.Exit`, `log.Fatal*`, `log.Panic*`, `runtime.Goexit`) entries.

Which packages count as the standard library is decided by a checked-in,
sorted list of import paths (`stdpackages.go`), generated by
`internal/genstd` and recorded with the Go version it came from. `go list
std` reports different packages for different targets (`syscall/js` only on
js/wasm, `internal/syscall/windows` only on windows, `runtime/cgo` only with
cgo), so the generator runs it for every target that `go tool dist list`
reports, with `CGO_ENABLED=0` and `CGO_ENABLED=1`, and lists the union. The
list therefore depends only on the Go version, not on the machine or the
environment that generates it. A test checks that the packages `go list std`
reports on the machine running the tests are all in the list, when the `go`
command has the same version; it doesn't check the other targets, which
would mean running `go list` about a hundred times. The list replaces the
old guess (no dot in the first path element, and a directory in the local
GOROOT), which depended on the machine. The census generates the same list.

`ReturnsResultOf` entries are not part of the interim table; they come with the
census. Replacing the interim table with the generated one is then a change
to one file, with no API change.

### Standard-library census (step 5)

Don't hand-maintain the table. Classify the **entire standard library**
ahead of time, once per Go release, and generate a Go file checked into
zoltar:

- Run an effect analysis over the standard library's SSA. Zoltar can do most
  of this itself: no stores through pointers reachable from arguments or
  globals, no calls to anything not already classified, nothing from
  `unsafe`, `reflect`, `runtime`, `syscall` or assembly, and deterministic
  (no `time`, randomness, maps' iteration order, environment) means
  **computable**; no stores but not deterministic means **read-only**;
  every path ends in `panic`, `os.Exit` or an infinite loop means **never
  returns**; a result that is always a fresh allocation means **fresh pointer**.
- Also tag functions that are **loop-free and small**, so they can be followed
  cheaply even without an adapter. `strconv.Itoa` and `strings.TrimSuffix`
  followed fine above.
- Generate the table in the same (package path, name) form as
  `WithAssumptions`, so it's readable and independent of which packages a
  given program loads.
- For each computable function, generate a `ReturnsResultOf` entry that names the
  real function, such as `ReturnsResultOf(strings.Join)`; no adapter is needed.
  Functions whose signatures don't have a shape `ReturnsResultOf` accepts
  (callbacks, io.Reader, ...) fall back to their other assumptions.
- Anything the census can't classify is skipped, as today.
- Regenerate when the module's minimum Go version changes, and record the Go
  version the table was generated from.

The census is also the honest answer to "why skip the standard library":
after it, nothing is skipped merely for being standard. Each function is
followed, computed, covered by assumptions, or opaque for a recorded reason.

## Soundness (step 2)

The promise is that a prediction is meant never to be wrong, only more or
less specific; the README and package docs list the known exceptions. A
soundness audit (322 programs, each run for real and compared with the
prediction, on darwin/arm64) found 33 wrong predictions: 32 from nine causes,
and one from `append`.
This step fixes all of them before anything else builds on the predictor.
Most fixes make an answer less specific, which is always allowed.

The audit's programs are this step's acceptance tests: each has a known
real result, and after this step no prediction may exclude it. They live in
`testdata/soundness/<category>/*.txtar`, one program per file, with the real
result recorded as data (`TestSoundness` never runs them;
`ZOLTAR_RECORD_SOUNDNESS=1 go test -run TestSoundness_record` records the
results again with `go run`). A prediction fails a case only if every
timeline is a known constant and none is the real result. Cases whose result
depends on the platform (G and H, and the other out-of-range float
conversions) record the platform, darwin/arm64, and are skipped elsewhere.

Each case that is wrong today carries a `known failure: <cause>` mark naming
a row of the table below, so the suite passes now; a marked case that stops
failing fails the suite, so each fix must remove its marks. 37 cases are
marked: A 8, B 5, B′ 3, C 3, D 3, E 3, F 3, G 2, H 2, append 4 (one from
the audit, and three of the four programs in "`append` and capacity" below),
and negative zero 1 (see "Exact run-time values" below).
The README and package docs say the promise has known exceptions until this
step removes the last mark.

| Cause | Example | Real | Predicted | Fix |
|---|---|---|---|---|
| **A.** A function literal created during package initialization is treated as part of initialization, so the writes it makes when called later are ignored | `var bump = func(){ n = 5 }`, then `bump(); return n` | 5 | 1 | Writes inside an init-created function literal count as ordinary writes that can happen after initialization. |
| **B.** Goroutines started by a function that has already returned are ignored | `NewPool()` starts a worker; `f` sends a job that sets `x = 2`, waits, returns `x` | 2 | 1 | Any goroutine whose start can happen before the read counts as a possible concurrent writer, wherever it was started. |
| **B′.** Callbacks that a skipped function runs later are treated as running at the call | `time.AfterFunc(0, func(){ ...; x = 2 })`, then `x = 3`, then wait | 2 | 3 | A callback passed to code that isn't followed may run at any later time, like a goroutine. |
| **C.** A package-level variable whose address skipped code keeps during initialization is treated as unchanged afterwards | `flag.IntVar(&x, ...)` in `init`, then `flag.Set("x", "7")` | 7 | 1 | An address passed to code that isn't followed escapes, during initialization as at any other time. |
| **D.** The result of a read-only function is assumed to share nothing with its arguments | `rename(slices.Clone(items))` changes `items[0].Name` | "x" | "a" | Unless a call is computed, its result may refer to anything reachable from its arguments. `ReadOnly` promises no writes; it does not promise an unshared result. |
| **E.** `unsafe` conversions are assumed not to alias | `p := (*int32)(unsafe.Pointer(&x)); *p = 5; return x` | 5 | 1 | Memory reached through an `unsafe` conversion, `unsafe.Slice` or similar may alias any variable whose address was taken. |
| **F.** Pointers to different zero-size variables are predicted unequal | `a, b := &E{}, &E{}` (escaping); `return a == b` | true | false | Comparisons of pointers to zero-size variables are unknown. |
| **G.** Out-of-range float-to-integer conversions get an exact answer | `uint64(-1.0)` via a variable | 0 | 18446744073709551615 | The result is unknown: Go leaves it to the platform. |
| **H.** Float arithmetic is rounded at every step, but the compiler may fuse `x*y + z` | `x*x - y` with `x = 0.1, y = x*x` | -8.3e-19 | 0 | The result is one of two values: rounded at every step, or fused (`math.FMA`). |
| **append.** Every `append` is assumed to copy | see "append and capacity" below | | | Track capacity (below). |

### `append` and capacity

Today every `append` is assumed to copy the slice into a new array. Go only
does that when the slice is full; with spare capacity it writes into the
existing array, and every slice sharing that array sees the change. So
today's predictions can be wrong, not just less specific:

| Program | Actual | Predicted today |
|---|---|---|
| `s := make([]int, 1, 4); a := append(s, 1); _ = append(s, 2); return a[1]` | `2` | `1` |
| `s := make([]int, 1, 4); r := append(s, 7); r[0] = 9; return s[0]` | `9` | `0` |
| `s := []int{0}; r := append(s, 7); r[0] = 9; return s[0]` | `0` | `0` |
| three chained appends to a nil slice, then `a := append(s, 4); _ = append(s, 5); return a[3]` | `5` | `4` |

Fix: track each slice's capacity where it can be known, and choose one of
three behaviors.

| Capacity | `append` | Typical source |
|---|---|---|
| Known, with room | Reuses the existing array: writes show in every slice sharing it | `make([]T, n, c)` with constant `c` |
| Known, full | Copies to a new array, as today | slice literals and variadic argument slices (capacity equals length) |
| Unknown | Might reuse: writes through either slice are conditional writes for the other, so answers gain timelines (the fourth example becomes `{4 \| 5}`) | the result of an `append` that grew the slice |

The capacity after growth is deliberately not computed. It depends on the Go
runtime's growth rule and element size classes, which change between Go
versions; an answer that's less specific but always right is preferred.

The conditional writes reuse the existing machinery for conditional slice
elements.

After this step, the docs' "known exception" about `append` goes, and the
promise can be stated without exceptions again.

### `ReturnsResultOf` and slice arguments

Until capacity is tracked, a computed call could be wrong whenever a slice
argument already exists elsewhere in the program. The arguments are passed
to `fn` as new slices whose capacity equals their length, and every slice
result is treated as new memory. So a function that returns part of an
argument (`bytes.TrimSpace`), writes into an argument's spare capacity
(`strconv.AppendInt`), or reads `cap()` can give a result that differs from
the program's, or hide a later write that the program would see.

So for now a call is computed only if every slice argument is nil, or is
the slice that the call itself creates for its variadic arguments
(`path.Join("a", "b")`): its capacity equals its length, and nothing else
can refer to it. A call that passes any other slice, including a spread
call `f(s...)`, isn't computed, and its results are unknown with their
declared types, exactly as if `ReturnsResultOf` didn't apply.

Once capacity is tracked, this restriction is lifted without any options on
`ReturnsResultOf`:

- pass each slice argument to `fn` with its real capacity, by building a Go
  array with the argument's whole backing array, and slicing it as the
  program does;
- after calling `fn`, check whether each slice result points into an
  argument's backing array. If it does, the result is that sub-slice of the
  argument, sharing its memory, rather than new memory;
- compare each argument's backing array before and after the call. Each
  element that differs becomes a write at that position, at the call, as for
  a call that is followed.

A slice argument whose capacity is still unknown leaves the call
uncomputed, as now.

### Exact run-time values

Known values are `go/constant` values today, which are exact, arbitrary
precision numbers. Run-time values aren't: they have a fixed size, and
floating-point values include negative zero, NaN and the infinities, which
constants can't hold. So `z := 0.0; return -z` is predicted to be `0`, when
the program really returns `-0`. This is a known exception, with one
soundness case marked `known failure: negative zero`. NaN and the infinities
already give unknown results (dividing a floating-point value by zero is
unknown), and `Known` panics for them.

Fix: represent known run-time values as real Go values of the right size,
rather than as `go/constant` values:

- integers as an `int64` or `uint64`, plus the width of the program's type;
- `float64` and `float32`;
- `complex128`;
- `bool` and `string`.

Arithmetic and conversions use Go's own operators on these values, so the
results round, wrap and overflow exactly as the program's do. A constant in
the source is converted once, to the type that SSA gives it, where SSA uses
it. This makes negative zero, NaN and the infinities exact, and lets `Known`
accept NaN and infinities.

It changes predictions (`-z` becomes `-0`, `0.0/0.0` becomes NaN rather
than unknown), so it belongs to this step, not step 1.

### Also in step 2

Precision changes that change predictions, so they can't be in step 1:

- **Recursion with known arguments is followed.** Today a function that
  calls itself is never followed a second time, so `factorial(5)` is
  unknown. Instead, a recursive call is followed while every argument is
  exactly known: that is running the program with constants, and it ends
  where the real program ends (`factorial(5)` is `120`). The depth limit for
  this case is raised well above 32, since each level is cheap; recursion
  that never ends hits it and gets a `Limit` blind spot. A recursive call with
  any argument not exactly known is still not followed, with an
  `Unsupported` blind spot, until merging exists (see "Not planned").
- **`ReturnsNew` for interface results** (`errors.New`, `fmt.Errorf`): the
  result holds a new pointer whose dynamic type is unknown, so `IsNil` is
  false and two calls differ. This makes `if err != nil` and comparisons
  with sentinel errors decidable. A call inside a loop creates a new value on
  each pass, so it shares identity's loop tests (step 3).
- **TODO: non-nil errors from `ReturnsResultOf`.** Today a non-nil error result
  of a computed call is unknown. It becomes the same "new pointer whose
  dynamic type is unknown" as `ReturnsNew` for interface results, so that
  `strconv.Atoi("abc")` gives an error that is known not to be nil.

## Identity for unknowns (step 3)

Today an unknown value has no identity, and `Equal` treats every value with
blind spots as different from everything, including itself. `Pure` makes a
stronger answer possible: the result of a pure call is fully determined by
the function and its arguments, so two calls with the same arguments return
the same value, even when neither is known.

```go
a := slugify(name) // name is unknown
b := slugify(name)
// today:  a.Equal(b) == false
// after:  a.Equal(b) == true, because both are "the result of slugify(name)"
```

This needs unknowns to carry an identity, and `Equal` to become "known to be
the same value, including two unknowns known to be the same unknown":

| Unknown | Same value when read again? |
|---|---|
| An `Unknown(t, ...)` argument passed to `Predict` | Yes: parameters never change |
| The result of a `Pure` call with equal arguments | Yes, by definition of `Pure` |
| A read from memory | Only if no write can happen in between, which zoltar already tracks |
| The result of any other call that isn't followed | No: another call may return something different |

What it enables: map lookups with unknown but identical keys
(`m[keyFor(x)] = v` then `m[keyFor(x)]` is a hit), branch conditions such as
`keyFor(x) == keyFor(x)`, and duplicate detection between values whose
contents aren't known.

The reshape designs `Value`'s internal representation with room for an
identity, and writes `Equal`'s doc in its final form. Until identity is
implemented, the "same unknown" clause simply never applies, so the doc is
true throughout.

Identity is where zoltar could become *wrong* rather than less specific, so
it needs its own soundness tests, at least:

- **Loops.** One instruction inside a loop produces a different value on
  each pass. `x := next()` in a loop must not make every `x` equal.
- **Shared predictions.** Zoltar reuses one prediction for calls with equal
  arguments, to save work. Two calls to an ordinary function can still
  return different results, so identity must come from `Pure` or from
  parameters, never from that reuse.

## Order of work

Each step is checked against the one before it.

1. **Reshape.** The new public API, hooks and budget removed, renames, the
   interim default assumptions, `Value` with room for identity, and the
   internal vocabulary brought in line. It must not change any answer: every
   existing test gives the same results before and after, and so do the
   tests of every program that uses zoltar.
1b. **Performance.** Remove the repeated work the profiling found (share
   memory-read results across nested reads; collapse unknowns while
   walking; cache walks inside cycles; don't store cache entries that can
   never match), make values and control flow eager, add `ctx` and the
   work limit, and the loop rule: when resolving a call from an earlier pass
   of a loop would depend on the call being resolved, answer "may be
   modified by a call within a loop". Acceptance: the five measured codebases
   (example, dapper, projectionkit, ferrite, configkit) finish, with
   identical answers wherever the old code finished. Eager control flow has
   been measured (affordable); it relies on the repeated-work fixes, because
   a pathological condition would otherwise cost every prediction of that
   function. A smaller sub-limit for conditions evaluated while working out
   control flow is the fallback if that cost ever matters.
2. **Soundness.** Fixes the nine causes of wrong predictions found by the
   audit, `append`, and negative zero (exact run-time values), and lifts the
   restriction on computing calls with slice arguments once capacity is
   tracked. The audit's programs are its acceptance tests. Changes answers on
   purpose.
3. **Identity for unknowns.** Changes `Equal` on purpose, with soundness
   tests for loops and shared predictions.
4. **Map tracking.** Builds on identity, so that lookups with identical
   unknown keys are hits.
5. **Standard-library census.** Generates the default assumptions and
   `ReturnsResultOf` entries, replacing the interim table.

## Not planned

- **Merging recursion with differing arguments.** Recursion is cut off at
  the internal depth limit instead. Merging would change the memory model
  (one allocation standing for many variables) and is a separate research
  item, to be considered only if real programs show the limit matters.
