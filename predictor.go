package zoltar

import (
	"context"

	"golang.org/x/tools/go/ssa"
)

// A Predictor predicts the values in a program without executing it.
//
// Methods that take a [context.Context] return its error if it's done. The
// predictor, and everything it returned, is still usable afterwards.
type Predictor struct{}

// New returns a [Predictor] for prog, which must already be built. See
// [ssa.Program.Build].
//
// It panics if [ReturnsResultOf] was given something that isn't a function, or
// a function of the wrong shape.
func New(prog *ssa.Program, opts ...Option) *Predictor {
	panic("not implemented")
}

// Predict returns a prediction about what happens when a function within the
// program is called with specific arguments.
//
// fn is the function to predict.
//
// args is the set of arguments to pass to the function call. If fn is a method,
// the first argument is the receiver.
//
// To construct an argument, use [Known], [KnownAs], [Sentinel] or [Unknown].
// Missing arguments, or the zero [Value], are treated as unknown.
func (p *Predictor) Predict(
	ctx context.Context,
	fn *ssa.Function,
	args ...Value,
) (*Prediction, error) {
	panic("not implemented")
}

// An Option changes how a [Predictor] analyzes a program.
type Option func(*options)

// options is the configuration built by applying each [Option].
type options struct{}
