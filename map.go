package zoltar

// MapEntries is what's known about the entries of a map.
//
// See [Value.MapEntries].
type MapEntries struct{}

// All returns the entries that may be in the map. They're in no particular
// order, and a key that may be written more than once has more than one entry.
func (m MapEntries) All() []MapEntry {
	panic("not implemented")
}

// IsComplete returns true if [MapEntries.All] lists every entry, and the key
// of each is known.
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

// IsConditional returns true if the entry is only present on some timelines,
// such as when it's written or deleted by an if statement, so it may not be in
// the map at all.
func (e MapEntry) IsConditional() bool {
	panic("not implemented")
}
