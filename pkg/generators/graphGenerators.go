package generators

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

// IsCompleteMatrixGraph checks whether a graph is a complete graph or not. A
// complete graph is a loopless graph where every pair of different vertices is
// adjacent.
func IsCompleteMatrixGraph(g *graph.StaticGraph) bool {
	a := g.Matrix()
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
}

// CompleteMatrixGraph returns a complete graph of order n. A complete graph is
// a loopless graph where every pair of different vertices is adjacent.
func CompleteMatrixGraph(n int) *graph.StaticGraph {
	a := make([][]byte, n, n)
	for i := range a {
		a[i] = make([]byte, n)
		for j := range a[i] {
			if i != j {
				a[i][j] = 1
			}
		}
	}
	return graph.NewGraphFromMatrixU(a)
}

// IsCompleteBipartiteMatrixGraph checks whether a graph is a complete bipartite
// graph or not.
func IsCompleteBipartiteMatrixGraph(g *graph.StaticGraph) bool {
	a := g.Matrix()
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
}

// CompleteBipartiteMatrixGraph returns a complete bipartite graph with parts of
// cardinality n and m.
func CompleteBipartiteMatrixGraph(n, m int) *graph.StaticGraph {
	a := make([][]byte, n+m, n+m)
	for i := range a {
		a[i] = make([]byte, n+m)
		for j := range a[i] {
			if i != j && ((i < n && n-1 < j) || (n-1 < i && j < n)) {
				a[i][j] = 1
			}
		}
	}
	return graph.NewGraphFromMatrixU(a)
}

// IsCycleMatrixGraph checks whether a graph is an irreflexive cycle or not.
func IsCycleMatrixGraph(g *graph.StaticGraph) bool {
	d := g.DegreeSequence()
	for _, v := range d {
		if v != 2 {
			return false
		}
	}
	a := g.Matrix()
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
}

/*
// CycleMatrixGraph returns an irreflexive cycle of order n in canonical order.
func CycleMatrixGraph(n int) *graph.StaticGraph {
	c := PathMatrixGraph(n)
	err := c.AddEdge(0, n-1)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
*/

// IsPathMatrixGraph checks whether a graph is an irreflexive path or not.
func IsPathMatrixGraph(g *graph.StaticGraph) bool {
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
	a := g.Matrix()
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
}

// PathMatrixGraph returns an irreflexive path of order n in canonical order.
func PathMatrixGraph(n int) *graph.StaticGraph {
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
	return graph.NewGraphFromMatrixU(a)
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
