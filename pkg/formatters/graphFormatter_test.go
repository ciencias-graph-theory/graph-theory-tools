package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
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
	Q3g6 := "Gr_iOk"

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

func TestParseByteSliceFormat6(t *testing.T) {
	// Byte slices examples.
	a := []byte{
		1, 1, 0, 0, 1, 1, 1,
		0, 0, 0, 0, 0, 1, 0,
		1, 0, 1, 0, 0, 1, 0,
		0, 0, 0, 1, 0, 1, 1}

	b := []byte{1, 1, 1, 1, 1, 1}

	c := []byte{1, 1, 0, 0, 1, 1}

	d := []byte{
		0, 0, 1, 0, 1,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 1, 0, 0, 1,
		0, 0, 0, 0, 0}

	// Expected int slices.
	ia := []int{114, 95, 105, 79, 107}
	ib := []int{126}
	ic := []int{114}
	id := []int{73, 63, 65, 79, 63}

	// Obtained int slices.
	IA := parseByteSliceFormat6(a, false)
	IB := parseByteSliceFormat6(b, false)
	IC := parseByteSliceFormat6(c, false)
	ID := parseByteSliceFormat6(d, false)

	// Check that the obtained slices are correct.
	if !sliceutils.EqualIntSlice(ia, IA) {
		t.Errorf("Parsing Error: Expected %v but got %v", ia, IA)
	}

	if !sliceutils.EqualIntSlice(ib, IB) {
		t.Errorf("Parsing Error: Expected %v but got %v", ib, IB)
	}

	if !sliceutils.EqualIntSlice(ic, IC) {
		t.Errorf("Parsing Error: Expected %v but got %v", ic, IC)
	}

	if !sliceutils.EqualIntSlice(id, ID) {
		t.Errorf("Parsing Error: Expected %v but got %v", id, ID)
	}
}

func TestParseOrderFormat6(t *testing.T) {
	// Test values.
	input := []int{8, 30, 136, 12345, 460175067}

	// Expected output.
	output := [][]int{
		{71},
		{93},
		{126, 63, 65, 71},
		{126, 66, 63, 120},
		{126, 126, 63, 90, 90, 90, 90, 90},
	}

	// Check that obtained values are correct.
	for i := 0; i < len(input); i++ {
		obOutput := parseOrderFormat6(input[i])
		if !sliceutils.EqualIntSlice(obOutput, output[i]) {
			t.Errorf("Parsing Error: Expected %v but got %v", output[i], obOutput)
		}
	}
}

func TestToLoop6(t *testing.T) {
	// The adj. matrix of complete graph with four vertices with loops.
	a := [][]byte{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	// The adj. matrix of a 4-cycle with loops.
	b := [][]byte{
		{1, 1, 1, 0},
		{1, 1, 0, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 1},
	}

	// The adj. matrix of a 3-cube with loops.
	c := [][]byte{
		{1, 1, 1, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 1, 1, 1},
	}

	K4l, _ := graph.NewGraphFromMatrix(a)
	C4l, _ := graph.NewGraphFromMatrix(b)
	Q3l, _ := graph.NewGraphFromMatrix(c)

	// Expected Loop6 strings from the previous graphs.
	K4l6 := ";C~{"
	C4l6 := ";C|["
	Q3l6 := ";G|]HYSV"

	// Check if obtained formats are correct.
	K4L6 := ToLoop6(K4l)
	C4L6 := ToLoop6(C4l)
	Q3L6 := ToLoop6(Q3l)

	if K4L6 != K4l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", K4l6, K4L6)
	}

	if C4L6 != C4l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", C4l6, C4L6)
	}

	if Q3L6 != Q3l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Q3l6, Q3L6)
	}
}

func TestToDigraph6(t *testing.T) {
	// Example of a digraph with n = 5 and edges:
	// 0 -> 2, 0 -> 4,
	// 3 -> 1, 3 -> 4.
	a := [][]byte{
		{0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
	}

	// Example of digraph with n = 8 and edges:
	// 0 -> 1, 0 -> 2,
	// 3 -> 1, 3 -> 2,
	// 4 -> 0,
	// 5 -> 2, 5 -> 4, 5 -> 7,
	// 6 -> 1, 6 -> 4, 6 -> 7,
	// 7 -> 3.
	b := [][]byte{
		{0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 0, 0},
	}

	Da := graph.NewDigraphFromMatrix(a)
	Db := graph.NewDigraphFromMatrix(b)

	// Expected Digraph6 strings from the previous graphs.
	Dad6 := "&DI?AO?"
	Dbd6 := "&GW???WG?hQP?"

	// Check if obtained formats are correct.
	DAD6 := ToDigraph6(Da)
	DBD6 := ToDigraph6(Db)

	if DAD6 != Dad6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Dad6, DAD6)
	}

	if DBD6 != Dbd6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Dbd6, DBD6)
	}
}
