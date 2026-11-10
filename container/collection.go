package container

import "iter"

// Collection is a Container whose elements can be walked one by one.
type Collection[T any] interface {
	Container
	// All yields every element.
	All() iter.Seq[T]
}
