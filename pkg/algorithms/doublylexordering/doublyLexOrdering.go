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
	rowIndexes := B.GetRows().GetValues()

	// Get the amount of columns in the block.
	numCols := B.GetColumns().Cardinality()

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
func updateAffectedBlocksColumns(
	M matrix, lRef, rRef *IntSet, Cj *IntSet,
	currentPosition int, sizeMap *BlockMap,
	orderedRows []*IntSet) {
	// SC is the smaller partition of the refinement, BC is the bigger partition
	// of the refinement.
	var SC, BC *IntSet

	// Indicate which partition is smaller.
	if lRef.Cardinality() <= rRef.Cardinality() {
		SC, BC = lRef, rRef
	} else {
		SC, BC = rRef, lRef
	}

	// Traverse all the affected blocks.
	for i := 0; i < len(orderedRows); i++ {
		Ri := orderedRows[i]

		// Obtain block B = (Ri, Cj)
		B := sizeMap.Get(Ri, Cj)

		// Let k be the current row part position. All the blocks above B = (Rk, Cj)
		// are constant. Do not consider the block B' = (Ri, Cjl) for all i < k, as
		// they are above (Rk, Cjl) then they are not considered; however blocks
		// above B'' = (Ri, Cjr) have to be considered, and thus must have a size,
		// as they're constant this operation is not heavy on performance.
		if i < currentPosition {
			constBlockL := NewBlockFromIntSets(Ri, lRef)
			constBlockR := NewBlockFromIntSets(Ri, rRef)
			if B.GetSize() > 0 {
				constBlockL.SetSize(Ri.Cardinality() * lRef.Cardinality())
				constBlockR.SetSize(Ri.Cardinality() * rRef.Cardinality())
			} else {
				constBlockL.SetSize(0)
				constBlockR.SetSize(0)
			}

			sizeMap.Add(Ri, rRef, constBlockL)
			sizeMap.Add(Ri, lRef, constBlockR)
		} else {
			// For all blocks B = (Ri, Cj) where i >= k. The size of the blocks B' =
			// (Ri, Cjl) and B'' = (Ri, Cjr) can be determined using only the smallest
			// refinement. This speeds ups the process detemining the size of produced
			// blocks by the refinement.
			smallBlock := NewBlockFromIntSets(Ri, SC)
			bigBlock := NewBlockFromIntSets(Ri, BC)

			// Determine the size of the small block.
			sizeSmall, rowBlockMapSmall := calculateSize(M, Ri, SC)
			smallBlock.SetSize(sizeSmall)
			smallBlock.SetRowBlockMap(rowBlockMapSmall)

			// Use the previous information to determine the size of the big block.
			sizeBig := 0
			for _, r := range Ri.GetValues() {
				currentRowSize := B.GetRowBlockSize(r)
				smallRowSize := smallBlock.GetRowBlockSize(r)
				bigRowSize := currentRowSize - smallRowSize
				bigBlock.SetRowBlockSize(r, bigRowSize)

				sizeBig += bigRowSize
			}

			bigBlock.SetSize(sizeBig)

			// Add blocks B' and B'' to the map.
			sizeMap.Add(Ri, SC, smallBlock)
			sizeMap.Add(Ri, BC, bigBlock)
		}

	}
}

// updateAffectedBlocksRows updates the size of the blocks affected by a
// row refinement.
func updateAffectedBlocksRows(
	M matrix, lRef, rRef *IntSet, Ri *IntSet,
	sizeMap *BlockMap, pendingColParts []*IntSet) {
	// SC is the smaller partition of the refinement, BC is the bigger partition
	// of the refinement.
	var SR, BR *IntSet

	// Indicate which partition is smaller.
	if lRef.Cardinality() <= rRef.Cardinality() {
		SR, BR = lRef, rRef
	} else {
		SR, BR = rRef, lRef
	}

	// In the refinement of columns we have to take into consideration all of the
	// new produced blocks to the right, even if these blocks are constant, we
	// have to save their size to be consistent. In the refinement of rows this
	// doesn't happen as all the produced blocks to the left are already constant
	// and don't have to traversed again; thus only the blocks next to (Ri, Cj)
	// have to be traversed.
	for _, Cj := range pendingColParts {
		// Get block B = (Ri, Cj).
		B := sizeMap.Get(Ri, Cj)

		// The size of the blocks B' = (Ril, Cj) and B'' = (Rir, Cj) can be
		// determined using only the smallest refinement. This speeds ups the
		// process detemining the size of produced blocks by the refinement.
		smallBlock := NewBlockFromIntSets(SR, Cj)
		bigBlock := NewBlockFromIntSets(BR, Cj)

		// Determine the small block's size by traversing the row blocks of (Ri, Cj)
		// that will be (SR, Cj).
		sizeSmall := 0
		for _, r := range SR.GetValues() {
			rbs := B.GetRowBlockSize(r)
			smallBlock.SetRowBlockSize(r, rbs)
			sizeSmall += rbs
		}

		smallBlock.SetSize(sizeSmall)
		bigBlock.SetSize(B.GetSize() - sizeSmall)

		// To avoid traversing all of the rows in BR, simply indicate that the row
		// blocks map of (BR, C) is the same as (Ri, C).
		bigBlock.SetRowBlockMap(B.GetRowBlockMap())

		// Add the blocks (SR, Cj) and (BR, Cj) to the map.
		sizeMap.Add(SR, Cj, smallBlock)
		sizeMap.Add(BR, Cj, bigBlock)

	}
}
