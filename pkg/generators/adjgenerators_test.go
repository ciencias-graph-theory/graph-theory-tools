package generators

import (
	"math/rand"
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

func TestIsCompleteMatrixGraph(t *testing.T) {
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
	m, _ := graph.NewGraphFromMatrix(c)
	n, _ := graph.NewGraphFromMatrix(d)
	if !IsCompleteMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteMatrixGraph(k),
		)
	}
	if IsCompleteMatrixGraph(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteMatrixGraph(k),
		)
	}
	if IsCompleteMatrixGraph(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteMatrixGraph(k),
		)
	}
	if !IsCompleteMatrixGraph(n) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteMatrixGraph(k),
		)
	}
}

// TestCompleteMatrixGraph calls CompleteMatrixGraph with five different
// randomly generated graphs, and checks each of them to be a complete graph by
// exploring their adjacency matrices.
func TestCompleteMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		k := CompleteMatrixGraph(n)
		a := k.Adjacency()
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

func TestIsCompleteBipartiteMatrixGraph(t *testing.T) {
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
	m, _ := graph.NewGraphFromMatrix(c)
	n, _ := graph.NewGraphFromMatrix(d)
	o, _ := graph.NewGraphFromMatrix(e)
	if !IsCompleteBipartiteMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteBipartiteMatrixGraph(k),
		)
	}
	if IsCompleteBipartiteMatrixGraph(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartiteMatrixGraph(k),
		)
	}
	if IsCompleteBipartiteMatrixGraph(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartiteMatrixGraph(k),
		)
	}
	if IsCompleteBipartiteMatrixGraph(n) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCompleteBipartiteMatrixGraph(k),
		)
	}
	if !IsCompleteBipartiteMatrixGraph(o) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCompleteBipartiteMatrixGraph(k),
		)
	}
}

// TestCompleteBipartiteMatrixGraph calls CompleteBipartiteMatrixGraph with five
// different randomly generated graphs, and checks each of them to be a complete
// bipartite graph by exploring their adjacency matrices.
func TestCompleteBipartiteMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		m := rand.Intn(1000)
		k := CompleteBipartiteMatrixGraph(n, m)
		a := k.Adjacency()
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

// TestIsCycleMatrixGraph calls IsCycleMatrixGraph with different hardcoded
// graphs, including cycles of different lengths, disconnected 2-regular graphs,
// and other non-cycle graphs.
func TestIsCycleMatrixGraph(t *testing.T) {
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
	if !IsCycleMatrixGraph(g) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(g),
		)
	}
	if !IsCycleMatrixGraph(h) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(h),
		)
	}
	if !IsCycleMatrixGraph(i) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsCycleMatrixGraph(i),
		)
	}
	if IsCycleMatrixGraph(j) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycleMatrixGraph(j),
		)
	}
	if IsCycleMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsCycleMatrixGraph(k),
		)
	}
}

// TestCycleMatrixGraph calls CycleMatrixGraph with five different
// randomly generated numbers, and checks each of them to be a cycle by
// exploring their adjacency matrices.
func TestCycleMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			c := CycleMatrixGraph(n)
			d := c.DegreeSequence()
			for _, v := range d {
				if v != 2 {
					t.Error("The graph is not 2-regular")
				}
			}
			a := c.Adjacency()
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

// TestIsPathMatrixGraph calls IsPathMatrixGraph with different hardcoded
// graphs, including cycles of different lengths, linear forests,
// and other non-cycle graphs.
func TestIsPathMatrixGraph(t *testing.T) {
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
	if !IsPathMatrixGraph(h) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPathMatrixGraph(h),
		)
	}
	if !IsPathMatrixGraph(i) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPathMatrixGraph(i),
		)
	}
	if !IsPathMatrixGraph(j) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPathMatrixGraph(j),
		)
	}
	if IsPathMatrixGraph(k) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPathMatrixGraph(k),
		)
	}
	if IsPathMatrixGraph(l) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPathMatrixGraph(l),
		)
	}
	if IsPathMatrixGraph(m) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsPathMatrixGraph(m),
		)
	}
	if !IsPathMatrixGraph(n) {
		t.Errorf(
			"Expected %v, but got %v",
			true,
			IsPathMatrixGraph(n),
		)
	}
}

// TestPathMatrixGraph calls PathMatrixGraph with five different
// randomly generated numbers, and checks each of them to be a path by
// exploring their adjacency matrices.
func TestPathMatrixGraph(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Intn(1000)
		if n > 2 {
			c := PathMatrixGraph(n)
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
			a := c.Adjacency()
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
		if !sliceutils.EqualByteMatrix(a, got.Adjacency()) {
			t.Errorf("Expected %v, but got %v", a, got.Adjacency())
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
	if !sliceutils.EqualByteMatrix(a, got.Adjacency()) {
		t.Errorf("Expected %v, but got %v", a, got.Adjacency())
	}
}

*/
