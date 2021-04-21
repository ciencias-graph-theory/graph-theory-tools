package graph

import (
	"github.com/Japodrilo/graph-theory-tools/pkg/sliceutils"
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
