package cache

import "github.com/maxclav/ds/container"

// Cache maps keys to values and evicts an entry when it is full.
// Implementations differ only in which entry they evict.
type Cache[K comparable, V any] interface {
	container.Container
	Get(key K) (value V, ok bool)
	Put(key K, value V)
	// Remove deletes key and reports whether it was present.
	Remove(key K) bool
}
