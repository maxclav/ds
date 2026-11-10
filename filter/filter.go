package filter

import "fmt"

// Filter tests membership approximately, using little memory: MayContain can
// return true for a value that was never added, but never false for one that was.
type Filter[T any] interface {
	fmt.Stringer
	Add(value T)
	MayContain(value T) bool
}
