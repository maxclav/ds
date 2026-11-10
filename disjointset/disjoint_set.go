package disjointset

import "fmt"

// DisjointSet splits elements, numbered from 0, into groups that don't overlap.
type DisjointSet interface {
	fmt.Stringer
	// Add adds an element in a group of its own and returns its number.
	Add() int
	// Union merges the groups of a and b and reports whether they were different.
	Union(a, b int) bool
	// Find returns the element that represents x's group.
	Find(x int) int
	Connected(a, b int) bool
	// Count returns the number of groups.
	Count() int
	// Len returns the number of elements.
	Len() int
}
