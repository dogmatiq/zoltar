package zoltar

// SliceElements is what's known about the elements of a slice or array.
//
// See [Value.SliceElements].
type SliceElements struct{}

// All returns the elements that may be in the slice. They aren't necessarily
// in index order, and a position that may be written more than once has more
// than one element.
func (s SliceElements) All() []SliceElement {
	panic("not implemented")
}

// IsComplete returns true if [SliceElements.All] lists every element, and the
// index of each is known.
func (s SliceElements) IsComplete() bool {
	panic("not implemented")
}

// A SliceElement is an element that may be in a slice or array.
type SliceElement struct{}

// Index returns the element's position in the slice.
//
// ok is false if it isn't known.
func (e SliceElement) Index() (i int, ok bool) {
	panic("not implemented")
}

// Value returns the element itself.
func (e SliceElement) Value() Value {
	panic("not implemented")
}

// IsConditional returns true if the element is only written on some
// timelines, so it may not be in the slice at all.
func (e SliceElement) IsConditional() bool {
	panic("not implemented")
}
