package graph

import (
  "testing"
  "reflect"

  "github.com/Japodrilo/graph-theory-tools/pkg/sliceutils"
)


// TestAdjacency calls NewMarixGraph with an
// adjacency matrix, and then compares this
// adjacency matrix with the one returned by
// Adjacency()
func TestAdjacency(t *testing.T) {
  adjacency := [][]byte{
    {0,1,0,0,1,1,0,0,0,0},
    {1,0,1,0,0,0,1,0,0,0},
    {0,1,0,1,0,0,0,1,0,0},
    {0,0,1,0,1,0,0,0,1,0},
    {1,0,0,1,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,1,1,0},
    {0,1,0,0,0,0,0,0,1,1},
    {0,0,1,0,0,1,0,0,0,1},
    {0,0,0,1,0,1,1,0,0,0},
    {0,0,0,0,1,0,1,1,0,0},
  }
  petersen := NewMatrixGraph(adjacency)
  got := petersen.Adjacency()
  if !reflect.DeepEqual(adjacency, got) {
    t.Errorf("Expected %v, got %v", adjacency, got)
  }
}

// TestNewMatrixGraphSequence calls NewMatrixGraph
// with an adjacency matrix,
func TestNewMatrixGraphSequence(t *testing.T) {
  adjacency := [][]byte{
    {0,1,0,0,1,1,0,0,0,0},
    {1,0,1,0,0,0,1,0,0,0},
    {0,1,0,1,0,0,0,1,0,0},
    {0,0,1,0,1,0,0,0,1,0},
    {1,0,0,1,0,0,0,0,0,1},
    {1,0,0,0,0,0,0,1,1,0},
    {0,1,0,0,0,0,0,0,1,1},
    {0,0,1,0,0,1,0,0,0,1},
    {0,0,0,1,0,1,1,0,0,0},
    {0,0,0,0,1,0,1,1,0,0},
  }
  petersen := NewMatrixGraph(adjacency)
  if petersen.degreeSequence != nil {
    t.Errorf("The degree sequence was expected to be nil")
  }
  want := []int{3,3,3,3,3,3,3,3,3,3}
  got := petersen.DegreeSequence()
  if !sliceutils.EqualIntSlice(want, got) {
    t.Errorf("Expected %v, got %v", want, got)
  }
  got = petersen.degreeSequence
  if !sliceutils.EqualIntSlice(want, got) {
    t.Errorf("Expected %v, got %v", want, got)
  }
}
