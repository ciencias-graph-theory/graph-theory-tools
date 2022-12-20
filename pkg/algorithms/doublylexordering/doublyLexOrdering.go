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

// calculateSize calculates the size of the block B = (Ri, Cj) and the blocks
// (r, Cj) for each r in Ri.
func calculateSize(M matrix, Ri, Cj *IntSet) (int, map[int]int) {
	// Number of rows in the block.
	numRows := Ri.Cardinality()

	// Get the indexes contained in the row and column part.
	rowIndexes := Ri.GetValues()
	colIndexes := Cj.GetValues()

	// Define a map for the row blocks.
	rowBlocks := make(map[int]int, numRows)

	// Total size of B.
	sizeB := 0

	// Traverse each row r in Ri and calculate the size of the block (r, Cj).
	for _, r := range rowIndexes {
		// Total size of (r, Cj).
		sizeR := 0

		for _, c := range colIndexes {
			sizeR += int(M[r][c])
		}

		rowBlocks[r] = sizeR

		sizeB += sizeR
	}

	return sizeB, rowBlocks
}

// getSplittingRow returns the index of a splitting row of B, if B has no
// splitting row then it returns -1. This function requires that the block's
// size has already been calculated. A splitting row of a block B = (R_i, C_j)
// is a row r in R_i such that the row slice M(r, C_j) is non-constant.
func getSplittingRow(M matrix, B *Block) int {
	// Get the indexes contained in the row part.
	rowIndexes := B.GetRowPart().GetSet().GetValues()

	// Get the amount of columns in the block.
	numCols := B.GetColumnPart().GetSet().Cardinality()

	// Traverse each row block in search of a non-constant one.
	for _, r := range rowIndexes {
		if (B.GetRowBlockSize(r) > 0) && (B.GetRowBlockSize(r) < numCols) {
			return r
		}
	}

	// If all row blocks are constant then there is no splitting row.
	return -1
}

// updateAffectedBlocksColumns updates the size of the blocks affected by a
// column refinement.
func updateAffectedBlocksColumns(M matrix, lRef, rRef *IntSet, B *Block, sizeMap *BlockMap) {
	// SC is the smaller partition of the refinement, BC is the bigger partition
	// of the refinement.
	var SC, BC *IntSet

	// Indicate which partition is smaller.
	if lRef.Cardinality() <= rRef.Cardinality() {
		SC, BC = lRef, rRef
	} else {
		SC, BC = rRef, lRef
	}

	// Columns indexes set of the current block.
	Cj := B.GetColumnPart().GetSet()

	// Update all of the affected blocks.
	rowPart := B.GetRowPart()
	for rowPart != nil {
		// Row indexes set of the current block.
		R := rowPart.GetSet()

		// Current block defined by (R, Cj)
		current := sizeMap.Get(R, Cj)

		// Define a block for the smaller refinement.
		smallerBlock := NewBlockFromIntSets(R, SC)

		// Calculate the small block's size.
		size, rowBlocksMap := calculateSize(M, R, SC)

		// Set small block's size and row blocks' sizes.
		smallerBlock.SetSize(size)
		smallerBlock.SetRowBlocksMap(rowBlocksMap)

		// Define a block for the bigger refinement.
		biggerBlock := NewBlockFromIntSets(R, BC)

		// Calculate the bigger block's size using the smaller one.
		sizeBg := 0
		for _, r := range R.GetValues() {
			// Set the bigger block's row blocks using the following formula.
			// size(r, BC) = size(r, Cj) - size(r, SC)
			currentRowSize := current.GetRowBlockSize(r)
			smallerRowSize := smallerBlock.GetRowBlockSize(r)
			biggerRowSize := currentRowSize - smallerRowSize
			biggerBlock.SetRowBlockSize(r, biggerRowSize)

			sizeBg += biggerRowSize
		}

		// Set the bigger block's size.
		biggerBlock.SetSize(sizeBg)

		// Add the blocks (R, SC) and (R, BC) to the map.
		sizeMap.Add(R, SC, smallerBlock)
		sizeMap.Add(R, BC, biggerBlock)

		// Move to the next row.
		rowPart = rowPart.GetNext()
	}
}
