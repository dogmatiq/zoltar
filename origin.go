package zoltar

import (
	"go/token"

	"golang.org/x/tools/go/ssa"
)

// Origin returns where the memory that v refers to was created, such as the
// variable that a pointer points to, or the array behind a slice.
//
// ok is false if v doesn't refer to memory, such as a constant, or if its
// timelines refer to memory created in different places.
func (v Value) Origin() (_ Origin, ok bool) {
	panic("not implemented")
}

// An Origin is where some memory was created. See [Value.Origin].
type Origin struct{}

// Prediction returns the prediction in which the memory was created, or nil
// for a package-level variable.
//
// For memory created inside a constructor such as NewGreeter(), its
// [Prediction.CallSite] is where NewGreeter was called.
func (o Origin) Prediction() *Prediction {
	panic("not implemented")
}

// Site returns the instruction that created the memory, such as &T{}, new()
// or make(), or the package-level variable.
func (o Origin) Site() ssa.Value {
	panic("not implemented")
}

// Pos returns the position of [Origin.Site], or [token.NoPos] if it has none.
func (o Origin) Pos() token.Pos {
	panic("not implemented")
}
