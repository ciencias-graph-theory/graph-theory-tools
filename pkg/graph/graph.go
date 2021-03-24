// Package graph provides the basic functionalities for creating
// and handling graphs.
package graph

type Graph interface {
	// Order returns the number of vertices in the graph.
	Order() int

	// DegreeSequence returns the degree sequence of the graph.
	DegreeSequence() []int
}
