// Package graph provides the basic functionalities for creating
// and handling graphs.
package graph

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
)

// Type aliases to improve code readability.
type AdjacencyMatrix = [][]byte
type AdjacencyList = [][]int
type EfficientAdjacencyList = []set.IntSet

type Graph interface {
	// Order returns the number of vertices in the graph.
	Order() int

	// DegreeSequence returns the degree sequence of the graph.
	DegreeSequence() []int

	// Size returns the size of the graph.
	Size() int

	// Matrix returns the adjacency matrix of the graph.
	Matrix() (AdjacencyMatrix, error)

	// Matrix returns the adjacency list of the graph.
	List() (AdjacencyList, error)

	// Neighbours returns a set of the neighbours to a given vertex in the graph.
	Neighbours(v int) *set.IntSet
}
