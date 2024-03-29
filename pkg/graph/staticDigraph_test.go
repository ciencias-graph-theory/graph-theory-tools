package graph

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// TestAdjacency calls NewMatrixDigraph with an
// adjacency matrix, and then compares this
// adjacency matrix with the one returned by
// Adjacency(). This test is identical to the one for regular
// graphs.
func TestNewDigraphFromMatrix(t *testing.T) {
	matrix := [][]byte{
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
	petersen := NewDigraphFromMatrix(matrix)
	got, _ := petersen.Matrix()
	if !reflect.DeepEqual(matrix, got) {
		t.Errorf("Expected %v, got %v", matrix, got)
	}
}

// TestDigraphOrder tests that order is computed correctly for
// digraphs.
func TestDigraphOrder(t *testing.T) {
	order := rand.Intn(100)
	matrix := make([][]byte, order)
	for i := 0; i < order; i++ {
		matrix[i] = make([]byte, order)
	}
	digraph := NewDigraphFromMatrix(matrix)
	if order != digraph.Order() {
		t.Errorf("Expected %v, got %v",
			order, digraph.Order())
	}
}

// TestDigraphDegreeSequence tests that degree, in-degree, and
// out-degree sequences are computed correctly for digraphs.
// First example digraph is C3*.
// Second is from page 11 of article "3-transitive digraphs" by
// Cesar Hernandez-Cruz, with loops added on vertices 4 and 5.
func TestDigraphDegreeSequence(t *testing.T) {
	matrix1 := [][]byte{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
	}
	matrix2 := [][]byte{
		{0, 1, 0, 0, 1, 0},
		{0, 0, 1, 1, 0, 1},
		{0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0, 1},
	}
	digraph1 := NewDigraphFromMatrix(matrix1)
	digraph2 := NewDigraphFromMatrix(matrix2)
	want1 := []int{2, 3, 3}
	want2 := []int{2, 5, 3, 3, 7, 4}

	// In-degree and out-degree sequences have not yet been computed.
	if digraph1.indegreeSequence != nil {
		t.Errorf("In digraph 1, in-degree" +
			"sequence was expected to be nil")
	}
	if digraph1.outdegreeSequence != nil {
		t.Errorf("In digraph 1, out-degree" +
			"sequence was expected to be nil")
	}
	if digraph2.indegreeSequence != nil {
		t.Errorf("In digraph 2, in-degree" +
			"sequence was expected to be nil")
	}
	if digraph2.outdegreeSequence != nil {
		t.Errorf("In digraph 2, out-degree" +
			"sequence was expected to be nil")
	}

	got1 := digraph1.DegreeSequence()
	got2 := digraph2.DegreeSequence()

	// Testing DegreeSequence
	if !sliceutils.EqualIntSlice(want1, got1) {
		t.Errorf("In graph 1, expected %v, got %v", want1,
			got1)
	}
	if !sliceutils.EqualIntSlice(want2, got2) {
		t.Errorf("In graph 2, expected %v, got %v", want2,
			got2)
	}

	want1 = []int{1, 2, 1}
	want2 = []int{0, 2, 2, 1, 4, 3}

	got1 = digraph1.IndegreeSequence()
	got2 = digraph2.IndegreeSequence()

	// Testing IndegreeSequence
	if !sliceutils.EqualIntSlice(want1, got1) {
		t.Errorf("In graph 1, expected %v, got %v", want1,
			got1)
	}
	if !sliceutils.EqualIntSlice(want2, got2) {
		t.Errorf("In graph 2, expected %v, got %v", want2,
			got2)
	}

	want1 = []int{1, 1, 2}
	want2 = []int{2, 3, 1, 2, 3, 1}

	got1 = digraph1.OutdegreeSequence()
	got2 = digraph2.OutdegreeSequence()

	// Testing IndegreeSequence
	if !sliceutils.EqualIntSlice(want1, got1) {
		t.Errorf("In graph 1, expected %v, got %v", want1,
			got1)
	}
	if !sliceutils.EqualIntSlice(want2, got2) {
		t.Errorf("In graph 2, expected %v, got %v", want2,
			got2)
	}
}

// TestSize tests that size of digraph is computed correctly. This test is
// very similar to the one for MatrixGraphs.
func TestDigraphSize(t *testing.T) {
	matrix := [][]byte{
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
	petersen := NewDigraphFromMatrix(matrix)
	if petersen.indegreeSequence != nil {
		t.Errorf("The indegree sequence was expected to be nil")
	}
	if petersen.outdegreeSequence != nil {
		t.Errorf("The out-degree sequence was expected to be nil")
	}
	want := 10 * 3
	got := petersen.Size()
	if want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}
	// Testing twice because size is computed differently depending on
	// whether we have degree sequence or not.
	got = petersen.Size()
	if want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}
	if petersen.indegreeSequence == nil {
		t.Errorf("The indegree sequence was not expected to be nil")
	}
	if petersen.outdegreeSequence == nil {
		t.Errorf("The out-degree sequence was not expected to be nil")
	}
	wantD := []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
	gotD := petersen.DegreeSequence()
	if !sliceutils.EqualIntSlice(wantD, gotD) {
		t.Errorf("Expected %v, got %v", wantD, gotD)
	}
}
