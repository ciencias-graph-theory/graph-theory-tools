package graph

import (
	"errors"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A Digraph represents a directed graph, modelled by its adjacency matrix, its
// adjacency list, or both. Usually, only one of the two representations will
// be used, and the other one will be calculated when an action is done more
// efficiently on it. The adjacency matrix is a two-dimensional byte array, and
// the adjacency list is a two-dimensional integer array. Since operations for
// adding and deleting edges are supported which only modify the adjacency
// matrix, a boolean variable is used to indicate if the adjacency list is
// up-to-date with respect to the adjacency matrix. A boolean variable is also
// used to indicate whether the digraph is a graph (the adjacency matrix should
// be symmetric).
type Digraph struct {
	matrix         [][]byte
	list           [][]int
	updatedEdges   bool
	isGraph        bool
	degreeSequence []int
}

// NewMatrixGraph initializes a graph modelled by its adjacency matrix. This
// method checks whether the matrix received as argument is symmetric; if it is
// not symmetric, it throws an error.
func NewMatrixGraph(matrix [][]byte) (*Digraph, error) {
	for i, v := range matrix {
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, errors.New("adjacency matrix is not symmetric")
			}
		}
	}
	return &Digraph{
		matrix:         matrix,
		list:           nil,
		updatedEdges:   true,
		isGraph:        true,
		degreeSequence: nil,
	}, nil
}

// NewMatrixGraphU initializes a graph modelled by its adjacency matrix. This is
// an unsafe method, it does not check whether the matrix received as argument
// is symmetric.
func NewMatrixGraphU(adjacency [][]byte) *Digraph {
	return &Digraph{
		matrix:         adjacency,
		list:           nil,
		updatedEdges:   true,
		isGraph:        true,
		degreeSequence: nil,
	}
}

// Matrix returns the adjacency matrix of the graph.
func (g *Digraph) Matrix() [][]byte {
	return g.matrix
}

// Order returns the number of vertices in the graph.
func (g *Digraph) Order() int {
	return len(g.matrix)
}

// DegreeSequence returns the degree sequence of the graph
// in non-increasing order.
func (g *Digraph) DegreeSequence() []int {
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

// AddEdge adds the edge between two given vertices. When the input does not
// correspond to two (possibly equal) vertices of the graph, an error is
// returned.
func (g *Digraph) AddEdge(u, v int) error {
	o := g.Order()
	if u < 0 || v < 0 || o-1 < u || o-1 < v {
		return errors.New("not a valid vertex pair")
	}
	g.matrix[u][v] = 1
	g.matrix[v][u] = 1
	return nil
}

// Size returns the size (number of edges) of a matrix graph.
func (g *Digraph) Size() int {
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
