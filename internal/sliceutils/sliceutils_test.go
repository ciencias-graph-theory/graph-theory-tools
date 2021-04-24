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
