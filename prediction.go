package zoltar

import (
	"context"
	"iter"

	"golang.org/x/tools/go/ssa"
)

// A Prediction describes what happens when a function is called with specific
// arguments.
//
// A [Value] that refers to memory, such as a pointer, belongs to the
// prediction that produced it. It may be used with any prediction related to
// that one by followed calls, and passed as an argument to any prediction.
// Using it with any other prediction, including another pass of the same loop,
// panics.
type Prediction struct{}

// ValueOf returns the value of v as of just before the instruction "before"
// runs.
//
// v is any [ssa.Value] in p's function, such as a parameter, a constant or the
// result of an instruction.
//
// before is an instruction in p's function, typically the one that uses v.
// For a value that refers to memory, only the writes that can happen before it
// are counted.
//
// It panics if v or before isn't in p's function, or before is nil.
func (p *Prediction) ValueOf(ctx context.Context, v ssa.Value, before ssa.Instruction) (Value, error) {
	panic("not implemented")
}

// Returns returns the values that p's function returns, one per result.
func (p *Prediction) Returns(ctx context.Context) ([]Value, error) {
	panic("not implemented")
}

// Targets returns the functions that a call may run, along with the arguments
// each receives. A call through an interface or a function value may have
// more than one.
//
// call is a call instruction in p's function.
//
// If the targets can't be predicted, such as for a call to a method of a
// [Sentinel], the error is a [BlindSpot]. It panics if call isn't in p's
// function.
func (p *Prediction) Targets(ctx context.Context, call ssa.CallInstruction) ([]Target, error) {
	panic("not implemented")
}

// Blocks yields each block of p's function that can run, starting with its
// first block. See [Prediction.CanRun].
func (p *Prediction) Blocks() iter.Seq[*ssa.BasicBlock] {
	panic("not implemented")
}

// NextBlocks yields the blocks that can run immediately after b.
func (p *Prediction) NextBlocks(b *ssa.BasicBlock) iter.Seq[*ssa.BasicBlock] {
	panic("not implemented")
}

// CanRun returns true if b can run: some path to it doesn't depend on a
// condition that is predicted to be false, or pass through a call to a
// function that doesn't return.
func (p *Prediction) CanRun(b *ssa.BasicBlock) bool {
	panic("not implemented")
}

// CanReach returns true if "to" can run after "from". It panics if either
// block isn't in p's function.
func (p *Prediction) CanReach(from, to *ssa.BasicBlock) bool {
	panic("not implemented")
}

// AlwaysRuns returns true if b runs whenever p's function returns normally.
// Paths that end in a panic aren't counted, unless [Prediction.MayRecover] is
// true.
func (p *Prediction) AlwaysRuns(b *ssa.BasicBlock) bool {
	panic("not implemented")
}

// MayRecover returns true if a deferred call in p's function may call
// recover(), so that the function may return normally after a panic.
func (p *Prediction) MayRecover() bool {
	panic("not implemented")
}

// Function returns the function whose call p predicts.
func (p *Prediction) Function() *ssa.Function {
	panic("not implemented")
}

// Parent returns the prediction that p was made from: the one containing the
// call that p follows, or the loop that p is a pass of. It returns nil if p
// was made by [Predictor.Predict] or [Func.Predict].
func (p *Prediction) Parent() *Prediction {
	panic("not implemented")
}

// CallSite returns the call in [Prediction.Parent] that p follows, or nil if
// p isn't the prediction of a followed call.
func (p *Prediction) CallSite() ssa.CallInstruction {
	panic("not implemented")
}

// A Target is a function that a call may run, along with the arguments it
// receives. See [Prediction.Targets].
type Target struct{}

// Func returns the function.
func (t Target) Func() *ssa.Function {
	panic("not implemented")
}

// Args returns the arguments. If the function is a method, the first argument
// is the receiver.
func (t Target) Args() []Value {
	panic("not implemented")
}

// Follow returns a prediction of t's function, as called from the call that t
// came from.
//
// Unlike the predictor's own handling of calls, it ignores assumptions.
//
// If the function can't be followed, such as when it has no body, the error
// is a [BlindSpot]. It panics if t didn't come from [Prediction.Targets].
func (t Target) Follow(ctx context.Context) (*Prediction, error) {
	panic("not implemented")
}

// FollowAll is like [Target.Follow], but returns a single prediction that
// covers every one of ts, combining their arguments. See [Combine].
//
// It panics if ts is empty, or if its targets don't all come from the same
// call to [Prediction.Targets] and have the same function.
func FollowAll(ctx context.Context, ts []Target) (*Prediction, error) {
	panic("not implemented")
}
