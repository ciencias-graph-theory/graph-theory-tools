package set

import (
	"math/rand"
	"testing"
)

// TestAdd tests the add function.
func TestAdd(t *testing.T) {
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

// TestRemove tests the remove function.
func TestRemove(t *testing.T) {
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
