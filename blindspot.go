package zoltar

import "go/token"

// A BlindSpot is something the predictor can't see about a value, and why.
// See [Value.BlindSpots].
//
// A BlindSpot is also an error, returned by the methods that can fail to give
// an answer at all, such as [Prediction.Targets].
type BlindSpot struct{}

// Kind returns the sort of limitation that caused the blind spot.
//
// A value worked out from an unknown value, such as an unknown number plus 1,
// has a blind spot of the same kind, so the kind names the original cause.
func (b BlindSpot) Kind() BlindSpotKind {
	panic("not implemented")
}

// Pos returns where the blind spot arose, such as the call that wasn't
// followed, or [token.NoPos] if it has no position.
func (b BlindSpot) Pos() token.Pos {
	panic("not implemented")
}

// String returns the explanation, such as "example.com/app.mustLoad never
// returns".
func (b BlindSpot) String() string {
	panic("not implemented")
}

// Error returns the same as [BlindSpot.String].
func (b BlindSpot) Error() string {
	panic("not implemented")
}

// BlindSpotKind is the sort of limitation that causes a [BlindSpot].
type BlindSpotKind int

const (
	// Input is a blind spot in a value that came from an unknown argument, or
	// from calling a method on a [Sentinel].
	Input BlindSpotKind = iota

	// NotFollowed is a blind spot in a value that came from code the predictor
	// doesn't follow, such as a function with no body, or one it makes
	// assumptions about instead.
	NotFollowed

	// Unsupported is a blind spot in a value that involves something the
	// predictor doesn't model.
	Unsupported

	// Limit is a blind spot in a value that the predictor stopped working out
	// because it reached one of its limits, such as how many calls deep it
	// follows, or how much work it does.
	Limit

	// Unpredictable is a blind spot in a value that can't be known without
	// running the program, such as the result of a select statement.
	Unpredictable
)

// String returns the name of the kind, such as "Limit".
func (k BlindSpotKind) String() string {
	panic("not implemented")
}
