package generators

import (
	"github.com/Japodrilo/graph-theory-tools/pkg/graph"
	"math/rand"
	"testing"
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
	k := graph.NewMatrixGraph(a)
	l := graph.NewMatrixGraph(b)
	m := graph.NewMatrixGraph(c)
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
