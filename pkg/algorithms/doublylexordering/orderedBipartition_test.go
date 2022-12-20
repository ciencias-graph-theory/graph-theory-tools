package doublylexordering

import (
	"testing"
)

// TestSetPartition tests the function SetPartition.
func TestSetPartition(t *testing.T) {
	// Define the following partition.
	I := []int{0, 1, 2, 3, 4, 5, 6, 7}
	P := NewOrderedBipartitionFromIntSlice(I)

	// Set P as the start.
	P.SetStart()

	// Check P has no partition defined.
	if P.HasPartition() {
		t.Errorf("Expected %v no have no partition.", P)
	}

	// Check P has no previous or next part.
	if P.GetPrevious() != nil {
		t.Errorf("Expected %v no have no previous partition but found %v.",
			P, P.GetPrevious())
	}

	if P.GetNext() != nil {
		t.Errorf("Expected %v no have no next partition but found %v.",
			P, P.GetNext())
	}

	// Define a left and right partition.
	L := NewOrderedBipartitionFromIntSlice([]int{0, 2, 4, 6})
	R := NewOrderedBipartitionFromIntSlice([]int{1, 3, 5, 7})

	P.SetPartition(L, R)

	// Check the order is correct.
	// The left partition shouldn't have a previous part.
	if L.GetPrevious() != nil {
		t.Errorf("Expected %v no have no previous partition but found %v.",
			L, L.GetPrevious())
	}

	// The left partition's next part should be the right partition.
	if !R.Equals(L.GetNext()) {
		t.Errorf("Expected %v to be to the right of %v but found %v.",
			R, L, L.GetNext())
	}

	// The right part's previous part should be the left partition..
	if !L.Equals(R.GetPrevious()) {
		t.Errorf("Expected %v to be to the left of %v but found %v.",
			L, R, R.GetPrevious())
	}

	// The right partition shouldn't have a next part.
	if R.GetNext() != nil {
		t.Errorf("Expected %v no have no previous partition but found %v.",
			R, R.GetNext())
	}

	// The left partition should be at the start of the ordered partition.
	if !L.IsStart() {
		t.Errorf("Expected %v to be at the start.", L)
	}

}
