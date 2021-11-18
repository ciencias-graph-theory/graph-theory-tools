// Package utils provides functions frequently used in go.
package sliceutils

import (
	"bytes"
)

// This function checks whether two int slices are equal.
func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// This function checks whether two 2D byte slices are equal.
func EqualByteMatrix(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !bytes.Equal(v, b[i]) {
			return false
		}
	}
	return true
}

// NextNonZero returns the position of the next non-zero entry to the right from
// a given position in a slice, starting over if the end of the slice is
// reached.   If no non-zero entry is found (different from the initial
// position), -1 is returned.
func NextNonZero(a []byte, i int) int {
	l := len(a)
	var j int
	if i+1 < l {
		j = i + 1
	} else {
		j = 0
	}
	for a[j] == 0 && j != i {
		if j+1 < l {
			j++
		} else {
			j = 0
		}
	}
	if j == i {
		return -1
	}
	return j
}

// Foldl takes a function, and initial argument, and a slice. It applies
// the function with the initial argument as the first parameter and the
// first element of the slice as the second parameter. The result is then
// operated with the second element of the slice and so on.
// TODO: Make the function usable for any type.
func Foldl(f func(int, int) int, init int, s []int) int {
	rv := init
	for _, v := range s {
		rv = f(rv, v)
	}
	return rv
}

// AddIntSlice adds the elements of a slice of ints.
func SumIntSlice(s []int) int {
	return Foldl(func(a, b int) int { return a + b }, 0, s)
}
