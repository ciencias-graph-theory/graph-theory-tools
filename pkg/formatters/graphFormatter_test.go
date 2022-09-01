package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
	"testing"
)

// TestToGraph6 calls ToGraph6 with a Graph G and check if the obtained format
// is correct.
func TestToGraph6(t *testing.T) {

	// The adj. matrix of complete graph with four vertices.
	a := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// The adj. matrix of a 4-cycle.
	b := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// The adj. matrix of a 3-cube.
	c := [][]byte{
		{0, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 0},
	}

	// Build the graphs.
	K4, _ := graph.NewGraphFromMatrix(a)
	C4, _ := graph.NewGraphFromMatrix(b)
	Q3, _ := graph.NewGraphFromMatrix(c)

	// Expected formats.
	K4g6 := "C~"
	C4g6 := "Cr"
	Q3g6 := "GKwIcJ"

	// Check if obtained formats are correct.
	K4G6 := ToGraph6(K4)
	C4G6 := ToGraph6(C4)
	Q3G6 := ToGraph6(Q3)

	if K4G6 != K4g6 {
		t.Errorf("Graph6 Error: Expected %s but got %v", K4g6, K4G6)
	}

	if C4G6 != C4g6 {
		t.Errorf("Graph6 Error: Expected %s but got %v", C4g6, C4G6)
	}

	if Q3G6 != Q3g6 {
		t.Errorf("Graph6 Error: Expected %s but got %v", Q3g6, Q3G6)
	}
}
