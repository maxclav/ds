package maps

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Map stores a value for each key.
type Map[K, V any] interface {
	container.Container
	// Put sets the value for key, replacing any previous one.
	Put(key K, value V)
	Get(key K) (value V, ok bool)
	// Remove deletes key and reports whether it was present.
	Remove(key K) bool
	Contains(key K) bool
	// All yields every key with its value.
	All() iter.Seq2[K, V]
}
