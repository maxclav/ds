package multiset

import "github.com/maxclav/ds/container"

// Multiset is a set that keeps count of repeated values.
// Len counts repeats, and All yields each value as many times as it is present.
type Multiset[T any] interface {
	container.Collection[T]
	Add(value T)
	// Remove removes one occurrence of value and reports whether there was one.
	Remove(value T) bool
	// Count returns how many times value is present.
	Count(value T) int
}
