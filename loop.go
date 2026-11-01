package zoltar

import (
	"context"

	"golang.org/x/tools/go/ssa"
)

// Loops returns the loops in p's function that can be predicted one pass at a
// time: range loops over a slice, an array or an integer.
func (p *Prediction) Loops() []Loop {
	panic("not implemented")
}

// A Loop is a loop that can be predicted one pass at a time. See
// [Prediction.Loops].
//
// Within a pass, values that change from one pass to the next, such as the
// current element, can be known. Anything worked out by an earlier pass is
// unknown.
type Loop struct{}

// Header returns the block that each pass starts from.
func (l Loop) Header() *ssa.BasicBlock {
	panic("not implemented")
}

// Contains returns true if b is inside the loop.
func (l Loop) Contains(b *ssa.BasicBlock) bool {
	panic("not implemented")
}

// Passes returns the number of passes the loop makes, if it's known.
func (l Loop) Passes() (n int, ok bool) {
	panic("not implemented")
}

// Pass returns a prediction of the i'th pass through the loop, counting from
// zero.
//
// i may be more than the number of passes, even if it isn't known: the
// timelines on which the loop makes fewer passes are left out. It panics if i
// is negative.
func (l Loop) Pass(ctx context.Context, i int) (*Prediction, error) {
	panic("not implemented")
}
