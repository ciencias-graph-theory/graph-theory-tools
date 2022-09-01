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

// TestDivideByteslice calls DivideByteslice with a slice v and a number n, it
// divides v into groups of n bits; it is expected that v's length is multiple
// of n.
func TestDivideByteSlice(t *testing.T) {
	// Example vectors.
	xa := []byte{1, 1, 1, 1, 1, 1}
	xb := []byte{1, 1, 0, 0, 1, 1}
	xc := []byte{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1}
	xd := []byte{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1}

	// Groups.
	ga := [][]byte{{1, 1, 1, 1, 1, 1}}
	gb := [][]byte{{1, 1, 0, 0, 1, 1}}
	gc := [][]byte{
		{0, 0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1},
	}

	gd := [][]byte{
		{0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1},
	}

	// Obtained vectors.
	Av := DivideByteSlice(xa, 6)
	Bv := DivideByteSlice(xb, 6)
	Cv := DivideByteSlice(xc, 6)
	Dv := DivideByteSlice(xd, 6)

	// Check that the obtainded vectors are equal to the extended ones.
	if !EqualByteMatrix(Av, ga) {
		t.Errorf("Expansion error: Expected %v but got %v", ga, Av)
	}

	if !EqualByteMatrix(Bv, gb) {
		t.Errorf("Expansion error: Expected %v but got %v", gb, Bv)
	}

	if !EqualByteMatrix(Cv, gc) {
		t.Errorf("Expansion error: Expected %v but got %v", gc, Cv)
	}

	if !EqualByteMatrix(Dv, gd) {
		t.Errorf("Expansion error: Expected %v but got %v", gd, Dv)
	}
}

// TestDivideByteslice calls DivideByteslice with a slice v and a number n, it
// divides v into groups of n bits; it is expected that v's length is multiple
// of n.
func TestByteMatrixToIntSlice(t *testing.T) {
	// Example groups.
	ga := [][]byte{{1, 1, 1, 1, 1, 1}}
	gb := [][]byte{{1, 1, 0, 0, 1, 1}}
	gc := [][]byte{
		{0, 0, 1, 1, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1},
	}

	gd := [][]byte{
		{0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1, 1},
	}

	ai := []int{63}
	bi := []int{51}
	ci := []int{12, 56, 10, 36, 11}
	di := []int{1, 35}

	// Obtained vectors.
	Ai := ByteMatrixToIntSlice(ga)
	Bi := ByteMatrixToIntSlice(gb)
	Ci := ByteMatrixToIntSlice(gc)
	Di := ByteMatrixToIntSlice(gd)

	// Check that the obtainded vectors are equal to the extended ones.
	if !EqualIntSlice(Ai, ai) {
		t.Errorf("Conversion error: Expected %v but got %v", ai, Ai)
	}

	if !EqualIntSlice(Bi, bi) {
		t.Errorf("Conversion error: Expected %v but got %v", bi, Bi)
	}

	if !EqualIntSlice(Ci, ci) {
		t.Errorf("Conversion error: Expected %v but got %v", ci, Ci)
	}

	if !EqualIntSlice(Di, di) {
		t.Errorf("Conversion error: Expected %v but got %v", di, Di)
	}
}
