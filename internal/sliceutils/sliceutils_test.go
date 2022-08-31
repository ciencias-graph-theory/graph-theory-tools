package sliceutils

import (
	"math/rand"
	"testing"
)

// TestFoldl tests the function Foldl.
func TestFoldl(t *testing.T) {
	pow2 := make([]int, 20)
	for i := 0; i < 20; i++ {
		pow2[i] = 2
	}
	one := Foldl(func(a, b int) int { return a / b }, 1048576, pow2)
	if one != 1 {
		t.Errorf("Expected %d, got %d", 1, one)
	}
	empty := []int{}
	value := rand.Intn(100)
	res := Foldl(func(a, b int) int { return 0 }, value, empty)
	if res != value {
		t.Errorf("Expected %d, got %d", res, value)
	}

}

// TestSumIntSlice tests the function SumIntSlice.
func TestSumIntSlice(t *testing.T) {
	empty := []int{}
	same := []int{2, 2, 2, 2, 2, 2, 2}
	gauss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumEmpty := SumIntSlice(empty)
	sumSame := SumIntSlice(same)
	sumGauss := SumIntSlice(gauss)
	if sumEmpty != 0 {
		t.Errorf("Expected %d, got %d", 0, sumEmpty)
	}
	if sumSame != 2*len(same) {
		t.Errorf("Expected %d, got %d", 2*len(same), sumSame)
	}
	if sumGauss != len(gauss)*(len(gauss)+1)/2 {
		t.Errorf("Expected %d, got %d",
			len(gauss)*(len(gauss)+1)/2,
			sumGauss)
	}
}

// TestExtendByteSlice calls extendByteSlice with a slice v of length m and a
// number n, it extends the slice by appending zeros to the left until its
// length is a multiple of n.
func TestExtendSliceOfBytes(t *testing.T) {
	// Example vectors.
	av := []byte{1, 1, 1, 1, 1, 1}
	bv := []byte{1, 1, 0, 0, 1, 1}
	cv := []byte{1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1}
	dv := []byte{1, 1, 0, 0, 0, 1, 1}

	// Extended vectors.
	xa := []byte{1, 1, 1, 1, 1, 1}
	xb := []byte{1, 1, 0, 0, 1, 1}
	xc := []byte{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1}
	xd := []byte{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1}

	// Obtained vectors.
	Av := ExtendByteSlice(av, 6)
	Bv := ExtendByteSlice(bv, 6)
	Cv := ExtendByteSlice(cv, 6)
	Dv := ExtendByteSlice(dv, 6)

	// Check that the obtainded vectors are equal to the extended ones.
	if !EqualByteSlice(Av, xa) {
		t.Errorf("Expansion error: Expected %v but got %v", xa, Av)
	}

	if !EqualByteSlice(Bv, xb) {
		t.Errorf("Expansion error: Expected %v but got %v", xb, Bv)
	}

	if !EqualByteSlice(Cv, xc) {
		t.Errorf("Expansion error: Expected %v but got %v", xc, Cv)
	}
	if !EqualByteSlice(Dv, xd) {
		t.Errorf("Expansion error: Expected %v but got %v", xd, Dv)
	}
}
