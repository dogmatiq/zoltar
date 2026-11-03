package zoltar

import (
	"context"

	"golang.org/x/tools/go/ssa"
)

// A Prediction describes what happens when a function is called with specific
// arguments.
type Prediction struct{}

// Function returns the predicted function.
func (p *Prediction) Function() *ssa.Function {
	panic("not implemented")
}

// Parent returns the prediction that p came from, or nil if p was made by
// [Predictor.Predict].
//
// If p was made by [Target.Follow] or [FollowAll], it's the prediction that
// contains the call. If p is a pass of a loop, it's the prediction that
// contains the loop. If p was made by [Func.Predict], it's the prediction that
// the function value came from.
func (p *Prediction) Parent() *Prediction {
	panic("not implemented")
}

// CallSite returns the call that p follows, or nil if p wasn't made by
// [Target.Follow] or [FollowAll].
//
// The call is in the parent's function. See [Prediction.Parent].
func (p *Prediction) CallSite() ssa.CallInstruction {
	panic("not implemented")
}

// ValueOf returns the value of v just before the given instruction executes.
//
// v is a parameter, free variable or instruction of the predicted function, or
// a constant, global or function that it uses.
//
// before is an instruction of the predicted function, typically one that uses
// v. The result describes v, and anything it refers to, immediately before
// that instruction executes.
//
// It panics if v or before is nil, or if either isn't as described above.
func (p *Prediction) ValueOf(ctx context.Context, v ssa.Value, before ssa.Instruction) (Value, error) {
	_ = ctx
	_ = v
	_ = before

	panic("not implemented")
}

// ReturnValues returns the values that the predicted function returns.
func (p *Prediction) ReturnValues(ctx context.Context) ([]Value, error) {
	_ = ctx

	panic("not implemented")
}

// CallTargets returns the functions that the given call may run, along with the
// arguments each receives.
//
// call is a call instruction in the predicted function.
//
// If the function to call isn't known statically, such as for a call through
// an interface or a function variable, there is one target for each function
// it could run.
//
// If the targets can't be predicted, such as for a call to a method of a
// [Sentinel], the error is a [BlindSpot]. It panics if call isn't in the
// predicted function.
func (p *Prediction) CallTargets(ctx context.Context, call ssa.CallInstruction) ([]Target, error) {
	_ = ctx
	_ = call

	panic("not implemented")
}

// IsReachable returns true if b may execute when the function is called with
// the predicted arguments.
//
// It's equivalent to p.IsReachableFrom(b, p.Function().Blocks[0]).
func (p *Prediction) IsReachable(b *ssa.BasicBlock) bool {
	_ = b

	panic("not implemented")
}

// IsReachableFrom returns true if b is reachable from src, given the predicted
// arguments. Every block is reachable from itself.
//
// It panics if either block isn't in the predicted function.
func (p *Prediction) IsReachableFrom(b, src *ssa.BasicBlock) bool {
	_ = b
	_ = src

	panic("not implemented")
}

// ReachableBlocks returns the reachable blocks of the predicted function, in
// the order of [ssa.Function.Blocks]. See [Prediction.IsReachable].
func (p *Prediction) ReachableBlocks() []*ssa.BasicBlock {
	panic("not implemented")
}

// ReachableSuccessors returns the successors of b that may actually execute
// after it, given the predicted arguments, in the order of
// [ssa.BasicBlock.Succs].
func (p *Prediction) ReachableSuccessors(b *ssa.BasicBlock) []*ssa.BasicBlock {
	_ = b

	panic("not implemented")
}

// IsReachableOnEveryReturnPath returns true if every path on which the
// function returns passes through b, given the predicted arguments.
//
// It returns false if [Prediction.MayReturn] is false.
func (p *Prediction) IsReachableOnEveryReturnPath(b *ssa.BasicBlock) bool {
	_ = b

	panic("not implemented")
}

// MayReturn returns true if the function may return, given the predicted
// arguments.
//
// Recovering from a panic counts as returning.
func (p *Prediction) MayReturn() bool {
	panic("not implemented")
}

// MayRecover returns true if a deferred call in the predicted function may
// call recover(), so that the function may return normally after a panic.
func (p *Prediction) MayRecover() bool {
	panic("not implemented")
}

// Loops returns the loops in the predicted function that can be predicted one
// pass at a time. Specifically, range loops over a slice, an array or an
// integer.
func (p *Prediction) Loops() []Loop {
	panic("not implemented")
}
