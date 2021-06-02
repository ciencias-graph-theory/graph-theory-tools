package sequencegen

import (
	"reflect"
	"testing"
)

func contains(s [][]bool, e []bool) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func TestWeight(t *testing.T) {
	size := 3
	// Test 1: weight 0
	gen := Weight(size, 0)
	res := make([]bool, size)
	got := gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	// Test 2: weight 1
	got2 := make([][]bool, size)
	gen = Weight(size, 1)
	for i := 0; i < size; i++ {
		got2[i] = make([]bool, size)
		copy(got2[i], gen())
	}
	elem := []bool{true, false, false}
	found := contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	elem = []bool{false, true, false}
	found = contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	elem = []bool{false, false, true}
	found = contains(got2, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got2)
	}
	// Test 3: weight 2
	got3 := make([][]bool, size)
	gen = Weight(size, 2)
	for i := 0; i < size; i++ {
		got3[i] = make([]bool, size)
		copy(got3[i], gen())
	}
	elem = []bool{false, true, true}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	elem = []bool{true, false, true}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	elem = []bool{true, true, false}
	found = contains(got3, elem)
	if !found {
		t.Errorf("%v is not in sequence %v", elem, got3)
	}
	// Test 4: Weight 3
	gen = Weight(size, 3)
	res = []bool{true, true, true}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
	got = gen()
	if !reflect.DeepEqual(res, got) {
		t.Errorf("Expected %v, got %v", res, got)
	}
}
