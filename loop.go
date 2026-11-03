package zoltar

import (
	"context"

	"golang.org/x/tools/go/ssa"
)

// A Loop is a loop that can be predicted one pass at a time. See
// [Prediction.Loops].
//
// Only range loops over a slice, an array or an integer can be predicted this
// way, because pass i of such a loop always handles element i, or the integer
// i. Other loops can't, such as a for loop with a condition, or a range loop
// over a map, a string, a channel or a function.
//
// In a prediction of a pass, the values that each pass gets afresh, such as
// the index and element of a range loop, can be known. Anything an earlier
// pass may have changed, such as a variable that each pass adds to, is
// unknown.
type Loop struct{}

// Header returns the block that each pass starts from.
func (l Loop) Header() *ssa.BasicBlock {
	panic("not implemented")
}

// Contains returns true if b is inside the loop.
func (l Loop) Contains(b *ssa.BasicBlock) bool {
	_ = b

	panic("not implemented")
}

// Passes returns the number of passes the loop makes.
//
// ok is false if that isn't known, or differs between timelines.
func (l Loop) Passes() (n int, ok bool) {
	panic("not implemented")
}

// Pass returns a prediction of pass i through the loop, counting from zero.
//
// The prediction only includes the timelines on which the loop makes more than
// i passes. If there are none, no block in it is reachable.
//
// The result's [Prediction.Parent] is the prediction that contains the loop,
// and it shares that prediction's variables. A variable allocated inside the
// loop is a different variable on each pass.
//
// It panics if i is negative.
func (l Loop) Pass(ctx context.Context, i int) (*Prediction, error) {
	_ = ctx
	_ = i

	panic("not implemented")
}
