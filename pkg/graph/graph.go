// Package graph provides the basic functionalities for creating
// and handling graphs.
package graph

type Graph interface {
	// Order returns the number of vertices in the graph.
	Order() int

	// DegreeSequence returns the degree sequence of the graph.
	DegreeSequence() []int

	// Size returns the size of the graph.
	Size() int
}

// PropertyCheck is an interface for checking a specific property in an
// induced subgraph.
type PropertyCheck interface {

	// Act checks whether a property for an induced subgraph is true.
	Check(g Graph, s []bool) bool
}

// PropertyCheck is an interface for checking a specific property in an
// induced subgraph.
type PropertyCheck interface {

	// Act checks whether a property for an induced subgraph is true.
	Check(g Graph, s []bool) bool
}
