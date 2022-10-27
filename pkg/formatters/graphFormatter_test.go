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

// TestFromGraph6 calls FromGraph6 with a string corresponding a the graph6
// format of a graph G, it checks that the obtained graph is the one
// corresponding to the format.
func TestFromGraph6(t *testing.T) {
	// Example strings.
	K4g6 := "C~"
	C4g6 := "Cr"
	Q3g6 := "Gr_iOk"

	// Expected adj. matrices of the examples.
	K4m := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// The adj. matrix of a 4-cycle.
	C4m := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// The adj. matrix of a 3-cube.
	Q3m := [][]byte{
		{0, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 0},
	}

	// Obtain the graphs based on the strings
	K4, _ := FromGraph6(K4g6).Matrix()
	C4, _ := FromGraph6(C4g6).Matrix()
	Q3, _ := FromGraph6(Q3g6).Matrix()

	// Compare that the obtained graphs are correct.
	if !sliceutils.EqualByteMatrix(K4m, K4) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", K4m, K4)
	}

	if !sliceutils.EqualByteMatrix(C4m, C4) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", C4m, C4)
	}

	if !sliceutils.EqualByteMatrix(Q3m, Q3) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", Q3m, Q3)
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

func TestDetermineOrderAndEdges(t *testing.T) {
	// Test slices.
	input := [][]int{
		{71},
		{93},
		{126, 63, 65, 71},
		{126, 66, 63, 120},
		{126, 126, 63, 90, 90, 90, 90, 90},
	}

	// Expected values.
	output := []int{8, 30, 136, 12345, 460175067}

	// Check that obtained values are correct.
	for i := 0; i < len(input); i++ {
		obOutput, _ := determineOrderAndEdges(input[i])
		if output[i] != obOutput {
			t.Errorf("Parsing Error: Expected %d but got %d", output[i], obOutput)
		}
	}
}

func TestInverseParseOrderFormat6(t *testing.T) {
	// Test slices.
	input := [][]int{
		{71},
		{93},
		{63, 65, 71},
		{66, 63, 120},
		{63, 90, 90, 90, 90, 90},
	}

	// Expected values.
	output := []int{8, 30, 136, 12345, 460175067}

	// Check that obtained values are correct.
	for i := 0; i < len(input); i++ {
		obOutput := inverseParseOrderFormat6(input[i])
		if output[i] != obOutput {
			t.Errorf("Parsing Error: Expected %d but got %d", output[i], obOutput)
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

	// The adj. matrix of an arbitrary graph with no loops.
	d := [][]byte{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 1},
		{0, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 1},
		{0, 0, 1, 1, 1, 0},
	}

	// The adj. matrix of an arbitrary graph with some loops.
	e := [][]byte{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
	}

	K4l, _ := graph.NewGraphFromMatrix(a)
	C4l, _ := graph.NewGraphFromMatrix(b)
	Q3l, _ := graph.NewGraphFromMatrix(c)
	Gdl, _ := graph.NewGraphFromMatrix(d)
	Gel, _ := graph.NewGraphFromMatrix(e)

	// Expected Loop6 strings from the previous graphs.
	K4l6 := ";C~{"
	C4l6 := ";C|["
	Q3l6 := ";G|]HYSV"
	Gdl6 := ";EUGPo"
	Gel6 := ";DUGW"

	// Check if obtained formats are correct.
	K4L6 := ToLoop6(K4l)
	C4L6 := ToLoop6(C4l)
	Q3L6 := ToLoop6(Q3l)
	GdL6 := ToLoop6(Gdl)
	GeL6 := ToLoop6(Gel)

	if K4L6 != K4l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", K4l6, K4L6)
	}

	if C4L6 != C4l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", C4l6, C4L6)
	}

	if Q3L6 != Q3l6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Q3l6, Q3L6)
	}

	if GdL6 != Gdl6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Gdl6, GdL6)
	}

	if GeL6 != Gel6 {
		t.Errorf("Loop6 Error: Expected %s but got %v", Gel6, GeL6)
	}
}

// TestFromLoop6 calls FromLoop6 with a string corresponding the loop6
// format of a graph G, it checks that the obtained graph is the one
// corresponding to the format.
func TestFromLoop6(t *testing.T) {
	// Example strings.
	K4l6 := ";C~{"
	C4l6 := ";C|["
	Q3l6 := ";G|]HYSV"

	// Expected adj. matrices of the examples.
	// The adj. matrix of complete graph with four vertices with loops.
	K4lm := [][]byte{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	// The adj. matrix of a 4-cycle with loops.
	C4lm := [][]byte{
		{1, 1, 1, 0},
		{1, 1, 0, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 1},
	}

	// The adj. matrix of a 3-cube with loops.
	Q3lm := [][]byte{
		{1, 1, 1, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 1, 1, 1},
	}

	// Obtain the graphs based on the strings
	K4l, _ := FromLoop6(K4l6).Matrix()
	C4l, _ := FromLoop6(C4l6).Matrix()
	Q3l, _ := FromLoop6(Q3l6).Matrix()

	// Compare that the obtained graphs are correct.
	if !sliceutils.EqualByteMatrix(K4lm, K4l) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", K4lm, K4l)
	}

	if !sliceutils.EqualByteMatrix(C4lm, C4l) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", C4lm, C4l)
	}

	if !sliceutils.EqualByteMatrix(Q3lm, Q3l) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", Q3lm, Q3l)
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

func TestFromDigraph6(t *testing.T) {
	// Examples of Digraph6 strings.
	Dad6 := "&DI?AO?"
	Dbd6 := "&GW???WG?hQP?"

	// Expected adj. matrices of the previous strings.
	a := [][]byte{
		{0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
	}

	// Example of digraph with n = 8 and edges:
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

	Da, _ := FromDigraph6(Dad6).Matrix()
	Db, _ := FromDigraph6(Dbd6).Matrix()

	// Compare that the obtained graphs are correct.
	if !sliceutils.EqualByteMatrix(a, Da) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", a, Da)
	}

	if !sliceutils.EqualByteMatrix(b, Db) {
		t.Errorf("Graph6 Conversion Error: Expected %v but got %v", b, Db)
	}
}

// TestFromSparse6 calls FromSparse6 with a string corresponding the sparse6
// format of a graph G, it checks that the obtained graph is the one
// corresponding to the format.
func TestFromSparse6(t *testing.T) {
	// Example string obtained from the format6 specifications website:
	// https://users.cecs.anu.edu.au/~bdm/data/formats.txt
	ex1 := ":Fa@x^"

	// Another example with the following adjacencies:
	// {{0, 0}, {1, 2}, {1, 3}, {2, 3}, {4, 5}, {4, 5}}
	ex2 := ":E?`cdPK"

	// Expected adjacency matrix from the previous strings.
	ex1A := [][]byte{
		{0, 1, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 0},
	}

	ex2A := [][]byte{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 2, 0},
	}

	// Build the graphs from the sparse6 string.
	ex1G := FromSparse6(ex1)
	ex2G := FromSparse6(ex2)

	// Obtained matrices.
	ex1Gm, _ := ex1G.Matrix()
	ex2Gm, _ := ex2G.Matrix()

	// Compare the obtained graphs with the expected ones.
	if !sliceutils.EqualByteMatrix(ex1A, ex1Gm) {
		t.Errorf("Sparse6 Conversion Error: Expected %v but got %v", ex1A, ex1Gm)
	}

	if !sliceutils.EqualByteMatrix(ex2A, ex2Gm) {
		t.Errorf("Sparse6 Conversion Error: Expected %v but got %v", ex2A, ex2Gm)
	}
}

// TestToSparse6 calls ToSparse6 with a graph, then checks that the
// obtained sparse6 string is the correct one.
func TestToSparse6(t *testing.T) {
	// Example adj. matrices.
	ex1A := [][]byte{
		{0, 1, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 0},
	}

	ex2A := [][]byte{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 2, 0},
	}

	// The sparse6 strings expected from the previous strings.
	ex1 := ":Fa@x^"
	ex2 := ":EA`clPN"

	// Build the graphs from the adj. matrix.
	G1, _ := graph.NewGraphFromMatrix(ex1A)
	G2, _ := graph.NewGraphFromMatrix(ex2A)

	// Get the sparse6 strings from the previous strings.
	G1S6 := ToSparse6(G1)
	G2S6 := ToSparse6(G2)

	// Compare the obtained sparse6 strings with the expected ones.
	if ex1 != G1S6 {
		t.Errorf("Sparse6 Conversion Error: Expected %s but got %s", ex1, G1S6)
	}

	if ex2 != G2S6 {
		t.Errorf("Sparse6 Conversion Error: Expected %s but got %s", ex2, G2S6)
	}
}
