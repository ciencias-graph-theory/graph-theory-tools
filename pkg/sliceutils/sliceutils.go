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
