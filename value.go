package zoltar

import (
	"go/types"
	"iter"
)

// A Value is a prediction of a value in the program.
//
// A value is never wrong, only more or less specific. It may differ between
// timelines, such as a variable that is assigned a different value on each
// branch of an if statement. Methods that return a [Value] combine the result
// from each timeline, and methods that return ok succeed only if every
// timeline gives the same result.
//
// Everything inside a value, such as what a pointer points to, is as of the
// point it was predicted at. See [Prediction.ValueOf].
//
// A value held in an interface is the value itself: there is no separate
// interface level, as there is in package reflect. Methods depend on the
// dynamic type. Calling one that the type doesn't support, such as Field on an
// int, panics. If the dynamic type isn't known, methods such as Field and Elem
// return an unknown value whose static type is the empty interface.
//
// The zero Value is not valid. Its methods panic, except IsValid and String.
type Value struct {
	v *value
}

// value is the internal representation of a [Value].
type value struct{}

// IsValid returns true if v is not the zero [Value].
func (v Value) IsValid() bool {
	panic("not implemented")
}

// String returns a description of v, for debugging. Its format may change.
func (v Value) String() string {
	panic("not implemented")
}

// StaticType returns the type of the expression that v came from, such as the
// type of x for [Prediction.ValueOf] or a field's declared type for
// [Value.Field]. It may be an interface.
func (v Value) StaticType() types.Type {
	panic("not implemented")
}

// DynamicType returns the type of v at run time, such as *bytes.Buffer for a
// value held in an io.Writer. If v's static type isn't an interface, it's the
// same as [Value.StaticType].
//
// ok is false if the type isn't known or differs between timelines, if v is a
// nil interface, or if v is a [Sentinel] of an interface type.
func (v Value) DynamicType() (t types.Type, ok bool) {
	panic("not implemented")
}

// HasSingleTimeline returns true if v is the same on every timeline. Such a
// value may still have blind spots.
func (v Value) HasSingleTimeline() bool {
	panic("not implemented")
}

// Timelines returns the value on each timeline, merging timelines that have
// the same value.
//
// A timeline is possible as far as the predictor can tell. If another
// timeline has blind spots, this one may never actually happen.
func (v Value) Timelines() []Value {
	panic("not implemented")
}

// HasBlindSpots returns true if something about v itself is unknown on any
// timeline: its dynamic type, its constant value, or the variable it refers
// to. What's unknown inside v, such as a field, is reported by that value.
//
// A [Sentinel] has no blind spots.
func (v Value) HasBlindSpots() bool {
	panic("not implemented")
}

// BlindSpots returns why something about v is unknown, one [BlindSpot] per
// cause. See [Value.HasBlindSpots].
func (v Value) BlindSpots() []BlindSpot {
	panic("not implemented")
}

// As returns v as a T, if v is a known constant whose type has the same kind
// as T, as in package reflect. For example, As[int] accepts a constant of type
// int or time.Month, but not int64 or float64.
func (v Value) As[T ~bool | ~string |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64]() (result T, ok bool) {
	panic("not implemented")
}

// IsNil returns whether v is nil. known is false if that can't be predicted.
// A nil pointer held in an interface isn't a nil interface.
//
// v must be a pointer, interface, slice, map, channel or function.
func (v Value) IsNil() (isNil, known bool) {
	panic("not implemented")
}

// Equal returns true if v and x are known to be the same value. That includes
// the same unknown value, such as two results of a [Pure] function called
// with equal arguments.
//
// Otherwise, a value with blind spots isn't equal to anything, not even
// itself, and neither is a value that differs between timelines. A [Sentinel]
// is equal only to itself. Values that refer to memory are equal if they
// refer to the same variable as of the same point.
func (v Value) Equal(x Value) bool {
	panic("not implemented")
}

// NumField returns the number of fields in v, which must be a struct, or 0 if
// it isn't known.
func (v Value) NumField() int {
	panic("not implemented")
}

// Field returns v's i'th field. v must be a struct. It panics if i is out of
// range.
func (v Value) Field(i int) Value {
	panic("not implemented")
}

// FieldByName returns v's field with the given name, including fields of
// embedded structs, as in package reflect. v must be a struct. It panics if
// there's no such field.
func (v Value) FieldByName(name string) Value {
	panic("not implemented")
}

// Elem returns the value that v points to. v must be a pointer.
func (v Value) Elem() Value {
	panic("not implemented")
}

// Len returns the number of elements in v, which must be a slice, array,
// string, map or channel. ok is false if it isn't known, which is always the
// case for a channel.
func (v Value) Len() (n int, ok bool) {
	panic("not implemented")
}

// Index returns v's i'th element, or for a string, its i'th byte. v must be a
// slice, array or string.
//
// Timelines on which i is out of range are left out. It panics if i is
// negative, or out of range on every timeline.
func (v Value) Index(i int) Value {
	panic("not implemented")
}

// Reachable yields v and everything reachable from it on every timeline: its
// timelines, elements, fields, map entries, what it points to, and the
// variables it captures.
//
// Each variable is visited once, so cycles end. An array with more than a few
// elements is visited as a single value that combines its elements.
func (v Value) Reachable() iter.Seq[Value] {
	panic("not implemented")
}
