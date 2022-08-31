package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	// "github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
	"testing"
)

// TestObtainUppertriangle calls obtainUppertriangle with an adjacency matrix,
// and then compares the obtained vector with the upper triangle of the given
// matrix.
func TestObtainUpperTriangle(t *testing.T) {

	// Complete graph with four vertices.
	a := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// A 4-cycle.
	b := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// A cube.
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

	// The corresponding vectors of the upper triangle of each adjacency matrix.
	av := []byte{1, 1, 1, 1, 1, 1}
	bv := []byte{1, 1, 0, 0, 1, 1}
	cv := []byte{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1}

	Av := obtainUpperTriangle(a)
	Bv := obtainUpperTriangle(b)
	Cv := obtainUpperTriangle(c)

	if !sliceutils.EqualByteSlice(Av, av) {
		t.Errorf("Formatting error: Expected %v but got %v", av, Av)
	}

	if !sliceutils.EqualByteSlice(Bv, bv) {
		t.Errorf("Formatting error: Expected %v but got %v", bv, Bv)
	}

	if !sliceutils.EqualByteSlice(Cv, cv) {
		t.Errorf("Formatting error: Expected %v but got %v", cv, Cv)
	}
}
