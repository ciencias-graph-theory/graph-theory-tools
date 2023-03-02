package drawer

import (
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

func TestCircular(t *testing.T) {

	l := [][]int{
		{1, 2, 17, 9},
		{0, 3},
		{0, 3},
		{1, 2},
		{5},
		{4},
		{7},
		{6},
		{9},
		{8, 0},
		{11},
		{10},
		{13},
		{12},
		{15},
		{14},
		{17},
		{16, 0},
		{19},
		{18},
		{21},
		{20},
		{23},
		{22},
	}
	g, _ := graph.NewGraphFromList(l)
	Circular(g, 200, 200)

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

	Circular(petersen, 200, 200)
	Circular(g, 200, 200)
}
