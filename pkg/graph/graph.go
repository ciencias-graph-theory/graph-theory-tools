// Package graph provides the basic functionalities for creating
// and handling graphs.
package graph

// Type aliases to improve code readability.
type AdjacencyMatrix = [][]byte
type AdjacencyList = [][]int

type Graph interface {
	// Order returns the number of vertices in the graph.
	Order() int

	// DegreeSequence returns the degree sequence of the graph.
	DegreeSequence() []int

	// Size returns the size of the graph.
	Size() int

	// Matrix returns the adjacency matrix of the graph.
	Matrix() AdjacencyMatrix

	// Matrix returns the adjacency list of the graph.
	List() AdjacencyList
}
