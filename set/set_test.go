package set

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/maxclav/ds/container"
)

// mapSet is a minimal Set used to test the package functions.
type mapSet[T comparable] map[T]struct{}

func newMapSet[T comparable](values ...T) mapSet[T] {
	s := mapSet[T]{}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

func (s mapSet[T]) Add(value T) bool {
	if s.Contains(value) {
		return false
	}
	s[value] = struct{}{}
	return true
}

func (s mapSet[T]) Remove(value T) bool {
	if !s.Contains(value) {
		return false
	}
	delete(s, value)
	return true
}

func (s mapSet[T]) Clear()                { clear(s) }
func (s mapSet[T]) Contains(value T) bool { _, ok := s[value]; return ok }
func (s mapSet[T]) Len() int              { return len(s) }
func (s mapSet[T]) All() iter.Seq[T]      { return maps.Keys(s) }
func (s mapSet[T]) String() string        { return container.String(s.All()) }

func TestOperations(t *testing.T) {
	a, b := newMapSet(1, 2, 3), newMapSet(2, 3, 4)
	tests := []struct {
		name string
		seq  iter.Seq[int]
		want []int
	}{
		{"Union", Union(a, b), []int{1, 2, 3, 4}},
		{"Intersection", Intersection(a, b), []int{2, 3}},
		{"Intersection larger first", Intersection(newMapSet(1, 2, 3, 4, 5), b), []int{2, 3, 4}},
		{"Difference", Difference(a, b), []int{1}},
		{"SymmetricDifference", SymmetricDifference(a, b), []int{1, 4}},
	}
	for _, tt := range tests {
		if got := slices.Sorted(tt.seq); !slices.Equal(got, tt.want) {
			t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
		}
		for range tt.seq {
			break // stopping early must not panic
		}
	}
}

func TestRelations(t *testing.T) {
	small, big, other := newMapSet(1, 2), newMapSet(1, 2, 3), newMapSet(7)
	if !IsSubset(small, big) || IsSubset(big, small) {
		t.Error("IsSubset")
	}
	if !Equal(small, newMapSet(2, 1)) || Equal(small, big) {
		t.Error("Equal")
	}
	if !IsDisjoint(small, other) || IsDisjoint(small, big) {
		t.Error("IsDisjoint")
	}
}
