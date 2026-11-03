package zoltar

// WithAssumptions is an [Option] that makes a [Predictor] use the given
// assumptions about a function, instead of following it.
//
// See [Assumption].
//
// pkgPath is the import path of the package that declares it, such as "errors"
// or "gopkg.in/yaml.v3".
//
// name is the name of a function, or the dot-separated name of a type and its
// method, such as "SomeType.MyMethod". The type name does not include the
// pointer asterisk, even if the method is declared on a pointer receiver.
//
// It replaces any prior assumptions about the same function, including the
// built-in assumptions that the [Predictor] makes about standard library
// functions. Assumptions about a function that the program doesn't contain are
// ignored.
//
// With no assumptions, the predictor doesn't follow the function and assumes
// nothing about it: its result is unknown, and it may modify anything reachable
// from its arguments, and any package-level variable.
func WithAssumptions(pkgPath, name string, assumptions ...Assumption) Option {
	_ = pkgPath
	_ = name
	_ = assumptions

	panic("not implemented")
}

// Assumption is something the predictor assumes about a function instead of
// following it.
//
// Each assumption claims only what its doc says. Anything else, such as the
// result, is unknown unless another assumption says otherwise. Assumptions
// combine: for example, a function can have both the [ReadOnly] and
// [ReturnsNew] assumptions, and the predictor gives the most specific result
// they allow together.
//
// Pass the [WithAssumptions] options to [New] to add assumptions.
type Assumption interface {
	isAssumption()
}

var (
	// ReadOnly assumes that calling the function leaves no variable different:
	// nothing reachable from its arguments, and no package-level variable.
	//
	// The function may read variables, perform I/O, and lock and unlock
	// mutexes, as long as nothing is left different afterwards, so variables
	// stay known across the call.
	ReadOnly Assumption = assumption{"ReadOnly"}

	// ReturnsNew assumes that the function returns a pointer to a new variable
	// on every call, which nothing else refers to.
	//
	// The result is known not to be nil, and its [Origin] is the call. The new
	// variable's contents are unknown.
	//
	// For a function whose result is an interface, such as errors.New, the
	// result holds a new pointer whose dynamic type is unknown. It's still not
	// nil, and two calls still give different values.
	ReturnsNew Assumption = assumption{"ReturnsNew"}

	// DoesNotReturn assumes that every call panics, exits or blocks forever, so
	// code after the call can't run.
	DoesNotReturn Assumption = assumption{"DoesNotReturn"}

	// Pure assumes that the function's result depends only on its arguments,
	// including the contents of any slices passed to it. It implies [ReadOnly].
	//
	// Two calls with equal arguments give equal results, even if the results
	// are unknown. See [Value.Equal].
	Pure Assumption = assumption{"Pure"}
)

// assumption is the [Assumption] behind each of the predefined assumptions,
// such as [ReadOnly].
type assumption struct{ name string }

func (assumption) isAssumption() {}

// ReturnsResultOf assumes that the function returns the same results as
// calling fn.
//
// fn is called by the predictor when every argument is known: a constant, or a
// slice of constants created by a call to a variadic function. Otherwise, or
// if fn panics, the results are unknown. fn must return or panic, and produce
// the same results as the function in the program.
//
// fn must have the same signature as the function in the program, with two
// exceptions. Its parameters and results only need to match in kind, so a
// func(string) string can stand in for a function that takes and returns a
// named string type. For a method, the receiver is fn's first parameter.
//
// Each parameter and result must be a boolean, a string, a number, or a slice
// of those. A result may also be an error. [New] panics if fn doesn't meet
// these requirements.
//
// It says nothing about what the function writes, so it's usually combined
// with another assumption, such as [Pure]:
//
//	WithAssumptions(
//		"strings", "ToUpper",
//		ReturnsResultOf(strings.ToUpper),
//		Pure,
//	)
func ReturnsResultOf(fn any) Assumption {
	_ = fn

	panic("not implemented")
}
