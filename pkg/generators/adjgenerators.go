package generators

import (
	"github.com/Japodrilo/graph-theory-tools/pkg/graph"
)

// IsCompleteMatrixGraph checks whether a graph is a complete graph or not. A
// complete graph is a loopless graph where every pair or different vertices is
// adjacent.
func IsCompleteMatrixGraph(g *graph.MatrixGraph) bool {
	a := g.Adjacency()
	for i := range a {
		for j := range a[i] {
			if i == j && a[i][j] != 0 {
				return false
			} else if i != j && a[i][j] != 1 {
				return false
			}
		}
	}
	return true
}

// CompleteMatrixGraph returns a complete graph of order n. A complete graph is
// a loopless graph where every pair or different vertices is adjacent.
func CompleteMatrixGraph(n int) *graph.MatrixGraph {
	a := make([][]byte, n, n)
	for i := range a {
		a[i] = make([]byte, n)
		for j := range a[i] {
			if i != j {
				a[i][j] = 1
			}
		}
	}
	return graph.NewMatrixGraph(a)
}
