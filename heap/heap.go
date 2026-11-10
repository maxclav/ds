package heap

import "github.com/maxclav/ds/container"

// Heap gives back its values smallest first, by its compare function.
type Heap[T any] interface {
	container.Container
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
