package orderings

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

// Let M be a matrix with indexed rows and columns, and R = (R_1, ..., R_n) and
// C = (C_1, ..., C_n) a partition of the indexed rows and columns respectively.
// A block is defined as an ordered pair, B = (R_i, C_j).
type block = [2]numset

// Given a square (0,1)-matrix M, a row index, and a set of columns Cj of M. The
// following function returns a column refinement of the given set Cj. A column
// refinement divides the columns whose entries in the given row are 0 or 1. If
// the inverseOrder flag is true then the first set contains all of the columns
// whose entries are 0.
func columnRefinement(M matrix, row int, Cj numset, inverseOrder bool) (numset, numset) {
	// The set of columns whose entries in the given row are 0.
	C0 := make(numset)

	// The set of columns whose entries in the given row are 0.
	C1 := make(numset)

	for col, _ := range Cj {
		if M[row][col] == 1 {
			C1[col] = true
		} else {
			C0[col] = true
		}
	}

	if inverseOrder {
		return C0, C1
	} else {
		return C1, C0
	}

}
