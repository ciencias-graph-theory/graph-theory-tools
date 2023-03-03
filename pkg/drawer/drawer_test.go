package drawer

import (
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/fileutils/writer"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

// TestCircular uses the Circular graph drawing algorithm and outputs two svg
// files.
func TestCircular(t *testing.T) {
	// Ad hoc adjacency list to represent a test graph with possible vertex/edge
	// intersections.
	l := [][]int{
		{1, 2, 17, 9},
		{0, 3},
		{0, 3},
		{1, 2},
		{5},
		{4, 9, 17},
		{7},
		{6},
		{9},
		{8, 5, 0},
		{11, 13, 15},
		{10},
		{13},
		{12, 10},
		{15},
		{14, 10},
		{17},
		{16, 0, 5},
		{19},
		{18},
		{21},
		{20},
		{23},
		{22},
	}
	g, _ := graph.NewGraphFromList(l)
	// Petersen graph's adjacency matrix.
	adjacency := [][]byte{
		{0, 1, 0, 0, 1, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 1, 1, 0, 0},
	}
	petersen := graph.NewFromMatrix(adjacency)

	matrixDraw := Circular(petersen, 200, 200)
	listDraw := Circular(g, 200, 200)

	writer.Write("testMatrix.svg", []byte(matrixDraw.Content()))
	writer.Write("testList.svg", []byte(listDraw.Content()))
}
