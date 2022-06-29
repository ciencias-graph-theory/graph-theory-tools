package graph

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A StaticGraph represents an undirected graph, modelled by its adjacency
// matrix, its adjacency list, or both. Usually, only one of the two
// representations will be used, and the other one will be calculated when an
// action is done more efficiently on it. The adjacency matrix is a
// two-dimensional byte array, and the adjacency list is a two-dimensional
// integer array.
type StaticGraph struct {
	matrix         [][]byte
	list           [][]int
	degreeSequence []int
}

// NewMatrixGraph initializes a graph modelled by its adjacency matrix. This
// method checks whether the matrix received as argument is symmetric; if it is
// not symmetric, it throws an error.
func NewGraphFromMatrix(matrix [][]byte) (*StaticGraph, error) {
	for i, v := range matrix {
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, assymetricMatrixError
			}
		}
	}
	return &StaticGraph{
		matrix:         matrix,
		list:           nil,
		degreeSequence: nil,
	}, nil
}

// NewMatrixGraphU initializes a graph modelled by its adjacency matrix. This is
// an unsafe method, it does not check whether the matrix received as argument
// is symmetric.
func NewGraphFromMatrixU(adjacency [][]byte) *StaticGraph {
	return &StaticGraph{
		matrix:         adjacency,
		list:           nil,
		degreeSequence: nil,
	}
}

// Matrix returns the adjacency matrix of the graph.
func (g *StaticGraph) Matrix() [][]byte {
	return g.matrix
}

// List returns the adjacency lisy of the graph.
func (g *StaticGraph) List() [][]int {
	return g.list
}

// Order returns the number of vertices in the graph.
func (g *StaticGraph) Order() int {
	return len(g.matrix)
}

// DegreeSequence returns the degree sequence of the graph
// in non-increasing order.
func (g *StaticGraph) DegreeSequence() []int {
	if g.degreeSequence != nil {
		return g.degreeSequence
	} else {
		degreeSequence := make([]int, len(g.matrix))
		for i, v := range g.matrix {
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
func (g *StaticGraph) Size() int {
	if g.degreeSequence != nil {
		return sliceutils.SumIntSlice(g.degreeSequence) / 2
	}
	size := 0
	degreeSequence := make([]int, len(g.matrix))
	for i, v := range g.matrix {
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
