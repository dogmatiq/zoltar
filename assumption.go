package zoltar

// WithAssumptions is an [Option] that configures a [Predictor] to make the
// given assumptions about a specific function instead of following it.
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
// It replaces any prior assumptions about the same function. This includes
// replacing the built-in assumptions that the [Predictor] makes about standard
// library functions.
func WithAssumptions(pkgPath, name string, assumptions ...Assumption) Option {
	if len(assumptions) == 0 {
		panic("at least one assumption must be provided")
	}

	panic("not implemented")
}

// Assumption is something the predictor assumes about a function instead of
// following it.
//
// Assumptions are combinatorial. For example, a function can have both the
// [ReadOnly] and [ReturnsNew] assumptions, and the predictor gives the most
// specific result they allow together.
//
// Pass the [WithAssumptions] options to [New] to add assumptions.
type Assumption interface {
	isAssumption()
}

var (
	// ReadOnly assumes that calling the function leaves no variable different:
	// nothing reachable from its arguments, and no package-level variable.
	//
	// It may read memory, perform I/O, and lock and unlock mutexes, as long as
	// nothing is left different afterwards. Memory stays known across the call.
	// On its own, it leaves the result unknown, and the result may share memory
	// with the arguments, as the result of slices.Clone shares its elements.
	ReadOnly Assumption = assumption{"ReadOnly"}

	// ReturnsNew assumes that the function returns a new, non-nil pointer on
	// every call, which nothing else refers to. [Value.IsNil] is known, and
	// [Value.Origin] points at the call. The contents are unknown.
	//
	// For a function whose result is an interface, such as errors.New, the
	// result holds a new pointer whose dynamic type is unknown: it is still
	// not nil, and two calls still give different values.
	ReturnsNew Assumption = assumption{"ReturnsNew"}

	// DoesNotReturn assumes that every call panics, exits or blocks forever, so
	// code after the call can't run.
	DoesNotReturn Assumption = assumption{"DoesNotReturn"}

	// Pure assumes that the function's result depends only on its arguments,
	// including the contents of any slices passed to it. It implies [ReadOnly].
	Pure Assumption = assumption{"Pure"}

	// Opaque assumes nothing: the result is unknown, and anything reachable
	// from the arguments, and any package-level variable, may be modified.
	// Combined with other assumptions, it has no effect.
	Opaque Assumption = assumption{"Opaque"}
)

// assumption is the [Assumption] behind each of the predefined assumptions,
// such as [ReadOnly].
type assumption struct{ name string }

func (assumption) isAssumption() {}

// ReturnsResultOf assumes that the function returns the same results as fn,
// such as strings.ToUpper for the function ToUpper in package "strings".
//
// fn is called by the predictor when every argument is known: a constant, or a
// slice of constants created by a call to a variadic function. Otherwise, or
// if fn panics, the results are unknown. fn must return or panic, and produce
// the same results as the function in the program.
//
// fn's signature must have the same shape as the function's: the same number
// of parameters and results, the same variadic-ness, and each parameter and
// result of the same kind, as in package reflect. The kinds allowed are bool,
// string, the numeric kinds, slices of those, and for a result, error. For a
// method, the receiver is the first parameter. [New] panics if fn isn't such a
// function.
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
	panic("not implemented")
}
