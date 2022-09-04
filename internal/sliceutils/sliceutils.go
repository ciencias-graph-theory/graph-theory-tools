// Package utils provides functions frequently used in go.
package sliceutils

import (
	"bytes"
	"math"
	"strconv"
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

// This function checks whether two int slices are equal.
func EqualByteSlice(a, b []byte) bool {
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

// Given a slice v of length m and a number n, we extend v until its length is a
// multiple of n by appending zeros. If leftPadding is true then the zeros are
// appended to the left; otherwise, they're appended to the right.
func ExtendByteSlice(v []byte, n int, leftPadding bool) []byte {
	m := len(v)

	if (m % n) == 0 {
		return v
	} else {
		missingBytes := n - (m % n)
		u := make([]byte, missingBytes)

		if leftPadding {
			return append(u, v...)
		} else {
			return append(v, u...)
		}
	}
}

// Divides the slice into groups of n bits; it is expected that the slice's
// length is a multiple of n.
func DivideByteSlice(v []byte, n int) [][]byte {
	numGroups := len(v) / n
	groups := make([][]byte, numGroups)

	for i := 0; i < numGroups; i++ {
		groups[i] = v[(i * n):((i + 1) * n)]
	}

	return groups
}

// Not to be confused with the methods provided by the library encoding/binary.
// The reason is the following: If a byte slice represents the binary of an int
// (e.g. (1, 1, 0, 0, 1, 1)) then the desired int would be 35. However, all the
// libraries available to work with bytes would consider each bit as a byte, not
// as a bit. Resulting in a wrong conversion. So, we have to do this manually.
// Consider this as an auxiliary function to the ByteMatrixtoIntSlice.
func ByteSliceToInt(v []byte) int {
	val := 0
	numBits := len(v)

	for i := 0; i < numBits; i++ {
		if v[i] == 1 {
			power := float64((numBits - 1) - i)
			val += int(math.Pow(2, power))
		}
	}

	return val
}

// Returns an int slice containing the int representation of the bytes in each
// row of the matrix given.
func ByteMatrixToIntSlice(v [][]byte) []int {
	cols := len(v)
	vals := make([]int, cols)

	for i := 0; i < cols; i++ {
		vals[i] = ByteSliceToInt(v[i])
	}

	return vals
}

// Returns a ASCII representation of the slice of ints. Note: The int's
// values should between 33 and 126, the reason being that all of the ASCII
// found in this range are printable.
func IntSliceToASCII(v []int) string {
	ASCII := ""

	for _, val := range v {
		ASCII += string(rune(val))
	}

	return ASCII
}

// Returns the binary representation of an intenger n as a byte slice.
// E.g. If n = 170 then the function returns {1,0,1,0,1,0,1,0}.
func IntToByteSlice(n int) []byte {
	binaryString := strconv.FormatInt(int64(n), 2)
	binary := make([]byte, len(binaryString))

	for i := 0; i < len(binaryString); i++ {
		if binaryString[i] == '1' {
			binary[i] = 1
		} else {
			binary[i] = 0
		}
	}

	return binary
}

// Returns a byte slice corresponding to the upper triangle of a byte matrix
// column by column, if diagonal is true then the diagonal is considered.
func ByteMatrixUpperTriangle(matrix [][]byte, diagonal bool) []byte {
	var v []byte

	// Number of columns.
	n := len(matrix)

	// Let k be the amount of elements in the upper triangle of the matrix,
	// we build a byte slice of size k.
	// If the diagonal is not considered then k = (n * (n-1)) / 2.
	// If the diagonal is considered then k = (n * (n+1)) / 2.
	if diagonal {
		v = make([]byte, (n * (n + 1) / 2))
	} else {
		v = make([]byte, (n * (n - 1) / 2))
	}

	// Travel the matrix column by column filling the slice.
	k := 0
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if (i == j) && !diagonal {
				continue
			}
			v[k] = matrix[i][j]
			k++
		}
	}

	return v
}

// Sum m to each value k in the int slice v.
func IntSliceSumToEach(v []int, m int) []int {
	w := make([]int, len(v))

	for i := 0; i < len(v); i++ {
		w[i] = v[i] + m
	}

	return w
}

// Returns the byte matrix as a slice of bytes, this is done by traversing
// the matrix row by row.
func ByteMatrixToSlice(matrix [][]byte) []byte {
	// A matrix has m rows and n columns.
	m := len(matrix)
	n := len(matrix[0])

	// We create a slice of size (m * n)
	slice := make([]byte, m*n)

	// Travel the matrix row by row.
	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			slice[k] = matrix[i][j]
			k++
		}
	}

	return slice
}
