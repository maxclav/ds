package searchtree

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// SearchTree keeps its keys in sorted order.
type SearchTree[K, V any] interface {
	container.Container
	Insert(key K, value V)
	Search(key K) (value V, ok bool)
	Delete(key K) (ok bool)
	Min() (key K, value V, ok bool)
	Max() (key K, value V, ok bool)
	// Floor returns the largest key less than or equal to key.
	Floor(key K) (k K, value V, ok bool)
	// Ceiling returns the smallest key greater than or equal to key.
	Ceiling(key K) (k K, value V, ok bool)
	// Rank returns the number of keys less than key.
	Rank(key K) int
	// Select returns the entry with the i-th smallest key, counting from 0.
	Select(i int) (key K, value V, ok bool)
	// All yields the entries in key order (an in-order walk).
	All() iter.Seq2[K, V]
	// Backward yields the entries in reverse key order.
	Backward() iter.Seq2[K, V]
}
