package drawer

import (
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

func TestCircular(t *testing.T) {

	l := [][]int{
		{1, 2},
		{0, 3},
		{0, 3},
		{1, 2},
	}
	g, _ := graph.NewGraphFromList(l)
	Circular(g, 200, 200)
}
