package stack

import "github.com/maxclav/ds/container"

// Stack is last in, first out: Pop returns the value pushed most recently.
// All yields the values from the top down.
type Stack[T any] interface {
	container.Collection[T]
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
