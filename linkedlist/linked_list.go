package linkedlist

import "github.com/maxclav/ds/container"

// LinkedList is a sequence of nodes, each pointing to the next one.
// It has no PopBack: a singly linked list would need O(n) to find the node
// before the last one.
type LinkedList[T any] interface {
	container.Collection[T]
	PushFront(value T)
	PushBack(value T)
	PopFront() (value T, ok bool)
	Front() (value T, ok bool)
	Back() (value T, ok bool)
}
