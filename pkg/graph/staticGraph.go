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
// An static graph cannot be modified (neither vertices nor edges can be added
// to it).
type StaticGraph struct {
	matrix         AdjacencyMatrix
	list           AdjacencyList
	degreeSequence []int
}

// NewGraphFromMatrix initializes a graph modelled by its adjacency matrix. This
// method checks whether the matrix received as argument is symmetric; if it is
// not symmetric, it throws an error.
func NewGraphFromMatrix(matrix AdjacencyMatrix) (*StaticGraph, error) {
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

// NewGraphFromMatrixU initializes a graph modelled by its adjacency matrix.
// This is an unsafe method, it does not check whether the matrix received as
// argument is symmetric.
func NewGraphFromMatrixU(adjacency AdjacencyMatrix) *StaticGraph {
	return &StaticGraph{
		matrix:         adjacency,
		list:           nil,
		degreeSequence: nil,
	}
}

// Auxiliar method. Gets an efficient adjacency list (list of sets) from a given
// adjacency list (list of lists), useful for testing membership.
func efficientAdjacencyList(list AdjacencyList) *[]map[int]struct{} {
	efficientAdjacencyList := new([]map[int]struct{})
	for _, v := range list {
		s := make(map[int]struct{})
		for _, w := range v {
			s[w] = struct{}{}
		}
		*efficientAdjacencyList = append(*efficientAdjacencyList, s)
	}
	return efficientAdjacencyList
}

// NewGraphFromList initializes a graph modelled by its adjacency list. This
// method checks whether the list received as argument is valid; if it is
// not valid, it throws an error.
func NewGraphFromList(list AdjacencyList) (*StaticGraph, error) {
	efficientList := efficientAdjacencyList(list)
	for i, v := range *efficientList {
		for w := range v {
			_, contains := (*efficientList)[w][i]
			if !contains {
				return nil, invalidListError
			}
		}
	}
	return &StaticGraph{
		matrix:         nil,
		list:           list,
		degreeSequence: nil,
	}, nil
}

// NewGraphFromListU initializes a graph modelled by its adjacency list.
// This is an unsafe method, it does not check whether the list received as
// argument is valid.
func NewGraphFromListU(list AdjacencyList) *StaticGraph {
	return &StaticGraph{
		matrix:         nil,
		list:           list,
		degreeSequence: nil,
	}
}

// Matrix returns the adjacency matrix of the graph.
func (g *StaticGraph) Matrix() AdjacencyMatrix {
	return g.matrix
}

// List returns the adjacency list of the graph.
func (g *StaticGraph) List() AdjacencyList {
	return g.list
}

// === Tentative auxiliar function ====
func matrixToList(matrix AdjacencyList) (AdjacencyList, error) {
	return nil, nil
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

// Size returns the size (number of edges) of a graph.
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
