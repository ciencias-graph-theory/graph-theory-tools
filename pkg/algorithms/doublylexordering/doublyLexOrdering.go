package doublylexordering

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
)

// A double lexicographical ordering of a matrix is an ordering in which the
// rows and columns, as vectors, are sorted lexicographically. For example, the
// double lexicographical ordering of this matrix would be the following:
//      c0  c1  c2  c3            c3  c1  c2  c0
// r0 [  0,  1,  0,  1 ]     r0 [  1,  1,  0,  0 ]
// r1 [  0,  1,  0,  0 ]     r3 [  1,  0,  0,  0 ]
// r2 [  0,  0,  1,  0 ]     r1 [  0,  1,  0,  0 ]
// r3 [  0,  0,  0,  1 ]     r2 [  0,  0,  1,  0 ]
//         Original                   Sorted
//
// To further explain the previous example. If each row is represented as a
// vector and they're compared from top to bottom, then we have the following
// vectors in the given order: (1100, 1000, 0100, 0010). If each column in
// represented as a vector and they're compared from left to right, then we have
// the following vectors in the given order: (1100, 1010, 0001, 0000). As we can
// observe, the rows and columns are now sorted lexicographically. One important
// detail to mention is that switching rows is independant from switching
// columns; in other words, we can switch rows without swithing columns and
// vice versa.

// Type aliases to improve readability.
type matrix = [][]byte

// There is no built-in data structure to represent a set. To get around this we
// will define a set as a boolean map where all the elements inside the map will
// always point to a true.
type numset = map[int]bool

// Given a square (0,1)-matrix M, a row index, and a set of columns Cj of M. The
// following function returns a column refinement of the given set Cj. A column
// refinement divides the columns whose entries in the given row are 0 or 1. If
// the inverseOrder flag is true then the first set contains all of the columns
// whose entries are 0.
func columnRefinement(M matrix, row int, Cj *IntSet, inverseOrder bool) (*IntSet, *IntSet) {
	// The set of columns whose entries in the given row are 0.
	C0 := set.NewIntSet()

	// The set of columns whose entries in the given row are 1.
	C1 := set.NewIntSet()

	// Get the columns indexes contained in Cj.
	colIndexes := Cj.GetValues()

	// Traverse the columns given by Cj.
	for _, col := range colIndexes {
		// Given the row r and column c, if the entry M[r][c] is one, add it to C1,
		// otherwise to C0.
		if M[row][col] == 1 {
			C1.Add(col)
		} else {
			C0.Add(col)
		}
	}

	// Return the partition of the set of column indexes.
	if inverseOrder {
		return C0, C1
	} else {
		return C1, C0
	}

}

// Given a square (0,1)-matrix M, a column index, and a set of rows Ri of M. The
// following function returns a row refinement of the given set Ri. A row
// refinement divides the rows whose entries in the given column are 0 or 1. If
// the inverseOrder flag is true then the first set contains all of the rows
// whose entries are 0.
func rowRefinement(M matrix, col int, Ri *IntSet, inverseOrder bool) (*IntSet, *IntSet) {
	// The set of rows whose entries in the given column are 0.
	R0 := set.NewIntSet()

	// The set of rows whose entries in the given column are 1.
	R1 := set.NewIntSet()

	// Get rows indexes contained in Ri.
	rowIndexes := Ri.GetValues()

	// Traverse the rows given by Ri.
	for _, row := range rowIndexes {
		// Given the row r and column c, if the entry M[r][c] is one, add it to R1,
		// otherwise to R0.
		if M[row][col] == 1 {
			R1.Add(row)
		} else {
			R0.Add(row)
		}
	}

	// Return the partition of the set of row indexes.
	if inverseOrder {
		return R0, R1
	} else {
		return R1, R0
	}

}

// The size of a block B = (Ri, Cj) is the amount of non-zero entries in
// sub-matrix defined by B. Given a square (0,1)-matrix M, a a set of rows Ri
// and columns Cj of M, and a map size. The following function calculates the
// size of the block B = (Ri, Cj) and the blocks (r, Cj) for each r in Ri.
// func calculateSize(M matrix, Ri, Cj numset, size map[block]int) {
// 	// Define the block B = (Ri, Cj).
// 	B := block{Ri, Cj}
//
// 	// Size of B.
// 	sizeB := 0
//
// 	// Traverse each row r in Ri and calculate the size of the block (r, Cj).
// 	for r, _ := range Ri {
// 		// Define a unitary set for r.
// 		rUnit := numset{r: true}
//
// 		// Define block (r, Cj).
// 		rowBlock := block{rUnit, Cj}
//
// 		// Size of (r, Cj).
// 		sizeR := 0
//
// 		for c, _ := range Cj {
// 			sizeR += int(M[r][c])
// 		}
//
// 		size[rowBlock] = sizeR
//
// 		sizeB += sizeR
// 	}
//
// 	size[B] = sizeB
// }