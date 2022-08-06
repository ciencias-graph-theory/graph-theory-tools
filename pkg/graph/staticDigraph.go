package graph

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A MatrixDigraph represents a digraph modelled by its adjacency
// matrix. The adjacency matrix is a two-dimensional byte array.
// Bytes were chosen for future compatibility with weighted
// digraphs, but in the case of a simple digraph, entries will be
// either 1 or 0.
type StaticDigraph struct {
	*StaticGraph
	indegreeSequence  []int
	outdegreeSequence []int
}

// NewDigraphFromMatrix initializes a digraph modelled by its adjacency matrix.
func NewDigraphFromMatrix(matrix AdjacencyMatrix) *StaticDigraph {
	return &StaticDigraph{
		StaticGraph:       NewFromMatrix(matrix),
		indegreeSequence:  nil,
		outdegreeSequence: nil,
	}
}

// NewDigraphFromList initializes a digraph modelled by its adjacency list.
func NewDigraphFromList(list AdjacencyList) *StaticDigraph {
	return &StaticDigraph{
		StaticGraph:       NewFromList(list),
		indegreeSequence:  nil,
		outdegreeSequence: nil,
	}
}

// Computes in-degree and out-degree sequences of the digraph
func (d *StaticDigraph) computeDegreeSequences() {
	inSequence := make([]int, len(d.matrix))
	outSequence := make([]int, len(d.matrix))
	for i, v := range d.matrix {
		for j, n := range v {
			if n != 0 {
				outSequence[i]++
				inSequence[j]++
			}
		}
	}
	d.indegreeSequence = inSequence
	d.outdegreeSequence = outSequence
}

// DegreeSequence returns the degree sequence of the digraph.
// The degree sequence of the digraph is the sum of the in-degree
// sequence and the out-degree sequence.
func (d *StaticDigraph) DegreeSequence() []int {
	if d.indegreeSequence == nil || d.outdegreeSequence == nil {
		d.computeDegreeSequences()
	}
	degreeSequence := make([]int, len(d.matrix))
	for i := 0; i < len(d.matrix); i++ {
		degreeSequence[i] = d.indegreeSequence[i] + d.outdegreeSequence[i]
	}
	return degreeSequence
}

// IndegreeSequence returns the in-degree sequence of the digraph
// in non-increasing order.
func (d *StaticDigraph) IndegreeSequence() []int {
	if d.indegreeSequence == nil {
		d.computeDegreeSequences()
	}
	return d.indegreeSequence
}

// OutdegreeSequence returns the out-degree sequence of the
// digraph in non-increasing order.
func (d *StaticDigraph) OutdegreeSequence() []int {
	if d.outdegreeSequence == nil {
		d.computeDegreeSequences()
	}
	return d.outdegreeSequence
}

// Size returns the size (number of arcs) of a digraph.
func (d *StaticDigraph) Size() int {
	if d.indegreeSequence == nil || d.outdegreeSequence == nil {
		size := 0
		inSequence := make([]int, len(d.matrix))
		outSequence := make([]int, len(d.matrix))
		for i, v := range d.matrix {
			for j, n := range v {
				if n != 0 {
					outSequence[i]++
					inSequence[j]++
					size++
				}
			}
		}
		d.indegreeSequence = inSequence
		d.outdegreeSequence = outSequence
		return size
	} else {
		return (sliceutils.SumIntSlice(d.indegreeSequence) +
			sliceutils.SumIntSlice(d.outdegreeSequence)) / 2
	}
}
