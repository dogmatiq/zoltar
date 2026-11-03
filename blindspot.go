package zoltar

import "go/token"

// A BlindSpot is something the predictor can't see about a value, and why.
//
// See [Value.BlindSpots].
//
// A BlindSpot is also an error, returned by the methods that can fail to give
// an answer at all, such as [Prediction.CallTargets].
type BlindSpot struct{}

// Kind returns the sort of limitation that caused the blind spot.
//
// A value computed from an unknown value, such as an unknown number plus 1,
// has a blind spot of the same kind, so the kind names the original cause.
func (b BlindSpot) Kind() BlindSpotKind {
	panic("not implemented")
}

// Pos returns where the blind spot arose, such as the call that wasn't
// followed.
//
// It returns [token.NoPos] if the blind spot has no position.
func (b BlindSpot) Pos() token.Pos {
	panic("not implemented")
}

// String returns a description of the blind spot, such as "the result of
// os.Getenv, which isn't followed".
//
// Its format may change.
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
	// InputBlindSpot is a blind spot in a value that came from an unknown
	// argument, or from calling a method on a [Sentinel].
	InputBlindSpot BlindSpotKind = iota

	// NotFollowedBlindSpot is a blind spot in a value that came from code the
	// predictor doesn't follow, such as a function with no body, or one it
	// makes assumptions about instead.
	NotFollowedBlindSpot

	// UnsupportedBlindSpot is a blind spot in a value that involves something
	// the predictor doesn't model, such as reflection or package unsafe.
	UnsupportedBlindSpot

	// LimitBlindSpot is a blind spot in a value that the predictor stopped
	// computing because it reached an internal limit, such as on how much
	// work it does for one prediction.
	LimitBlindSpot

	// UnpredictableBlindSpot is a blind spot in a value that can't be known
	// without running the program, such as which case of a select statement
	// runs.
	UnpredictableBlindSpot
)

// String returns the name of the kind, such as "Limit" for [LimitBlindSpot].
func (k BlindSpotKind) String() string {
	panic("not implemented")
}
