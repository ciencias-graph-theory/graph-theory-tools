package graph

// A MatrixGraph represents a graph, modelled by its adjacency matrix.
// The adjacency matrix is a two-dimensional byte array.
type MatrixGraph struct {
	adjacency      [][]byte
	degreeSequence []int
}

// NewMatrixGraph initializes a graph modelled by its adjacency matrix.
func NewMatrixGraph(adjacency [][]byte) *MatrixGraph {
	return &MatrixGraph{
		adjacency: adjacency,
		degreeSequence: nil,
	}
}

// Adjacency returns the adjacency matris of the graph.
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
