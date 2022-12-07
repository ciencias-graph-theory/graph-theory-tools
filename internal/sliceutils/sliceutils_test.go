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

	cv := []byte{
		1, 1, 0, 0, 1, 1, 1,
		0, 0, 0, 0, 0, 1, 0,
		1, 0, 1, 0, 0, 1, 0,
		0, 0, 0, 1, 0, 1, 1}

	dv := []byte{1, 1, 0, 0, 0, 1, 1}

	// Extended vectors with left padding.
	xal := []byte{1, 1, 1, 1, 1, 1}

	xbl := []byte{1, 1, 0, 0, 1, 1}

	xcl := []byte{
		0, 0, 1, 1, 0, 0,
		1, 1, 1, 0, 0, 0,
		0, 0, 1, 0, 1, 0,
		1, 0, 0, 1, 0, 0,
		0, 0, 1, 0, 1, 1}

	xdl := []byte{
		0, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 1, 1}

	// Extended vectors with right padding.
	xar := []byte{1, 1, 1, 1, 1, 1}

	xbr := []byte{1, 1, 0, 0, 1, 1}

	xcr := []byte{
		1, 1, 0, 0, 1, 1,
		1, 0, 0, 0, 0, 0,
		1, 0, 1, 0, 1, 0,
		0, 1, 0, 0, 0, 0,
		1, 0, 1, 1, 0, 0}

	xdr := []byte{
		1, 1, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 0}

	// Obtained vectors.
	Avl := ExtendByteSlice(av, 6, true)
	Bvl := ExtendByteSlice(bv, 6, true)
	Cvl := ExtendByteSlice(cv, 6, true)
	Dvl := ExtendByteSlice(dv, 6, true)

	Avr := ExtendByteSlice(av, 6, false)
	Bvr := ExtendByteSlice(bv, 6, false)
	Cvr := ExtendByteSlice(cv, 6, false)
	Dvr := ExtendByteSlice(dv, 6, false)

	// Check that the obtainded vectors are equal to the extended ones.
	if !EqualByteSlice(Avl, xal) {
		t.Errorf("Expansion error: Expected %v but got %v", xal, Avl)
	}

	if !EqualByteSlice(Bvl, xbl) {
		t.Errorf("Expansion error: Expected %v but got %v", xbl, Bvl)
	}

	if !EqualByteSlice(Cvl, xcl) {
		t.Errorf("Expansion error: Expected %v but got %v", xcl, Cvl)
	}
	if !EqualByteSlice(Dvl, xdl) {
		t.Errorf("Expansion error: Expected %v but got %v", xdl, Dvl)
	}

	if !EqualByteSlice(Avr, xar) {
		t.Errorf("Expansion error: Expected %v but got %v", xar, Avr)
	}

	if !EqualByteSlice(Bvr, xbr) {
		t.Errorf("Expansion error: Expected %v but got %v", xbr, Bvr)
	}

	if !EqualByteSlice(Cvr, xcr) {
		t.Errorf("Expansion error: Expected %v but got %v", xcr, Cvr)
	}
	if !EqualByteSlice(Dvr, xdr) {
		t.Errorf("Expansion error: Expected %v but got %v", xdr, Dvr)
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
	Av, errA := DivideByteSlice(xa, 6)
	Bv, errB := DivideByteSlice(xb, 6)
	Cv, errC := DivideByteSlice(xc, 6)
	Dv, errD := DivideByteSlice(xd, 6)

	// The previous operations should not return an error.
	if errA != nil {
		t.Errorf("Division Error: Unexpected Error: %v", errA)
	}

	if errB != nil {
		t.Errorf("Division Error: Unexpected Error: %v", errB)
	}

	if errC != nil {
		t.Errorf("Division Error: Unexpected Error: %v", errC)
	}

	if errD != nil {
		t.Errorf("Division Error: Unexpected Error: %v", errD)
	}

	// Check that the obtainded vectors are equal to the extended ones.
	if !EqualByteMatrix(Av, ga) {
		t.Errorf("Division error: Expected %v but got %v", ga, Av)
	}

	if !EqualByteMatrix(Bv, gb) {
		t.Errorf("Division error: Expected %v but got %v", gb, Bv)
	}

	if !EqualByteMatrix(Cv, gc) {
		t.Errorf("Division error: Expected %v but got %v", gc, Cv)
	}

	if !EqualByteMatrix(Dv, gd) {
		t.Errorf("Division error: Expected %v but got %v", gd, Dv)
	}

	// The following operation should return an error.
	Ev, errE := DivideByteSlice(xa[1:], 6)

	if errE == nil {
		t.Errorf("Division error: Expected an error but got %v", Ev)
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

// TestIntSliceToASCII calls IntSliceToASCII with a slice v, and return an ASCII
// representation of the values of the slice.
func TestIntSliceToASCII(t *testing.T) {
	// Examples.
	a := []int{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75}
	b := []int{71, 65, 84, 84, 79, 95, 82, 85, 76, 69, 83}
	c := []int{59, 60, 61, 62, 63, 64}

	// Expected string.
	sa := "ABCDEFGHIJK"
	sb := "GATTO_RULES"
	sc := ";<=>?@"

	// Obtained strings.
	SA := IntSliceToASCII(a)
	SB := IntSliceToASCII(b)
	SC := IntSliceToASCII(c)

	// Check that the obtained strings are correct.
	if SA != sa {
		t.Errorf("Conversion error: Expected %v but got %v", sa, SA)
	}

	if SB != sb {
		t.Errorf("Conversion error: Expected %v but got %v", sb, SB)
	}

	if SC != sc {
		t.Errorf("Conversion error: Expected %v but got %v", sc, SC)
	}
}

func TestASCIIToIntSlice(t *testing.T) {
	// String examples.
	sa := "ABCDEFGHIJK"
	sb := "GATTO_RULES"
	sc := ";<=>?@"
	sd := "SG_WAS_HERE"

	// Expected int slices.
	a := []int{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75}
	b := []int{71, 65, 84, 84, 79, 95, 82, 85, 76, 69, 83}
	c := []int{59, 60, 61, 62, 63, 64}
	d := []int{83, 71, 95, 87, 65, 83, 95, 72, 69, 82, 69}

	// Obtained int slices.
	A := ASCIIToIntSlice(sa)
	B := ASCIIToIntSlice(sb)
	C := ASCIIToIntSlice(sc)
	D := ASCIIToIntSlice(sd)

	// Compare slices.
	if !EqualIntSlice(A, a) {
		t.Errorf("Conversion error: Expected %v but got %v", a, A)
	}

	if !EqualIntSlice(B, b) {
		t.Errorf("Conversion error: Expected %v but got %v", b, B)
	}

	if !EqualIntSlice(C, c) {
		t.Errorf("Conversion error: Expected %v but got %v", a, A)
	}

	if !EqualIntSlice(D, d) {
		t.Errorf("Conversion error: Expected %v but got %v", a, A)
	}
}

// TestIntToByteSlice calls IntToByteslice with an int n. Returns the binary
// representation of n as a byte slice.
func TestIntToByteSlice(t *testing.T) {
	// Expected byte slices.
	b170 := []byte{1, 0, 1, 0, 1, 0, 1, 0}
	b63 := []byte{1, 1, 1, 1, 1, 1}
	b35 := []byte{1, 0, 0, 0, 1, 1}
	b89 := []byte{1, 0, 1, 1, 0, 0, 1}
	b120 := []byte{1, 1, 1, 1, 0, 0, 0}

	// Obtained byte slices.
	B170 := IntToByteSlice(170)
	B63 := IntToByteSlice(63)
	B35 := IntToByteSlice(35)
	B89 := IntToByteSlice(89)
	B120 := IntToByteSlice(120)

	// Check that the obtained byte slices are correct.
	if !EqualByteSlice(B170, b170) {
		t.Errorf("Conversion error: Expected %v but got %v", b170, B170)
	}

	if !EqualByteSlice(B63, b63) {
		t.Errorf("Conversion error: Expected %v but got %v", b63, B63)
	}

	if !EqualByteSlice(B35, b35) {
		t.Errorf("Conversion error: Expected %v but got %v", b35, B35)
	}

	if !EqualByteSlice(B89, b89) {
		t.Errorf("Conversion error: Expected %v but got %v", b89, B89)
	}

	if !EqualByteSlice(B120, b120) {
		t.Errorf("Conversion error: Expected %v but got %v", b120, B120)
	}
}

// TestByteMatrixUpperTriangle calls ByteMatrixUpperTriangle with a byte matrix,
// and then compares the obtained vector with the upper triangle of the given
// matrix.
func TestByteMatrixUpperTriangle(t *testing.T) {

	// Complete graph with four vertices.
	a := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// A 4-cycle.
	b := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// A cube.
	c := [][]byte{
		{0, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 0},
	}

	// The corresponding vectors of the upper triangle of each adjacency matrix
	// without considering the diagonal.
	av := []byte{1, 1, 1, 1, 1, 1}
	bv := []byte{1, 1, 0, 0, 1, 1}
	cv := []byte{
		1, 1, 0, 0, 1, 1, 1,
		0, 0, 0, 0, 0, 1, 0,
		1, 0, 1, 0, 0, 1, 0,
		0, 0, 0, 1, 0, 1, 1,
	}

	Av := ByteMatrixUpperTriangle(a, false)
	Bv := ByteMatrixUpperTriangle(b, false)
	Cv := ByteMatrixUpperTriangle(c, false)

	if !EqualByteSlice(Av, av) {
		t.Errorf("Formatting error: Expected %v but got %v", av, Av)
	}

	if !EqualByteSlice(Bv, bv) {
		t.Errorf("Formatting error: Expected %v but got %v", bv, Bv)
	}

	if !EqualByteSlice(Cv, cv) {
		t.Errorf("Formatting error: Expected %v but got %v", cv, Cv)
	}

	// The corresponding vectors of the upper triangle of each adjacency matrix
	// without considering the diagonal.
	avd := []byte{0, 1, 0, 1, 1, 0, 1, 1, 1, 0}
	bvd := []byte{0, 1, 0, 1, 0, 0, 0, 1, 1, 0}
	cvd := []byte{
		0, 1, 0, 1, 0, 0, 0,
		1, 1, 0, 1, 0, 0, 0,
		0, 0, 0, 1, 0, 1, 0,
		0, 1, 0, 0, 1, 0, 0,
		0, 0, 0, 1, 0, 1, 1, 0,
	}

	Avd := ByteMatrixUpperTriangle(a, true)
	Bvd := ByteMatrixUpperTriangle(b, true)
	Cvd := ByteMatrixUpperTriangle(c, true)

	if !EqualByteSlice(Avd, avd) {
		t.Errorf("Formatting error: Expected %v but got %v", avd, Avd)
	}

	if !EqualByteSlice(Bvd, bvd) {
		t.Errorf("Formatting error: Expected %v but got %v", bvd, Bvd)
	}

	if !EqualByteSlice(Cvd, cvd) {
		t.Errorf("Formatting error: Expected %v but got %v", cvd, Cvd)
	}
}

// TestByteMatrixToSlice calls ByteMatrixToSlice with a byte matrix, and then
// compares that the obtained vector is correct.
func TestByteMatrixToSlice(t *testing.T) {
	// Test matrices.
	a := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// A 4-cycle.
	b := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// A cube.
	c := [][]byte{
		{0, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 1, 0, 1, 1, 0},
	}

	// Expected byte slices.
	sa := []byte{0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0}
	sb := []byte{0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0}
	sc := []byte{
		0, 1, 1, 0, 1, 0, 0, 0,
		1, 0, 0, 1, 0, 0, 1, 0,
		1, 0, 0, 1, 0, 1, 0, 0,
		0, 1, 1, 0, 0, 0, 0, 1,
		1, 0, 0, 0, 0, 1, 1, 0,
		0, 0, 1, 0, 1, 0, 0, 1,
		0, 1, 0, 0, 1, 0, 0, 1,
		0, 0, 0, 1, 0, 1, 1, 0,
	}

	// Obtained byte slices.
	SA := ByteMatrixToSlice(a)
	SB := ByteMatrixToSlice(b)
	SC := ByteMatrixToSlice(c)

	// Check that the obtainded slices are correct.
	if !EqualByteSlice(sa, SA) {
		t.Errorf("Conversion error: Expected %v but got %v", sa, SA)
	}

	if !EqualByteSlice(sb, SB) {
		t.Errorf("Conversion error: Expected %v but got %v", sb, SB)
	}

	if !EqualByteSlice(sc, SC) {
		t.Errorf("Conversion error: Expected %v but got %v", sc, SC)
	}
}
