package trie

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Trie stores words and answers prefix queries. All yields the words in sorted order.
type Trie interface {
	container.Collection[string]
	Insert(word string)
	// Delete removes word and reports whether it was present.
	Delete(word string) bool
	Search(word string) bool
	StartsWith(prefix string) bool
	// WithPrefix yields the words that start with prefix, in sorted order.
	WithPrefix(prefix string) iter.Seq[string]
}
