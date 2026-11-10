package queue

import "iter"

// Deque is a double-ended queue: values go in and out at both ends.
// Every Deque is also a Queue: Enqueue adds at the back, Dequeue removes
// from the front and Peek reads the front.
type Deque[T any] interface {
	Queue[T]
	PushFront(value T)
	PushBack(value T)
	PopFront() (value T, ok bool)
	PopBack() (value T, ok bool)
	Front() (value T, ok bool)
	Back() (value T, ok bool)
	// Backward yields the values from back to front.
	Backward() iter.Seq[T]
}
