package graph

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/Japodrilo/graph-theory-tools/pkg/sliceutils"
)

// TestAdjacency calls NewMatrixDigraph with an
// adjacency matrix, and then compares this
// adjacency matrix with the one returned by
// Adjacency(). This test is identical to the one for regular
// graphs.
func TestAdjacency(t *testing.T) {
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
	petersen := NewMatrixDigraph(adjacency)
	got := petersen.Adjacency()
	if !reflect.DeepEqual(adjacency, got) {
		t.Errorf("Expected %v, got %v", adjacency, got)
	}
}

// TestDigraphOrder tests that order is computed correctly for
// digraphs.
func TestDigraphOrder(t *testing.T) {
	order := rand.Intn(100)
	adjacency := make([][]byte, order)
	for i := 0; i < order; i++ {
		adjacency[i] = make([]byte, order)
	}
	digraph := NewMatrixDigraph(adjacency)
	if order != digraph.Order() {
		t.Errorf("Expected %v, got %v",
			order, digraph.Order())
	}
}

// TestDigraphDegreeSequence tests that degree sequence is
// computed correctly for digraphs.
// Example digraph is a triangle:
// 0 -> 1 <-> 2 -> 0
func TestDigraphDegreeSequence(t *testing.T) {
	adjacency := [][]byte{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
	}
	digraph := NewMatrixDigraph(adjacency)
	want := []int{2, 3, 3}
	got := digraph.DegreeSequence()
	if !sliceutils.EqualIntSlice(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
