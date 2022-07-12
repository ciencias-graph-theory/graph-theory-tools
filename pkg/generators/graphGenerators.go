package generators

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// IsComplete checks whether a graph/digraph is a complete graph/digraph or not.
func IsComplete(g Graph) bool {
	if a, err := g.Matrix(); err != graph.NilAdjacencyMatrix {
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
	} else if _, err := g.List(); err != graph.NilAdjacencyList {
		// TODO: deal with adjacency list case
		return false
	} else {
		return false
	}
}

// Makes a complete adjacency matrix of order n.
func completeMatrix(n int) graph.AdjacencyMatrix {
	a := make([][]byte, n, n)
	for i := range a {
		a[i] = make([]byte, n)
		for j := range a[i] {
			if i != j {
				a[i][j] = 1
			}
		}
	}
	return a
}

// CompleteMatrixGraph returns a complete graph of order n modelled by an
// adjacency matrix. A complete graph is a loopless graph where every pair of
// different vertices is adjacent.
func CompleteMatrixGraph(n int) *StaticGraph {
	return graph.NewFromMatrix(completeMatrix(n))
}

// CompleteMatrixDigraph returns a complete digraph of order n modelled by an
// adjacency matrix. A complete digraph is a digraph where every pair of
// different vertices is adjacent.
func CompleteMatrixDigraph(n int) *StaticDigraph {
	return graph.NewDigraphFromMatrix(completeMatrix(n))
}

// CompleteListGraph returns a complete graph of order n modelled by an
// adjacency list.
func CompleteListGraph(n int) *StaticGraph {
	// TODO: implement
	return nil
}

// CompleteListDigraph returns a complete digraph of order n modelled by an
// adjacency list.
func CompleteListDigraph(n int) *StaticDigraph {
	// TODO: implement
	return nil
}

// IsCompleteBipartite checks whether a graph/digraph is a complete bipartite
// graph/digraph or not.
func IsCompleteBipartite(g Graph) bool {
	if a, err := g.Matrix(); err != graph.NilAdjacencyMatrix {
		x := make([]int, 0)
		y := make([]int, 0)
		for i := range a[0] {
			if a[0][i] == 0 {
				x = append(x, i)
			} else if a[0][i] == 1 {
				y = append(y, i)
			} else {
				return false
			}
		}
		for _, v := range x {
			for _, w := range y {
				if a[v][w] != 1 || a[w][v] != 1 {
					return false
				}
			}
			for _, w := range x {
				if a[v][w] != 0 {
					return false
				}
			}
		}
		for _, v := range y {
			for _, w := range y {
				if a[v][w] != 0 {
					return false
				}
			}
		}
		return true
	} else if _, err := g.List(); err != graph.NilAdjacencyList {
		// TODO: deal with adjacency list case
		return false
	} else {
		return false
	}
}

// Makes a complete bipartite adjacency matrix of order n+m.
func completeBipartiteMatrix(n, m int) graph.AdjacencyMatrix {
	a := make([][]byte, n+m, n+m)
	for i := range a {
		a[i] = make([]byte, n+m)
		for j := range a[i] {
			if i != j && ((i < n && n-1 < j) || (n-1 < i && j < n)) {
				a[i][j] = 1
			}
		}
	}
	return a
}

// CompleteBipartiteMatrixGraph returns a complete bipartite graph modelled by
// an adjacency matrix, with parts of cardinality n and m.
func CompleteBipartiteMatrixGraph(n, m int) *StaticGraph {
	return graph.NewFromMatrix(completeBipartiteMatrix(n, m))
}

// CompleteBipartiteMatrixDigraph returns a complete bipartite digraph modelled
// by an adjacency matrix, with parts of cardinality n and m.
func CompleteBipartiteMatrixDigraph(n, m int) *StaticDigraph {
	return graph.NewDigraphFromMatrix(completeBipartiteMatrix(n, m))
}

// CompleteBipartiteListGraph returns a complete bipartite graph modelled by an
// adjacency list, with parts of cardinality n and m.
func CompleteBipartiteListGraph(n, m int) *StaticGraph {
	// TODO: implement
	return nil
}

// CompleteBipartiteListDigraph returns a complete bipartite digraph modelled by an
// adjacency list, with parts of cardinality n and m.
func CompleteBipartiteListDigraph(n, m int) *StaticDigraph {
	// TODO: implement
	return nil
}

// IsCycle checks whether a graph is an irreflexive cycle or not.
func IsCycle(g *StaticGraph) bool {
	d := g.DegreeSequence()
	for _, v := range d {
		if v != 2 {
			return false
		}
	}
	if a, err := g.Matrix(); err != graph.NilAdjacencyMatrix {
		n := len(d)
		i := 0
		j := sliceutils.NextNonZero(a[i], 0)
		k := j
		l := sliceutils.NextNonZero(a[k], i)
		var t int
		m := 0
		for k != i && l != j && m < n {
			t = k
			k = l
			l = sliceutils.NextNonZero(a[k], t)
			m++
		}
		return m+1 == n
	} else if _, err := g.List(); err != graph.NilAdjacencyMatrix {
		// TODO: deal with adjacency list case
		return false
	} else {
		return false
	}
}

// IsDirectedCycle checks whether a digraph is an irreflexive directed cycle
// or not.
func IsDirectedCycle(d StaticDigraph) bool {
	// TODO: implement
	return false
}

// MatrixCycle returns an irreflexive cycle of order n in canonical order.
func MatrixCycle(n int) *graph.StaticGraph {
	c := matrixPath(n)
	if n > 1 {
		c[0][n-1] = 1
		c[n-1][0] = 1
	}
	return graph.NewFromMatrix(c)
}

// IsPath checks whether a graph is an irreflexive path or not.
func IsPath(g *StaticGraph) bool {
	d := g.DegreeSequence()
	if len(d) == 1 && d[0] == 0 {
		return true
	}
	start := -1
	end := -1
	for i, v := range d {
		if v == 1 {
			if start == -1 {
				start = i
			} else if end == -1 {
				end = i
			} else {
				return false
			}
		} else if v != 2 {
			return false
		}
	}
	if end == -1 {
		return false
	}
	if a, err := g.Matrix(); err != graph.NilAdjacencyMatrix {
		n := len(d)
		i := start
		var j int
		if a[i][0] == 1 {
			j = 0
		} else {
			j = sliceutils.NextNonZero(a[start], 0)
		}
		var t int
		m := 0
		for i != end && m < n {
			t = i
			i = j
			j = sliceutils.NextNonZero(a[i], t)
			m++
		}
		return m+1 == n
	} else if _, err := g.List(); err != graph.NilAdjacencyMatrix {
		// TODO: deal with adjacency list case
		return false
	} else {
		return false
	}
}

// IsDirectedPath checks whether a graph is an irreflexive directed path or not.
func IsDirectedPath(g *StaticDigraph) bool {
	// TODO: implement
	return false
}

func matrixPath(n int) graph.AdjacencyMatrix {
	a := make([][]byte, n, n)
	a[0] = make([]byte, n, n)
	if n > 1 {
		a[0][1] = 1
		a[n-1] = make([]byte, n, n)
		a[n-1][n-2] = 1
		for i := 1; i < n-1; i++ {
			a[i] = make([]byte, n, n)
			a[i][i-1] = 1
			a[i][i+1] = 1
		}
	}
	return a
}

// Path returns an irreflexive path of order n in canonical order.
func MatrixPath(n int) *StaticGraph {
	return graph.NewFromMatrix(matrixPath(n))
}

// DirectedPath returns an irreflexive dircted path of order n in canonical order.
func MatrixDirectedPath(n int) *StaticDigraph {
	a := make([][]byte, n, n)
	a[0] = make([]byte, n, n)
	if n > 1 {
		a[0][1] = 1
		a[n-1] = make([]byte, n, n)
		for i := 1; i < n-1; i++ {
			a[i] = make([]byte, n, n)
			a[i][i+1] = 1
		}
	}
	return graph.NewDigraphFromMatrix(a)
}

/*
// CirculantMatrixDigraph returns a circulant digraph of order n
// with set of integer jumps s.
func CirculantMatrixDigraph(n int, jumps map[int]bool) *graph.MatrixDigraph {
	a := make([][]byte, n, n)
	var k int
	for i := range a {
		a[i] = make([]byte, n)
		for j := range jumps {
			if 0 < j && j < n {
				k = (i + j) % n
				a[i][k] = 1
			} else if j < 0 && -n < j {
				k = (i + n + j) % n
				a[i][k] = 1
			}
		}
	}
	return graph.NewMatrixDigraph(a)
}
*/
