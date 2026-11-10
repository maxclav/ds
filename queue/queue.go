package queue

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Queue is first in, first out: Dequeue returns the value enqueued earliest.
// All yields the values from front to back.
type Queue[T any] interface {
	container.Collection[T]
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)
}

func Enqueue[T any](queue Queue[T], values ...T) {
	for _, value := range values {
		queue.Enqueue(value)
	}
}

func Broadcast[T any](queues []Queue[T], values ...T) {
	for _, q := range queues {
		Enqueue(q, values...)
	}
}

func Dequeue[T any](queue Queue[T]) (value T, ok bool) {
	return queue.Dequeue()
}

func Drain[T any](q Queue[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			v, ok := q.Dequeue()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

func Peek[T any](queue Queue[T]) (value T, ok bool) {
	return queue.Peek()
}
