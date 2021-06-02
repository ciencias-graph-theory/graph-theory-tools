package sequencegen

import (
	"reflect"
	"testing"
)

// contains tells us whether a matrix contains a specific row.
func contains(s [][]bool, e []bool) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

// TestWeight is a test for the Weight function.
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

// TestBinary is a test for the Binary function.
func TestBinary(t *testing.T) {
	gen := Binary(5)
	want := make([][]bool, 32)
	want[0] = []bool{true, true, true, true, true}
	for i := 1; i < 6; i++ {
		want[i] = []bool{true, true, true, true, true}
		want[i][i-1] = false
	}
	want[6] = []bool{false, false, true, true, true}
	want[7] = []bool{false, true, false, true, true}
	want[8] = []bool{false, true, true, false, true}
	want[9] = []bool{false, true, true, true, false}
	want[10] = []bool{true, false, false, true, true}
	want[11] = []bool{true, false, true, false, true}
	want[12] = []bool{true, false, true, true, false}
	want[13] = []bool{true, true, false, false, true}
	want[14] = []bool{true, true, false, true, false}
	want[15] = []bool{true, true, true, false, false}
	want[16] = []bool{true, true, false, false, false}
	want[17] = []bool{true, false, true, false, false}
	want[18] = []bool{true, false, false, true, false}
	want[19] = []bool{true, false, false, false, true}
	want[20] = []bool{false, true, true, false, false}
	want[21] = []bool{false, true, false, true, false}
	want[22] = []bool{false, true, false, false, true}
	want[23] = []bool{false, false, true, true, false}
	want[24] = []bool{false, false, true, false, true}
	want[25] = []bool{false, false, false, true, true}
	for i := 26; i < 31; i++ {
		want[i] = []bool{false, false, false, false, false}
		want[i][i-26] = true
	}
	want[31] = []bool{false, false, false, false, false}
	got := make([][]bool, 32)
	for i := 0; i < 32; i++ {
		got[i] = make([]bool, 5)
		copy(got[i], gen())
	}

	for i := 0; i < 32; i++ {
		if !contains(got, want[i]) {
			t.Errorf("%v is not in sequence %v", want[i], got)
		}
	}

}
