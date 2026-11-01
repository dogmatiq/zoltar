package zoltar

// Elements returns what's known about the elements of v, which must be a slice
// or array. ok is false if v differs between timelines.
func (v Value) Elements() (_ SliceElements, ok bool) {
	panic("not implemented")
}

// SliceElements is what's known about the elements of a slice or array. See
// [Value.Elements].
type SliceElements struct{}

// All returns the elements that may be in the slice. They aren't necessarily
// in index order, and a position that may be written more than once has more
// than one element.
func (s SliceElements) All() []Element {
	panic("not implemented")
}

// IsComplete returns true if [SliceElements.All] lists every element, and the
// index of each is known.
func (s SliceElements) IsComplete() bool {
	panic("not implemented")
}

// An Element is an element that may be in a slice or array.
type Element struct{}

// Value returns the element itself.
func (e Element) Value() Value {
	panic("not implemented")
}

// Index returns the element's position in the slice. ok is false if it isn't
// known.
func (e Element) Index() (i int, ok bool) {
	panic("not implemented")
}

// IsConditional returns true if the element is only written on some
// timelines, so it may not be in the slice at all.
func (e Element) IsConditional() bool {
	panic("not implemented")
}

// MapIndex returns the value stored under key in v, which must be a map, or
// the zero value if there's none. If the key may or may not be present, the
// zero value is one of the timelines.
//
// It panics if key isn't assignable to the map's key type.
func (v Value) MapIndex(key Value) Value {
	panic("not implemented")
}

// MapEntries returns what's known about the entries of v, which must be a
// map. ok is false if v differs between timelines.
func (v Value) MapEntries() (_ MapEntries, ok bool) {
	panic("not implemented")
}

// MapEntries is what's known about the entries of a map. See
// [Value.MapEntries].
type MapEntries struct{}

// All returns the entries that may be in the map, in no particular order.
func (m MapEntries) All() []MapEntry {
	panic("not implemented")
}

// IsComplete returns true if [MapEntries.All] lists every entry, and every key
// is known.
func (m MapEntries) IsComplete() bool {
	panic("not implemented")
}

// A MapEntry is an entry that may be in a map.
type MapEntry struct{}

// Key returns the entry's key.
func (e MapEntry) Key() Value {
	panic("not implemented")
}

// Value returns the value stored under the entry's key.
func (e MapEntry) Value() Value {
	panic("not implemented")
}

// IsConditional returns true if the entry is only written, or only not
// deleted, on some timelines, so it may not be in the map at all.
func (e MapEntry) IsConditional() bool {
	panic("not implemented")
}
