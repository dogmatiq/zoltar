package zoltar

import (
	"context"

	"golang.org/x/tools/go/ssa"
)

// A Target is a function that a call may run, along with the arguments it
// receives. See [Prediction.CallTargets].
type Target struct{}

// Func returns the function.
func (t Target) Func() *ssa.Function {
	panic("not implemented")
}

// Args returns the arguments.
//
// If the function is a method, the first argument is the receiver.
func (t Target) Args() []Value {
	panic("not implemented")
}

// Follow predicts t's function, called with t's arguments.
//
// The result's [Prediction.Parent] is the prediction that contains the call,
// and its [Prediction.CallSite] is the call itself.
//
// It follows the function even if the predictor has assumptions about it,
// which would otherwise be used in its place.
//
// If the function can't be followed, such as when it has no body, the error
// is a [BlindSpot]. It panics if t is the zero Target.
func (t Target) Follow(ctx context.Context) (*Prediction, error) {
	_ = ctx

	panic("not implemented")
}

// FollowAll is like [Target.Follow], but makes one prediction for several
// targets that have the same function, such as a method called on one of
// several possible receivers.
//
// Each argument is the combination of that argument from every target, as for
// [Combine].
//
// It panics if targets is empty, or if they don't all come from the same
// call to [Prediction.CallTargets] and have the same function.
func FollowAll(ctx context.Context, targets []Target) (*Prediction, error) {
	_ = ctx
	_ = targets

	panic("not implemented")
}
