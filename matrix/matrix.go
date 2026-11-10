package matrix

import "fmt"

// Matrix is a grid of values addressed by row and column.
type Matrix[T any] interface {
	fmt.Stringer
	At(row, col int) T
	Set(row, col int, value T)
	Rows() int
	Cols() int
}
