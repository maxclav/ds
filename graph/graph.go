package graph

import (
	"iter"

	"github.com/maxclav/ds/container"
)

// Graph is a set of vertices, numbered from 0, joined by weighted directed edges.
// For an undirected graph, add each edge in both directions. Len counts vertices.
type Graph[W any] interface {
	container.Container
	// AddVertex adds a vertex and returns its number.
	AddVertex() int
	AddEdge(from, to int, weight W)
	// RemoveEdge removes the edge and reports whether it was present.
	RemoveEdge(from, to int) bool
	HasEdge(from, to int) bool
	// Neighbors yields each vertex one edge away from v, with that edge's weight.
	Neighbors(v int) iter.Seq2[int, W]
}
