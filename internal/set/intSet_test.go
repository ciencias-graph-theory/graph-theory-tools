package set

import (
	"math/rand"
	"testing"
)

// Tests if NewIntegerSet returns an empty integer set.
func TestNewIntSet(t *testing.T) {
	A := NewIntSet()

	// The empty set's cardinality should be zero.
	if !A.IsEmpty() {
		t.Errorf("Expected an empty set, but got set of cardinality: %v",
			A.Cardinality())
	}

	// The empty set should not have any values.
	if len(A.GetValues()) != 0 {
		t.Errorf("Expected an empty set, but got a set with values: %v",
			A.GetValues())
	}
}

// Tests if NewIntSetFromValues returns an integer set which contains the
// given values.
func TestNewIntSetFromValues(t *testing.T) {
	// Test the function with a slice with no repetitions.
	sliceWithNoRepetitions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	A := NewIntSetFromValues(sliceWithNoRepetitions)

	// A's size should be the same that of the original slice.
	if A.Cardinality() != len(sliceWithNoRepetitions) {
		t.Errorf("Expected set of length %v, but got set of length: %v",
			len(sliceWithNoRepetitions),
			A.Cardinality())
	}

	// Check that A contains the slice's elements.
	for _, val := range sliceWithNoRepetitions {
		if !A.Contains(val) {
			t.Errorf("Expected set to contain %v", val)
		}
	}

	sliceWithRepetitions := []int{1, 1, 2, 2, 2, 3, 4, 5}

	// Test the function with a slice with repetitions.
	B := NewIntSetFromValues(sliceWithRepetitions)

	// B's size should be 5, as there are only 5 unique elements.
	if B.Cardinality() != 5 {
		t.Errorf("Expected set of length %v, but got set of length: %v",
			5, A.Cardinality())
	}

	// Check that B contains the slice's elements.
	for _, val := range sliceWithRepetitions {
		if !B.Contains(val) {
			t.Errorf("Expected set to contain %v", val)
		}
	}

}

// Given a random slice of integers, tests if NewIntSetFromValues returns an
// integer set which contains the given values.
func TestNewIntSetFromValuesRandom(t *testing.T) {
	// Defines a random length r.
	r := rand.Intn(100) + 1

	// Creates a slice of length r and fills it with random values.
	randomSlice := make([]int, r, r)

	for k := 0; k < r; k++ {
		randomSlice[k] = rand.Intn(100)
	}

	A := NewIntSetFromValues(randomSlice)

	for k := 0; k < r; k++ {
		if !A.Contains(randomSlice[k]) {
			t.Errorf("Expected to contain %v", randomSlice[k])
		}
	}

}

// Given a set, tests if Add effectively adds an element to the set.
func TestAdd(t *testing.T) {
	// Define an empty set.
	A := NewIntSet()

	if !A.IsEmpty() {
		t.Errorf("Expected an empty set, but got set of length: %v",
			A.Cardinality())
	}

	// Add several unique values.
	vals := []int{1, 2, 3, 4}
	l := 0
	for _, v := range vals {
		A.Add(v)
		l++

		// Check that the value was added correctly.
		if !A.Contains(v) {
			t.Errorf("Expected to contain %v", v)
		}
	}
}

// TestAddRandom tests the add function with a random integer.
func TestAddRandom(t *testing.T) {
	randomInt := rand.Intn(100)
	set := NewIntSet()
	isPresent := set.Add(randomInt)
	if isPresent {
		t.Errorf("Expected %v, got %v", false, isPresent)
	}
	if !set.Contains(randomInt) {
		t.Errorf("Expected %v, got %v", true, set.Contains(randomInt))
	}
	isPresent = set.Add(randomInt)
	if !isPresent {
		t.Errorf("Expected %v, got %v", true, isPresent)
	}
}

// TestCardinality tests if Cardinality effectively returns the set's length.
func TestCardinality(t *testing.T) {
	// Define an empty set.
	A := NewIntSet()

	if A.Cardinality() != 0 {
		t.Errorf("Expected an empty set, but got set of length: %v",
			A.Cardinality())
	}

	// Add several unique values.
	vals := []int{1, 2, 3, 4}
	l := 0
	for _, v := range vals {
		A.Add(v)
		l++

		// Check that the set's length matches with the elements added.
		if A.Cardinality() != l {
			t.Errorf("Expected set of length %v, but got set of length: %v",
				l, A.Cardinality())
		}
	}

	// Add the same values, the set should not be affected.
	for _, v := range vals {
		A.Add(v)

		// Check that the set's length didn't increase or decrease.
		if A.Cardinality() != l {
			t.Errorf("Expected set of length %v, but got set of length: %v",
				l, A.Cardinality())
		}
	}
}

// TestRemoveRandom tests the remove function with a random integer.
func TestRemoveRandom(t *testing.T) {
	randomInt := rand.Intn(100)
	set := NewIntSet()
	wasPresent := set.Remove(randomInt)
	if wasPresent {
		t.Errorf("Expected %v, got %v", false, wasPresent)
	}
	set.Add(randomInt)
	wasPresent = set.Remove(randomInt)
	if !wasPresent {
		t.Errorf("Expected %v, got %v", true, wasPresent)
	}
	if set.Contains(randomInt) {
		t.Errorf("Expected %v, got %v", false, set.Contains(randomInt))
	}
}

// TestRemove tests the IsEmpty function.
func TestIsEmpty(t *testing.T) {
	// Define an empty set.
	A := NewIntSet()

	if !A.IsEmpty() {
		t.Errorf("Expected an empty set, but got set of length: %v",
			A.Cardinality())
	}

	// Add an element to the set.
	A.Add(5)

	// The set should no longer be empty.
	if A.IsEmpty() {
		t.Errorf("Expected an non-empty set.")
	}
}

// TestRemoveRandom tests the IsEmpty function with a random integer.
func TestIsEmptyRandom(t *testing.T) {
	randomInt := rand.Intn(100)
	set := NewIntSet()
	if !set.IsEmpty() {
		t.Errorf("Expected %v, got %v", true, set.IsEmpty())
	}
	set.Add(randomInt)
	if set.IsEmpty() {
		t.Errorf("Expected %v, got %v", false, set.IsEmpty())
	}
	set.Remove(randomInt)
	if !set.IsEmpty() {
		t.Errorf("Expected %v, got %v", true, set.IsEmpty())
	}
}

// TestGetValues tests if GetValues returns the values contained in the random
// slice.
func TestGetValuesRandom(t *testing.T) {
	// Defines a random length r.
	r := rand.Intn(100) + 1

	// Creates a slice of length r and fills it with random values.
	randomSlice := make([]int, r, r)

	for k := 0; k < r; k++ {
		randomSlice[k] = rand.Intn(100)
	}

	A := NewIntSetFromValues(randomSlice)

	avals := A.GetValues()

	// Check that the values in A are actually contained in the original slice.
	for _, aval := range avals {
		found := false
		for _, orig := range randomSlice {
			if aval == orig {
				found = true
			}
		}

		if !found {
			t.Errorf("Set contains value %v but this is not present in original slice.",
				aval)
		}
	}

}

// TestEquals tests the Equals function.
func TestEquals(t *testing.T) {
	// Define two empty sets.
	A := NewIntSet()
	B := NewIntSet()

	// Two empty sets should be equal.
	if !A.Equals(B) {
		t.Errorf("Two empty sets should be equal.")
		t.Errorf("Expected %v but got %v", A, B)
	}

	// After adding one element, they should differ.
	A.Add(5)

	if A.Equals(B) {
		t.Errorf("The following should not be equal: %v, %v", A, B)
	}

	// After adding a different element, they should still differ.
	B.Add(6)

	if A.Equals(B) {
		t.Errorf("The following should not be equal: %v, %v", A, B)
	}

	// After adding the missing elements, they should be equal now.
	A.Add(6)
	B.Add(5)

	if !A.Equals(B) {
		t.Errorf("Expected %v but got %v", A, B)
	}

	// A set should be equal to itself.
	if !A.Equals(A) {
		t.Errorf("A set should be equal to itself.")
	}

}
