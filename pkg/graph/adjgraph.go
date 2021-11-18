package graph

import (
	"errors"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A MatrixGraph represents a graph, modelled by its adjacency matrix.
// The adjacency matrix is a two-dimensional byte array.
type MatrixGraph struct {
	adjacency      [][]byte
	degreeSequence []int
}

// NewMatrixGraph initializes a graph modelled by its adjacency matrix.
func NewMatrixGraph(adjacency [][]byte) *MatrixGraph {
	return &MatrixGraph{
		adjacency:      adjacency,
		degreeSequence: nil,
	}
}

// Adjacency returns the adjacency matrix of the graph.
func (g *MatrixGraph) Adjacency() [][]byte {
	return g.adjacency
}

// Order returns the number of vertices in the graph.
func (g *MatrixGraph) Order() int {
	return len(g.adjacency)
}

// DegreeSequence returns the degree sequence of the graph
// in non-increasing order.
func (g *MatrixGraph) DegreeSequence() []int {
	if g.degreeSequence != nil {
		return g.degreeSequence
	} else {
		degreeSequence := make([]int, len(g.adjacency))
		for i, v := range g.adjacency {
			for _, n := range v {
				if n == 1 {
					degreeSequence[i] += 1
				}
			}
		}
		g.degreeSequence = degreeSequence
		return degreeSequence
	}
}

// AddEdge adds the edge between two given vertices. When the input does not
// correspond to two (possibly equal) vertices of the graph, an error is
// returned.
func (g *MatrixGraph) AddEdge(u, v int) error {
	o := g.Order()
	if u < 0 || v < 0 || o-1 < u || o-1 < v {
		return errors.New("not a valid vertex pair")
	}
	g.adjacency[u][v] = 1
	g.adjacency[v][u] = 1
	return nil
}

// Size returns the size (number of edges) of a matrix graph.
func (g *MatrixGraph) Size() int {
	if g.degreeSequence != nil {
		return sliceutils.SumIntSlice(g.degreeSequence) / 2
	}
	size := 0
	degreeSequence := make([]int, len(g.adjacency))
	for i, v := range g.adjacency {
		for j := 0; j < i+1; j++ {
			if v[j] != 0 {
				degreeSequence[i]++
				degreeSequence[j]++
				size++
			}
		}
	}
	g.degreeSequence = degreeSequence
	return size
}
