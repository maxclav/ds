package rangequery

import "fmt"

// RangeQuery answers questions about a range of positions while values change.
type RangeQuery[T any] interface {
	fmt.Stringer
	// Update sets the value at position i.
	Update(i int, value T)
	// Query combines the values at positions l to r-1, like the slice s[l:r].
	Query(l, r int) T
	// Len returns the number of positions.
	Len() int
}
