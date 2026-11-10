package tree

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Tree is a rooted tree. Len counts nodes.
type Tree[T any] interface {
	container.Container
	PreOrder() iter.Seq[T]
	PostOrder() iter.Seq[T]
	// LevelOrder yields the values one level at a time, from the root down.
	LevelOrder() iter.Seq[[]T]
	// Height returns the number of levels (0 for an empty tree).
	Height() int
}
