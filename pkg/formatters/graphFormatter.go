package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// Given a graph G. We build a vector by traveling the upper triangle of the
// adjacency matrix of G column by column without considering the diagonal.
func obtainUpperTriangle(matrix graph.AdjacencyMatrix) []byte {

	// Where n is the order of the Graph.
	n := len(matrix)

	// We build a vector of size (n * (n-1))/ 2); the amount of elements in the
	// upper triangle.
	v := make([]byte, (n * (n - 1) / 2))

	k := 0
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			v[k] = matrix[i][j]
			k++
		}
	}

	return v
}
