package set

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Set is a collection of unique values.
type Set[T any] interface {
	container.Collection[T]
	// Add adds value and reports whether it was new.
	Add(value T) bool
	// Remove removes value and reports whether it was present.
	Remove(value T) bool
	Contains(value T) bool
}

// Union yields every value in a or b, once.
func Union[T any](a, b Set[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range a.All() {
			if !yield(v) {
				return
			}
		}
		for v := range b.All() {
			if !a.Contains(v) && !yield(v) {
				return
			}
		}
	}
}

// Intersection yields the values in both a and b.
func Intersection[T any](a, b Set[T]) iter.Seq[T] {
	if a.Len() > b.Len() {
		a, b = b, a // walk the smaller set
	}
	return func(yield func(T) bool) {
		for v := range a.All() {
			if b.Contains(v) && !yield(v) {
				return
			}
		}
	}
}

// Difference yields the values in a that are not in b.
func Difference[T any](a, b Set[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range a.All() {
			if !b.Contains(v) && !yield(v) {
				return
			}
		}
	}
}

// SymmetricDifference yields the values in exactly one of a and b.
func SymmetricDifference[T any](a, b Set[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range Difference(a, b) {
			if !yield(v) {
				return
			}
		}
		for v := range Difference(b, a) {
			if !yield(v) {
				return
			}
		}
	}
}

// IsSubset reports whether every value in a is also in b.
func IsSubset[T any](a, b Set[T]) bool {
	if a.Len() > b.Len() {
		return false
	}
	for v := range a.All() {
		if !b.Contains(v) {
			return false
		}
	}
	return true
}

// Equal reports whether a and b hold the same values.
func Equal[T any](a, b Set[T]) bool {
	return a.Len() == b.Len() && IsSubset(a, b)
}

// IsDisjoint reports whether a and b have no value in common.
func IsDisjoint[T any](a, b Set[T]) bool {
	for range Intersection(a, b) {
		return false
	}
	return true
}
