package graph

import (
	"reflect"
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// TestMatrix calls NewMatrixGraph with an adjacency matrix, and then compares
// this adjacency matrix with the one returned by Matrix()
func TestDigraphMatrix(t *testing.T) {
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
	petersen, err := NewMatrixGraph(matrix)
	got := petersen.Matrix()
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}
	if !reflect.DeepEqual(matrix, got) {
		t.Errorf("Expected %v, got %v", matrix, got)
	}
	a := [][]byte{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}
	b := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 0},
	}
	_, err = NewMatrixGraph(a)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}
	_, err = NewMatrixGraph(b)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}
}

// TestNewMatrixGraphSequence calls NewMatrixGraph with a fixed adjacency
// matrix, then compares the sequence obtained from DegreeSequence and the
// (previously known) degree sequence corresponding to the adjacency matrix.
func TestNewMatrixGraphSequence(t *testing.T) {
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
	petersen := NewMatrixGraphU(adjacency)
	if petersen.degreeSequence != nil {
		t.Errorf("The degree sequence was expected to be nil")
	}
	want := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	got := petersen.DegreeSequence()
	if !sliceutils.EqualIntSlice(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
	got = petersen.degreeSequence
	if !sliceutils.EqualIntSlice(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

// TestAddEdge calls AddEdge with different invalid values and check that an
// error is successfully thrown.   Then, it calls AddEdge with a valid value,
// and checks that the given edge was added successfully.
func TestAddEdge(t *testing.T) {
	adjacency := [][]byte{
		{0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
	}
	g := NewMatrixGraphU(adjacency)
	err := g.AddEdge(-1, 2)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	err = g.AddEdge(1, 5)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	err = g.AddEdge(-1, -2)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	err = g.AddEdge(-1, 6)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	err = g.AddEdge(19, 25)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	err = g.AddEdge(0, 2)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if g.matrix[0][2] != 1 || g.matrix[2][0] != 1 {
		t.Errorf("The edge was not successfully added")
	}
	if g.matrix[0][1] != 1 || g.matrix[1][0] != 1 {
		t.Errorf("The edge was not successfully added")
	}
}

// TestSize calls Size on a matrix and checks that is correctly computed
func TestSize(t *testing.T) {
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
	petersen := NewMatrixGraphU(adjacency)
	if petersen.degreeSequence != nil {
		t.Errorf("The degree sequence was expected to be nil")
	}
	want := (10 * 3) / 2
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
	if petersen.degreeSequence == nil {
		t.Errorf("The degree sequence was not expected to be nil")
	}
	wantD := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	gotD := petersen.DegreeSequence()
	if !sliceutils.EqualIntSlice(wantD, gotD) {
		t.Errorf("Expected %v, got %v", wantD, gotD)
	}
}
