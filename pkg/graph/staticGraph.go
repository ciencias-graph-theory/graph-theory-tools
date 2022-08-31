package graph

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A StaticGraph represents an undirected graph, modelled by its adjacency
// matrix, its adjacency list, or both. Usually, only one of the two
// representations will be used, and the other one will be calculated when an
// action is done more efficiently on it. The adjacency matrix is a
// two-dimensional byte array, and the adjacency list is a two-dimensional
// integer array.
// A static graph cannot be modified (neither vertices nor edges can be added
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

// NewFromMatrix initializes a graph modelled by its adjacency matrix.
// This is an unsafe method, it does not check whether the matrix received as
// argument is symmetric.
func NewFromMatrix(adjacency AdjacencyMatrix) *StaticGraph {
	return &StaticGraph{
		matrix:         adjacency,
		list:           nil,
		degreeSequence: nil,
	}
}

// Gets an efficient adjacency list (list of sets) from a given adjacency list
// (list of lists), useful for testing membership.
func efficientAdjacencyList(list AdjacencyList) *EfficientAdjacencyList {
	efficientAdjacencyList := new(EfficientAdjacencyList)
	for _, v := range list {
		s := set.NewIntSet()
		for _, w := range v {
			s.Add(w)
		}
		*efficientAdjacencyList = append(*efficientAdjacencyList, *s)
	}
	return efficientAdjacencyList
}

// NewGraphFromList initializes a graph modelled by its adjacency list. This
// method checks whether the list received as argument is valid; if it is
// not valid, it throws an error.
func NewGraphFromList(list AdjacencyList) (*StaticGraph, error) {
	efficientList := efficientAdjacencyList(list)
	for i, v := range *efficientList {
		for w := range v.Items() {
			if !(*efficientList)[w].Contains(i) {
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

// NewFromList initializes a graph modelled by its adjacency list.
// This is an unsafe method, it does not check whether the list received as
// argument is valid.
func NewFromList(list AdjacencyList) *StaticGraph {
	return &StaticGraph{
		matrix:         nil,
		list:           list,
		degreeSequence: nil,
	}
}

// Matrix returns the adjacency matrix of the graph.
func (g *StaticGraph) Matrix() (AdjacencyMatrix, error) {
	if g.matrix == nil {
		return nil, NilAdjacencyMatrix
	}
	return g.matrix, nil
}

// List returns the adjacency list of the graph.
func (g *StaticGraph) List() (AdjacencyList, error) {
	if g.list == nil {
		return nil, NilAdjacencyList
	}
	return g.list, nil
}

// Transforms an adjacency matrix into an adjacency list.
func matrixToList(matrix AdjacencyMatrix) (*AdjacencyList, error) {
	list := new(AdjacencyList)
	for i, v := range matrix {
		var s []int
		for j, w := range v {
			if i < j && w != matrix[j][i] {
				return nil, assymetricMatrixError
			} else if w == 1 {
				s = append(s, j)
			}
		}
		*list = append(*list, s)
	}
	return list, nil
}

// Transforms an adjacency list into an adjacency matrix.
func listToMatrix(list AdjacencyList) (*AdjacencyMatrix, error) {
	matrix := make([][]byte, len(list))
	for i := range matrix {
		matrix[i] = make([]byte, len(list))
	}
	efficientList := efficientAdjacencyList(list)
	for i, v := range *efficientList {
		for w := range v.Items() {
			if !(*efficientList)[w].Contains(i) {
				return nil, invalidListError
			} else {
				matrix[i][w] = 1
			}
		}
	}
	return &matrix, nil
}

// Order returns the number of vertices in the graph.
func (g *StaticGraph) Order() int {
	return len(g.matrix)
	//TODO list case
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
