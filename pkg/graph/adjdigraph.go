package graph

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// A MatrixDigraph represents a digraph modelled by its adjacency
// matrix. The adjacency matrix is a two-dimensional byte array.
// Bytes were chosen for future compatibility with weighted
// digraphs, but in the case of a simple digraph, entries will be
// either 1 or 0.
type MatrixDigraph struct {
	adjacency         [][]byte
	indegreeSequence  []int
	outdegreeSequence []int
}

// NewMatrixDigraph initializes a digraph modelled by its
// adjacency matrix.
func NewMatrixDigraph(adjacency [][]byte) *MatrixDigraph {
	return &MatrixDigraph{
		adjacency:         adjacency,
		indegreeSequence:  nil,
		outdegreeSequence: nil,
	}
}

// Adjacency returns the adjacency matrix of the digraph.
func (d *MatrixDigraph) Adjacency() [][]byte {
	return d.adjacency
}

// Order returns the number of vertices in the digraph.
func (d *MatrixDigraph) Order() int {
	return len(d.adjacency)
}

// Computes in-degree and out-degree sequences of the digraph
func (d *MatrixDigraph) computeDegreeSequences() {
	inSequence := make([]int, len(d.adjacency))
	outSequence := make([]int, len(d.adjacency))
	for i, v := range d.adjacency {
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
func (d *MatrixDigraph) DegreeSequence() []int {
	if d.indegreeSequence == nil ||
		d.outdegreeSequence == nil {
		d.computeDegreeSequences()
	}
	degreeSequence := make([]int, len(d.adjacency))
	for i := 0; i < len(d.adjacency); i++ {
		degreeSequence[i] = d.indegreeSequence[i] +
			d.outdegreeSequence[i]
	}
	return degreeSequence
}

// IndegreeSequence returns the in-degree sequence of the digraph
// in non-increasing order.
func (d *MatrixDigraph) IndegreeSequence() []int {
	if d.indegreeSequence == nil {
		d.computeDegreeSequences()
	}
	return d.indegreeSequence
}

// OutdegreeSequence returns the out-degree sequence of the
// digraph in non-increasing order.
func (d *MatrixDigraph) OutdegreeSequence() []int {
	if d.outdegreeSequence == nil {
		d.computeDegreeSequences()
	}
	return d.outdegreeSequence
}

// Size computes the size (number of arcs) of the digraph.
func (d *MatrixDigraph) Size() int {
	if d.indegreeSequence == nil || d.outdegreeSequence == nil {
		size := 0
		inSequence := make([]int, len(d.adjacency))
		outSequence := make([]int, len(d.adjacency))
		for i, v := range d.adjacency {
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
