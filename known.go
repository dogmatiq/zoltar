package zoltar

import "go/types"

// Known returns a known constant. Its static type is the predeclared type that
// matches T, such as int for Known(1) or string for Known("hello").
//
// For one of the program's own types, use [KnownAs]. It panics if v is NaN or
// an infinity.
func Known[T bool | string |
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 | uintptr |
	float32 | float64](v T) Value {
	panic("not implemented")
}

// KnownAs returns a known constant of type t, which may be one of the
// program's named types, such as time.Month.
//
// t's underlying type must have the same kind as T, as for [Value.As]. It
// panics if it doesn't, or if v is NaN or an infinity.
func KnownAs[T ~bool | ~string |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64](t types.Type, v T) Value {
	panic("not implemented")
}

// Zero returns the zero value of type t, such as 0, "", false, or nil.
func Zero(t types.Type) Value {
	panic("not implemented")
}

// Unknown returns a value of static type t whose contents are unknown. Its
// [BlindSpot] has kind [Input] and the given reason.
func Unknown(t types.Type, reason string) Value {
	panic("not implemented")
}

// Sentinel returns a new, unique value of type t, which you can recognize
// wherever it ends up with [Value.Equal], like a sentinel error such as
// io.EOF. For example, pass a sentinel for a parameter to find the calls made
// on it.
//
// A sentinel has no blind spots. If t is an interface, it has no dynamic type.
// Calls to its methods aren't followed: their results have a [BlindSpot] of
// kind [Input].
func Sentinel(t types.Type) Value {
	panic("not implemented")
}

// Combine returns a value whose timelines are those of vs, which must all have
// the same static type. It panics if they don't, or if vs is empty.
func Combine(vs ...Value) Value {
	panic("not implemented")
}
