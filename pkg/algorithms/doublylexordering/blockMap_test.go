package doublylexordering

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
	"testing"
)

// TestBlockMapAdd tests the Add function for BlockMaps.
func TestBlockMapAdd(t *testing.T) {
	// Define an emtpy block map.
	BM := NewBlockMap()

	// Integer sets.
	integerSetsPairs := [][][]int{
		{{0, 1, 2}, {0, 1, 2}},
		{{2, 4}, {0, 1, 2}},
		{{3}, {3}},
		{{0, 1}, {0, 1}},
		{{1}, {0, 1}},
		{{1}, {1}},
	}

	// Values for the previous pairs.
	values := []int{3, 4, 2, 5, 7, 1}

	// Empty map should not contain (R1, C1)

	for i, pair := range integerSetsPairs {
		// Define the sets.
		Ri := set.NewIntSetFromValues(pair[0])
		Cj := set.NewIntSetFromValues(pair[1])

		// Create a block and set a size.
		B := NewBlockFromIntSets(Ri, Cj)
		B.SetSize(values[i])

		// Block should not be on the map.
		if BM.Contains(Ri, Cj) {
			t.Errorf("Block Map should not contain the key (%v, %v)", Ri, Cj)
		}

		// Add block to the map.
		BM.Add(Ri, Cj, B)

		// Block should now be on the map.
		if !BM.Contains(Ri, Cj) {
			t.Errorf("Block Map should contain the key (%v, %v)", Ri, Cj)
		}

		// Test that the block's attributes were saved.
		X := BM.Get(Ri, Cj)
		if X.GetSize() != values[i] {
			t.Errorf("Expected block's size to be %v but got %v",
				values[i],
				X.GetSize())
		}
	}

}
