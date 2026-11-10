package queue

import (
	"iter"
	"slices"

	"github.com/maxclav/ds/container"
)

var _ Queue[string] = (*ArrayQueue[string])(nil)

type ArrayQueue[T any] struct {
	values []T
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{}
}

func (q *ArrayQueue[T]) Enqueue(value T) {
	q.values = append(q.values, value)
}

func (q *ArrayQueue[T]) Dequeue() (value T, ok bool) {
	// TODO
	var t T
	return t, false
}

func (q *ArrayQueue[T]) Peek() (value T, ok bool) {
	if len(q.values) == 0 {
		return value, false
	}
	return q.values[0], true
}

func (q *ArrayQueue[T]) Len() int {
	return len(q.values)
}

func (q *ArrayQueue[T]) Clear() {
	q.values = nil
}

func (q *ArrayQueue[T]) All() iter.Seq[T] {
	return slices.Values(q.values)
}

func (q *ArrayQueue[T]) String() string {
	return container.String(q.All())
}
