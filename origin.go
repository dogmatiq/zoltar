package zoltar

import "golang.org/x/tools/go/ssa"

// An Origin is where a variable was allocated, such as by &T{}, new(), make()
// or append(). It identifies the variable.
//
// See [Value.Origin].
//
// Origins can be compared with ==. Equal origins refer to the same variable,
// so an Origin can be used as a map key, such as to avoid visiting the same
// variable twice when following a value's structure. Origins from separate
// calls to [Predictor.Predict] are never equal.
type Origin struct{}

// Prediction returns the prediction in which the variable was allocated, or
// nil for a package-level variable.
//
// For a variable allocated inside a followed call, such as to a constructor,
// it's the prediction of that call, and its [Prediction.CallSite] is the
// call.
func (o Origin) Prediction() *Prediction {
	panic("not implemented")
}

// AllocationSite returns the instruction that allocated the variable, such as
// &T{}, new() or make(), or the package-level variable itself.
//
// For a variable returned by a function assumed to be [ReturnsNew], it's the
// call to that function.
func (o Origin) AllocationSite() ssa.Value {
	panic("not implemented")
}
