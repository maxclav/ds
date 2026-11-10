package multimap

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Multimap maps each key to one or more values. Len counts key-value pairs.
type Multimap[K, V any] interface {
	container.Container
	Put(key K, value V)
	// Get yields the values stored under key.
	Get(key K) iter.Seq[V]
	// Remove removes one key-value pair and reports whether it was present.
	Remove(key K, value V) bool
	// RemoveAll removes key and all its values.
	RemoveAll(key K)
	// All yields every key-value pair.
	All() iter.Seq2[K, V]
}
