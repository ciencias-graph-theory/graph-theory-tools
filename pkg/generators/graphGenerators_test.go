package generators

import (
	"math/rand"
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

func TestIsComplete(t *testing.T) {
	a := [][]byte{
		{0, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
	}
	b := [][]byte{
		{0, 1, 1},
		{1, 1, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	d := [][]byte{
		{0},
	}
	k, _ := graph.NewGraphFromMatrix(a)
	l, _ := graph.NewGraphFromMatrix(b)
	m := graph.NewFromMatrix(c)
	n, _ := graph.NewGraphFromMatrix(d)
	if !IsComplete(k) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsComplete(k),
		)
	}
	if IsComplete(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsComplete(l),
		)
	}
	if IsComplete(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsComplete(m),
		)
	}
	if !IsComplete(n) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsComplete(n),
		)
	}

	e := [][]int{
		{1, 2, 3, 4, 5},
		{0, 2, 3, 4, 5},
		{0, 1, 3, 4, 5},
		{0, 1, 2, 4, 5},
		{0, 1, 2, 3, 5},
		{0, 1, 2, 3, 4},
	}
	f := [][]int{
		{1, 2},
		{0, 1, 2},
		{0, 1},
	}
	g := [][]int{
		{2},
		{0, 2},
		{0, 1},
	}
	h := [][]int{
		{},
	}
	o, _ := graph.NewGraphFromList(e)
	p, _ := graph.NewGraphFromList(f)
	q := graph.NewFromList(g)
	r, _ := graph.NewGraphFromList(h)
	if !IsComplete(o) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsComplete(o),
		)
	}
	if IsComplete(p) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsComplete(p),
		)
	}
	if IsComplete(q) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsComplete(q),
		)
	}
	if !IsComplete(r) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsComplete(r),
		)
	}
}

// TestCompleteMatrix calls CompleteGraph with five different
// randomly generated graphs, and checks each of them to be a complete graph by
// exploring their adjacency matrices.
func TestCompleteMatrix(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		k := CompleteMatrixGraph(n)
		a, _ := k.Matrix()
		o := k.Order()
		if o != n {
			t.Errorf("Graph with incorrect order. Expected %v, got %v", n, o)
		}
		for i := range a {
			for j := range a[i] {
				if i == j && a[i][j] != 0 {
					t.Error("Graph is not irreflexive")
				} else if i != j && a[i][j] != 1 {
					t.Errorf("No adjacency between %v and %v", i, j)
				}
			}
		}
	}
}

func TestIsCompleteBipartite(t *testing.T) {
	a := [][]byte{
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 0, 0},
	}
	b := [][]byte{
		{0, 1, 1},
		{1, 1, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	d := [][]byte{
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0},
		{0, 1, 1, 1, 0, 0},
	}
	e := [][]byte{
		{0},
	}
	k, _ := graph.NewGraphFromMatrix(a)
	l, _ := graph.NewGraphFromMatrix(b)
	m := graph.NewFromMatrix(c)
	n := graph.NewFromMatrix(d)
	o, _ := graph.NewGraphFromMatrix(e)
	if !IsCompleteBipartite(k) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteBipartite(k),
		)
	}
	if IsCompleteBipartite(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartite(k),
		)
	}
	if IsCompleteBipartite(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartite(k),
		)
	}
	if IsCompleteBipartite(n) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartite(k),
		)
	}
	if !IsCompleteBipartite(o) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteBipartite(k),
		)
	}
}

// TestCompleteBipartiteMatrixGraph calls CompleteBipartiteMatrixGraph with five
// different randomly generated graphs, and checks each of them to be a complete
// bipartite graph by exploring their adjacency matrices.
func TestCompleteBipartiteGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		m := rand.Intn(1000)
		k := CompleteBipartiteMatrixGraph(n, m)
		a, _ := k.Matrix()
		for i := range a {
			for j := range a[i] {
				if i == j && a[i][j] != 0 {
					t.Error("Graph is not irreflexive")
				} else if i != j {
					if (i < n && n-1 < j && a[i][j] != 1) ||
						(n-1 < i && j < n && a[i][j] != 1) {
						t.Errorf("No adjacency between %v and %v", i, j)
					}
				} else if a[i][j] != 0 {
					t.Errorf("Unexpected adjacency between %v and %v", i, j)
				}
			}
		}
	}
}

// TestIsCycle calls IsCycleMatrixGraph with different hardcoded
// graphs, including cycles of different lengths, disconnected 2-regular graphs,
// and other non-cycle graphs.
func TestIsCycle(t *testing.T) {
	a := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	b := [][]byte{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
	}
	d := [][]byte{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
	}
	e := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	g, _ := graph.NewGraphFromMatrix(a)
	h, _ := graph.NewGraphFromMatrix(b)
	i, _ := graph.NewGraphFromMatrix(c)
	j, _ := graph.NewGraphFromMatrix(d)
	k, _ := graph.NewGraphFromMatrix(e)
	if !IsCycle(g) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycle(g),
		)
	}
	if !IsCycle(h) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycle(h),
		)
	}
	if !IsCycle(i) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycle(i),
		)
	}
	if IsCycle(j) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycle(j),
		)
	}
	if IsCycle(k) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycle(k),
		)
	}
}

// TestMatrixCycle calls MatrixCycle with five different
// randomly generated numbers, and checks each of them to be a cycle by
// exploring their adjacency matrices.
func TestMatrixCycle(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			c := MatrixCycle(n)
			d := c.DegreeSequence()
			for _, v := range d {
				if v != 2 {
					t.Error("The graph is not 2-regular")
				}
			}
			a, _ := c.Matrix()
			if a[0][1] != 1 || a[0][n-1] != 1 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					0,
				)
			}
			if a[n-1][0] != 1 || a[n-1][n-2] != 1 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					n-1,
				)
			}
			for i := 1; i < n-1; i++ {
				if a[i][i-1] != 1 || a[i][i+1] != 1 {
					t.Errorf(
						"Adjacencies of vertex %v are not as expected",
						i,
					)
				}
			}
		}
	}
}

// TestIsPath calls IsPath with different hardcoded
// graphs, including cycles of different lengths, linear forests,
// and other non-cycle graphs.
func TestIsPath(t *testing.T) {
	a := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	b := [][]byte{
		{0, 0, 1},
		{0, 0, 1},
		{1, 1, 0},
	}
	c := [][]byte{
		{0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0},
	}
	d := [][]byte{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0},
	}
	e := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
	}
	f := [][]byte{
		{0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0},
	}
	g := [][]byte{
		{0},
	}
	h, _ := graph.NewGraphFromMatrix(a)
	i, _ := graph.NewGraphFromMatrix(b)
	j, _ := graph.NewGraphFromMatrix(c)
	k, _ := graph.NewGraphFromMatrix(d)
	l, _ := graph.NewGraphFromMatrix(e)
	m, _ := graph.NewGraphFromMatrix(f)
	n, _ := graph.NewGraphFromMatrix(g)
	if !IsPath(h) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPath(h),
		)
	}
	if !IsPath(i) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPath(i),
		)
	}
	if !IsPath(j) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPath(j),
		)
	}
	if IsPath(k) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPath(k),
		)
	}
	if IsPath(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPath(l),
		)
	}
	if IsPath(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPath(m),
		)
	}
	if !IsPath(n) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPath(n),
		)
	}
}

// TestPathMatrixGraph calls PathMatrixGraph with five different
// randomly generated numbers, and checks each of them to be a path by
// exploring their adjacency matrices.
func TestPath(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			c := MatrixPath(n)
			d := c.DegreeSequence()
			for i, v := range d {
				if i == 0 || i == n-1 {
					if v != 1 {
						t.Error("First or last vertex does not have degree 1")
					}
				} else if v != 2 {
					t.Error("The graph is not 2-regular")
				}
			}
			a, _ := c.Matrix()
			if a[0][1] != 1 || a[0][n-1] != 0 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					0,
				)
			}
			if a[n-1][0] != 0 || a[n-1][n-2] != 1 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					n-1,
				)
			}
			for i := 1; i < n-1; i++ {
				if a[i][i-1] != 1 || a[i][i+1] != 1 {
					t.Errorf(
						"Adjacencies of vertex %v are not as expected",
						i,
					)
				}
			}
		}
	}
}

// TestPathDigraph calls PathDigraph with five different
// randomly generated numbers, and checks each of them to be a directed path by
// exploring their adjacency matrices.
func TestDirectedPath(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			p := MatrixDirectedPath(n)
			d := p.DegreeSequence()
			for i, v := range d {
				if i == 0 || i == n-1 {
					if v != 1 {
						t.Error("First or last vertex does not have degree 1")
					}
				} else if v != 2 {
					t.Error("The vertices do not have degree 2")
				}
			}
			o := p.OutdegreeSequence()
			for i, v := range o {
				if i == n-1 {
					if v != 0 {
						t.Error("Last vertex does not have outdegree 0")
					}
				} else if v != 1 {
					t.Error("The vertices do not have outdegree 1")
				}
			}
			a, _ := p.Matrix()
			if a[0][1] != 1 || a[0][n-1] != 0 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					0,
				)
			}
			if a[n-1][0] != 0 || a[n-1][n-2] != 0 {
				t.Errorf(
					"Adjacencies of vertex %v are not as expected",
					n-1,
				)
			}
			for i := 1; i < n-1; i++ {
				if a[i][i-1] != 0 || a[i][i+1] != 1 {
					t.Errorf(
						"Adjacencies of vertex %v are not as expected",
						i,
					)
				}
			}
		}
	}
}

/*

// TestCirculantMatrixDigraph randomly generates five circulant digraphs by
// constructing their adjacency matrices from a set of randomly generated
// jumps.   Then, it calls the CirculantMatrixDigraph function with the same set
// of jumps, and compares the obtained adjacency matrix with the previously
// generated one.   A second test is performed by hard coding the adjacency
// matrix of a circulant digraph, and testing for equality against the
// adjacency matrix of the circulant digraph generated by
// CirculantMatrixDigraph.
func TestCirculantMatrixDigraph(t *testing.T) {
	n := 0
	var a [][]byte
	var s int
	var k int
	var l int
	jumps := make(map[int]bool)
	var got *graph.MatrixDigraph
	for i := 0; i < 5; i++ {
		for n == 0 {
			n = rand.Intn(30)
		}
		a = make([][]byte, n, n)
		for j := range a {
			a[j] = make([]byte, n, n)
		}
		s = rand.Intn(n)
		for j := 0; j < s; j++ {
			k = rand.Intn(n)
			jumps[k] = true
		}
		for j := range a {
			for jump := range jumps {
				l = (j + jump) % n
				if l != j {
					a[j][l] = 1
				}
			}
		}
		got = CirculantMatrixDigraph(n, jumps)
		if !sliceutils.EqualByteMatrix(a, got.Matrix()) {
			t.Errorf("Expected %v, but got %v", a, got.Matrix())
		}
	}
	a = [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1, 0},
	}
	jumps = map[int]bool{
		1:  true,
		-1: false,
	}
	got = CirculantMatrixDigraph(8, jumps)
	if !sliceutils.EqualByteMatrix(a, got.Matrix()) {
		t.Errorf("Expected %v, but got %v", a, got.Matrix())
	}
}

*/
