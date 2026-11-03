package zoltar

import "go/types"

// A Value is a prediction of a value in the program.
//
// A value may differ between timelines, such as a variable that is assigned a
// different value on each branch of an if statement. Methods that return a
// [Value] combine the result from each timeline. See [Value.Timelines].
//
// Various aspects of a value may be unknown. See [Value.BlindSpots].
//
// Everything inside a value, such as what a pointer points to, describes the
// program at the point where the value was predicted. See [Prediction.ValueOf].
//
// If a value's static type is an interface, its methods act on its dynamic
// value. If the dynamic type isn't known, methods such as FieldByName and Elem
// return an unknown value whose static type is the empty interface.
type Value struct{}

// Known returns a known constant with the value v.
//
// Its static type is the predeclared type that matches T, such as int for
// Known(1) or string for Known("hello"). For one of the program's own types,
// use [KnownAs].
//
// It panics if v is NaN or an infinity.
func Known[T bool | string |
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 | uintptr |
	float32 | float64](v T) Value {
	_ = v

	panic("not implemented")
}

// KnownAs returns a known constant of type t with the value v.
//
// t may be one of the program's named types, such as time.Month. Its
// underlying type must have the same kind as T, as for [Value.As].
//
// It panics if t's underlying type doesn't have the same kind as T, or if v is
// NaN or an infinity.
func KnownAs[T ~bool | ~string |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64](t types.Type, v T) Value {
	_ = t
	_ = v

	panic("not implemented")
}

// Zero returns the zero value of type t, such as 0, "", false or nil.
func Zero(t types.Type) Value {
	_ = t

	panic("not implemented")
}

// Unknown returns a value of static type t whose contents are unknown.
//
// reason explains why it's unknown, such as "chosen by the caller". The value
// has a blind spot of kind [InputBlindSpot] with that reason.
func Unknown(t types.Type, reason string) Value {
	_ = t
	_ = reason

	panic("not implemented")
}

// Sentinel returns a new, unique value of type t.
//
// Like a sentinel error such as io.EOF, it can be recognized wherever it ends
// up, using [Value.Equal]. For example, pass a sentinel for a parameter to find
// the calls made on it.
//
// A sentinel has no blind spots, but what's inside it, such as its fields or
// what it points to, is unknown. If t is an interface, it has no dynamic type.
//
// Method calls on a sentinel aren't followed: their results have a [BlindSpot]
// of kind [InputBlindSpot].
func Sentinel(t types.Type) Value {
	_ = t

	panic("not implemented")
}

// Combine returns a value whose timelines are those of the given values.
//
// It panics if values is empty, if they don't all have the same static type, or
// if two of them refer to variables from separate calls to [Predictor.Predict].
func Combine(values ...Value) Value {
	_ = values

	panic("not implemented")
}

// IsValid returns true if v is not the zero [Value].
func (v Value) IsValid() bool {
	panic("not implemented")
}

// StaticType returns v's type as declared in the program, such as the
// declared type of a variable or field.
//
// It may be an interface type. See [Value.DynamicType].
func (v Value) StaticType() types.Type {
	panic("not implemented")
}

// DynamicType returns v's type when the program runs, such as *bytes.Buffer for
// a value held in an io.Writer.
//
// ok is false if it isn't known, or differs between timelines. If v is a nil
// interface, or a [Sentinel] of an interface type, it has no dynamic type, so
// ok is also false.
//
// If v's static type is concrete, it's the same as [Value.StaticType].
func (v Value) DynamicType() (t types.Type, ok bool) {
	panic("not implemented")
}

// HasSingleTimeline returns true if v is the same on every timeline.
//
// A value with a single timeline may still have blind spots.
func (v Value) HasSingleTimeline() bool {
	panic("not implemented")
}

// Timelines returns the distinct values of v on each of its timelines.
//
// Between them, the timelines cover every way the program can run. Each one
// may happen, but none is certain to.
func (v Value) Timelines() []Value {
	panic("not implemented")
}

// HasBlindSpots returns true if something about v itself is unknown on any
// timeline.
//
// Blind spots inside v belong to the inner value. For example, a struct
// with an unknown field has no blind spots of its own, but the value returned
// by [Value.FieldByName] does.
//
// A [Sentinel] has no blind spots.
func (v Value) HasBlindSpots() bool {
	panic("not implemented")
}

// BlindSpots returns the reasons that something about v is unknown, one per
// cause..
func (v Value) BlindSpots() []BlindSpot {
	panic("not implemented")
}

// As returns v as a value of type T.
//
// ok is false if v isn't a known constant, differs between timelines, or has a
// type whose kind isn't the same as T's.
func (v Value) As[
	T ~bool | ~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64,
]() (result T, ok bool) {
	panic("not implemented")
}

// IsNil returns true if v is nil.
//
// ok is false if that isn't known, or differs between timelines. isNil is only
// true if ok is also true, so if isNil is true, v is known to be nil.
//
// As in Go, a nil pointer held in an interface isn't a nil interface.
//
// It panics if v isn't a pointer, interface, slice, map, channel or function.
func (v Value) IsNil() (isNil, ok bool) {
	panic("not implemented")
}

// Equal returns true if v and x are the same whenever the program runs.
//
// ok is false if that isn't known, or differs between timelines. equal is only
// true if ok is also true, so if equal is true, v and x are known to be the
// same.
//
// Unknown values can be equal. For example, a parameter is equal to itself, and
// so are two results of a [Pure] function called with equal arguments. A
// [Sentinel] is equal only to itself. Values that refer to a variable, such as
// pointers, are equal if they refer to the same one.
//
// For values that refer to variables from separate calls to
// [Predictor.Predict], ok is always false.
func (v Value) Equal(x Value) (equal, ok bool) {
	_ = x

	panic("not implemented")
}

// Fields returns the values of v's fields, in the order they're declared.
//
// ok is false if v's dynamic type isn't known, or differs between timelines.
//
// It panics if v isn't a struct.
func (v Value) Fields() (_ []Value, ok bool) {
	panic("not implemented")
}

// FieldByName returns v's field with the given name, including fields promoted
// from embedded structs.
//
// It returns the zero [Value] if v has no such field on any timeline. A name
// that's ambiguous between embedded structs counts as missing.
//
// Timelines on which v has no such field, including those where an embedded
// pointer it must follow is nil, are left out. If the field's type differs
// between timelines, the result's static type is the empty interface.
//
// It panics if v isn't a struct.
func (v Value) FieldByName(name string) Value {
	_ = name

	panic("not implemented")
}

// Elem returns the value that v points to.
//
// It returns the zero [Value] if v is nil on every timeline. Timelines on which
// v is nil are left out.
//
// It panics if v isn't a pointer.
func (v Value) Elem() Value {
	panic("not implemented")
}

// Len returns the number of elements in v.
//
// ok is false if that isn't known, or differs between timelines. It's never
// known for a channel.
//
// It panics if v isn't a slice, array, string, map or channel.
func (v Value) Len() (n int, ok bool) {
	panic("not implemented")
}

// Index returns element i of v, or for a string, byte i. It returns the zero
// [Value] if i is out of range on every timeline.
//
// Timelines on which i is out of range are left out.
//
// It panics if v isn't a slice, array or string, or if i is negative.
func (v Value) Index(i int) Value {
	_ = i

	panic("not implemented")
}

// SliceElements returns what's known about the elements of v.
//
// ok is false if v's dynamic type isn't known, or v differs between timelines.
//
// If v's contents aren't known, the result has no elements and is incomplete.
// See [SliceElements.IsComplete].
//
// It panics if v isn't a slice or array.
func (v Value) SliceElements() (_ SliceElements, ok bool) {
	panic("not implemented")
}

// MapIndex returns the element stored under key in v.
//
// Timelines on which key isn't present give the zero value of the map's
// element type, as for m[key] in Go. If key isn't known, the result includes
// every element it might select. A key that refers to a variable from a
// separate call to [Predictor.Predict] counts as unknown.
//
// It panics if v isn't a map, or if key isn't assignable to the map's key
// type.
func (v Value) MapIndex(key Value) Value {
	_ = key

	panic("not implemented")
}

// MapEntries returns what's known about the entries of v.
//
// ok is false if v's dynamic type isn't known, or v differs between timelines.
//
// If v's contents aren't known, the result has no entries and is incomplete.
// See [MapEntries.IsComplete].
//
// It panics if v isn't a map.
func (v Value) MapEntries() (_ MapEntries, ok bool) {
	panic("not implemented")
}

// Func returns v as a function value.
//
// ok is false if the function isn't known, or differs between timelines.
// Timelines on which v is nil are left out, and ok is false if v is nil on
// every timeline.
//
// It panics if v isn't a function.
func (v Value) Func() (_ Func, ok bool) {
	panic("not implemented")
}

// Origin returns where the variable that v refers to was allocated, such as
// the variable that a pointer points to, or the array behind a slice.
//
// ok is false if that isn't known, or differs between timelines. Timelines on
// which v is nil are left out. ok is also false if v doesn't refer to a
// variable, such as a constant, or is nil on every timeline.
func (v Value) Origin() (_ Origin, ok bool) {
	panic("not implemented")
}

func (v Value) String() string {
	panic("not implemented")
}
