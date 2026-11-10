package container

import (
	"fmt"
	"iter"
	"strings"
)

// Container is the base of every collection: it knows its size, can be
// emptied and prints its contents.
type Container interface {
	fmt.Stringer
	// Len returns the number of elements.
	Len() int
	// Clear removes every element.
	Clear()
}

// String formats the values like Go prints a slice: [1 2 3].
func String[T any](seq iter.Seq[T]) string {
	var b strings.Builder
	b.WriteByte('[')
	sep := ""
	for v := range seq {
		fmt.Fprintf(&b, "%s%v", sep, v)
		sep = " "
	}
	b.WriteByte(']')
	return b.String()
}

// String2 formats the pairs like Go prints a map: map[a:1 b:2].
func String2[K, V any](seq iter.Seq2[K, V]) string {
	var b strings.Builder
	b.WriteString("map[")
	sep := ""
	for k, v := range seq {
		fmt.Fprintf(&b, "%s%v:%v", sep, k, v)
		sep = " "
	}
	b.WriteByte(']')
	return b.String()
}
