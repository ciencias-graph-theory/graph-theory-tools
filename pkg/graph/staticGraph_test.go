package graph

import (
	"reflect"
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

func TestImplementsGraphInterface(t *testing.T) {
	var _ Graph = &StaticGraph{}
}

// TestMatrix calls NewMatrixGraph with an adjacency matrix, and then compares
// this adjacency matrix with the one returned by Matrix()
func TestNewGraphFromMatrix(t *testing.T) {
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
	petersen, err := NewGraphFromMatrix(matrix)
	got, _ := petersen.Matrix()
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
	_, err = NewGraphFromMatrix(a)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}
	_, err = NewGraphFromMatrix(b)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}
}

func TestNewGraphFromList(t *testing.T) {
	a := [][]int{
		{0, 1, 2, 3},
		{0, 2, 3},
		{0, 3},
		{0, 1, 2, 3},
	}

	b := [][]int{
		{1, 2, 3},
		{0, 2, 3},
		{0, 3},
		{0, 1, 2},
	}
	_, err := NewGraphFromList(a)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}
	_, err = NewGraphFromList(b)
	if err == nil {
		t.Errorf("Expected an error, got %v", nil)
	}

	c := [][]int{
		{1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{0, 1, 2},
	}
	_, err = NewGraphFromList(c)
	if err != nil {
		t.Errorf("Didn't expect an error, got %v", nil)
	}
}

func TestMatrix(t *testing.T) {
	a := [][]byte{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
	}
	g1 := NewFromMatrix(a)
	m1, _ := g1.Matrix()
	if !reflect.DeepEqual(a, m1) {
		t.Errorf("Returned matrix differs from original one")
	}

	b := [][]byte{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}
	g2 := NewFromMatrix(b)
	m2, _ := g2.Matrix()
	if !reflect.DeepEqual(b, m2) {
		t.Errorf("Returned matrix differs from original one")
	}
}

func TestList(t *testing.T) {
	a := [][]int{
		{0, 1, 2, 3},
		{0, 2, 3},
		{0, 3},
		{0, 1, 2, 3},
	}
	g1 := NewFromList(a)
	l1, _ := g1.List()
	if !reflect.DeepEqual(a, l1) {
		t.Errorf("Returned list differs from original one")
	}

	b := [][]int{
		{1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{0, 1, 2},
	}
	g2 := NewFromList(b)
	l2, _ := g2.List()
	if !reflect.DeepEqual(b, l2) {
		t.Errorf("Returned list differs from original one")
	}
}

func TestMatrixToList(t *testing.T) {
	m1 := [][]byte{
		{1, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}
	m2 := [][]byte{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}
	r, err := matrixToList(m1)
	l := [][]int{
		{0, 1, 2},
		{0, 2, 3},
		{0, 1, 3},
		{1, 2, 3},
	}
	if err != nil {
		t.Errorf("Didn't expect an error, got %v", nil)
	}
	if !reflect.DeepEqual(*r, l) {
		t.Errorf("List does not correspond to given matrix")
	}
	_, err = matrixToList(m2)
	if err != assymetricMatrixError {
		t.Errorf("Expected an error, got %v", err)
	}
}

func TestListToMatrix(t *testing.T) {
	l1 := [][]int{
		{0, 1, 2},
		{0, 2, 3},
		{0, 1, 3},
		{1, 2, 3},
	}
	l2 := [][]int{
		{0, 1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{1, 2, 3},
	}
	m := [][]byte{
		{1, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}
	r, err := listToMatrix(l1)
	if err != nil {
		t.Errorf("Didn't expect an error, got %v", nil)
	}
	if !reflect.DeepEqual(*r, m) {
		t.Errorf("Matrix does not correspond to given list")
	}
	_, err = listToMatrix(l2)
	if err != invalidListError {
		t.Errorf("Expected an error, got %v", err)
	}
}

// TestNewMatrixGraphSequence calls NewMatrixGraph with a fixed adjacency
// matrix, then compares the sequence obtained from DegreeSequence and the
// (previously known) degree sequence corresponding to the adjacency matrix.
func TestDegreeSequence(t *testing.T) {
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
	petersen := NewFromMatrix(adjacency)
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
	petersen := NewFromMatrix(adjacency)
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

// TestNeighbours calls NeighboursSet on a graph and checks that the neighbours
// set is correctly computed for each vertex.
func TestNeighboursSet(t *testing.T) {
	m := [][]byte{
		{1, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}
	g1 := NewFromMatrix(m)
	for i := 0; i < len(m); i++ {
		s := g1.NeighboursSet(i)
		for j := 0; j < len(m); j++ {
			if m[i][j] == 1 && !s.Contains(j) {
				t.Errorf("Neighbour %v expected but not present", j)
			}
		}
	}
	l := [][]int{
		{0, 1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{1, 2, 3},
	}
	g2 := NewFromList(l)
	for i := 0; i < len(l); i++ {
		s := g2.NeighboursSet(i)
		want := l[i]
		for _, n := range want {
			if !s.Contains(n) {
				t.Errorf("Neighbour %v expected but not present", n)
			}
		}
	}
}
