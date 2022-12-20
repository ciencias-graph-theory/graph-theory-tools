package doublylexordering

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
	"testing"
)

// This function calls columnRefinement with a matrix, a row index and a set of
// columns and tests if the set of columns have been partitioned correctly.
func TestColumnRefinement(t *testing.T) {
	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	// The columns indexes.
	C := set.NewIntSetFromValues([]int{0, 1, 2, 3, 4})

	// Expected column refinements for each row.
	expectedColRefinement := [][][]int{
		// First row.
		{{0, 2}, {1, 3, 4}},
		// Second row.
		{{0, 2, 3, 4}, {1}},
		// Third row.
		{{0, 1, 3}, {2, 4}},
		// Fourth row.
		{{0, 1, 2}, {3, 4}},
		// Fifth row.
		{{0, 1, 2, 3, 4}, {}},
	}

	// Test for each row, that the obtained column refinement is correct.
	for row, expectedRefinement := range expectedColRefinement {
		// Expected column refinement.
		expectedC0 := set.NewIntSetFromValues(expectedRefinement[0])
		expectedC1 := set.NewIntSetFromValues(expectedRefinement[1])

		// Obtained column refinement.
		obtainedC1, obtainedC0 := columnRefinement(m, row, C, false)

		// Check that the column refinement is correct.
		if !expectedC0.Equals(obtainedC0) {
			t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
		}

		if !expectedC1.Equals(obtainedC1) {
			t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
		}

		// Check that the inverse order flag works.
		obtainedC0, obtainedC1 = columnRefinement(m, row, C, true)

		if !expectedC0.Equals(obtainedC0) {
			t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
		}

		if !expectedC1.Equals(obtainedC1) {
			t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
		}
	}

}

// This function calls rowRefinement with a matrix, a column index and a set of
// columns and tests if the set of rows have been partitioned correctly.
func TestRowRefinement(t *testing.T) {
	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	// The rows indexes.
	R := set.NewIntSetFromValues([]int{0, 1, 2, 3, 4})

	// Expected row refinements for each column.
	expectedRowRefinement := [][][]int{
		// First row.
		{{0, 1, 2, 3, 4}, {}},
		// Second row.
		{{2, 3, 4}, {0, 1}},
		// Third row.
		{{0, 1, 3, 4}, {2}},
		// Fourth row.
		{{1, 2, 4}, {0, 3}},
		// Fifth row.
		{{1, 4}, {0, 2, 3}},
	}

	// Test for each column, that the obtained row refinement is correct.
	for col, expectedRefinement := range expectedRowRefinement {
		// Expected row refinement.
		expectedR0 := set.NewIntSetFromValues(expectedRefinement[0])
		expectedR1 := set.NewIntSetFromValues(expectedRefinement[1])

		// Obtained row refinement.
		obtainedR1, obtainedR0 := rowRefinement(m, col, R, false)

		// Check that the row refinement is correct.
		if !expectedR0.Equals(obtainedR0) {
			t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
		}

		if !expectedR1.Equals(obtainedR1) {
			t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
		}

		// Check that the inverse order flag works.
		obtainedR0, obtainedR1 = rowRefinement(m, col, R, true)

		if !expectedR0.Equals(obtainedR0) {
			t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
		}

		if !expectedR1.Equals(obtainedR1) {
			t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
		}
	}
}

// This function calls calculateSize with a matrix, and a block. Tests if
// block's size has been calculated correctly.
func TestCalculateSize(t *testing.T) {
	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	// Multiple columns and rows indexes.
	rowIndexesTests := [][]int{
		{0, 1, 2, 3, 4},
		{0, 1, 3},
		{4},
	}
	colIndexesTests := [][]int{
		{0, 1, 2, 3, 4},
		{0, 1, 2},
		{0},
	}

	// Number of non-zero entries in the block conformed by the row and column
	// indexes.
	nonZeroEntries := [][]int{
		{8, 3, 0},
		{6, 2, 0},
		{0, 0, 0},
	}

	nonZeroEntriesRows := [][][]int{
		{{3, 1, 2, 2, 0}, {1, 1, 1, 0, 0}, {0, 0, 0, 0, 0}},
		{{3, 1, 2}, {1, 1, 0}, {0, 0, 0}},
		{{0}, {0}, {0}},
	}

	for ri, rowIndexes := range rowIndexesTests {
		for cj, colIndexes := range colIndexesTests {
			// Set Ri and Cj of row and column indexes respectively.
			Ri := set.NewIntSetFromValues(rowIndexes)
			Cj := set.NewIntSetFromValues(colIndexes)

			// Calculate block B=(Ri, Cj) size.
			sizeB, rowBlocksMapB := calculateSize(m, Ri, Cj)

			// Test if block's size has been calculated correctly.
			if nonZeroEntries[ri][cj] != sizeB {
				t.Errorf("Expected %v but got %v", nonZeroEntries[ri][cj], sizeB)
			}

			// Test if the row blocks' size has been calculated correctly.
			for i, r := range rowIndexes {
				if nonZeroEntriesRows[ri][cj][i] != rowBlocksMapB[r] {
					t.Errorf("Expected %v but got %v",
						nonZeroEntriesRows[ri][cj][i],
						rowBlocksMapB[r])
				}
			}
		}
	}
}

// This function calls obtainSplittingRow with a matrix, and a block. Tests if
// it returns a splitting row of a block.
func TestObtainSplittingRow(t *testing.T) {
	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	// Pairs of rows and columns indexes.
	pairIndexes := [][][]int{
		{{0, 1, 2, 3, 4, 5}, {0, 1, 2, 3, 4, 5}},
		{{0, 1, 4, 5}, {0, 1, 2, 3, 4, 5}},
		{{0, 3}, {3, 4}},
		{{4, 5}, {0, 1, 2, 3, 4, 5}},
	}

	// Possible splitting rows of the previous pairs.
	allSplittingRows := [][]int{
		{0, 1, 2, 3},
		{0, 1},
		{},
		{},
	}

	for i, pair := range pairIndexes {
		// Obtain the row and column indexes.
		Ri := set.NewIntSetFromValues(pair[0])
		Cj := set.NewIntSetFromValues(pair[1])

		// Create the respective block.
		B := NewBlockFromIntSets(Ri, Cj)

		// Calculate B's size.
		sizeB, rowBlocksMapB := calculateSize(m, Ri, Cj)

		// Set B size and row blocks.
		B.SetSize(sizeB)
		B.SetRowBlockMap(rowBlocksMapB)

		// Get the possible splitting rows as a set.
		possibleSplittingRows := set.NewIntSetFromValues(allSplittingRows[i])

		// Get B's splitting row index.
		r := getSplittingRow(m, B)

		// If B has a splitting row, check it is a possible splitting row.
		if r != -1 {
			if !possibleSplittingRows.Contains(r) {
				t.Errorf("Row %v should not be a splitting row", r)
			}
		}

		// If B has no splitting rows, check that there are no possible splitting
		// rows.
		if r == -1 {
			if !possibleSplittingRows.IsEmpty() {
				t.Errorf("Expected a splitting row.")
			}
		}
	}

}

// TestUpdateAffectedBlocksColumns tests if the function updates correctly the
// size of the affected blocks by a column refinement.
func TestUpdateAffectedBlocksColumns(t *testing.T) {
	// Let's suppose we have the following matrix:
	//
	// {0, 1, 0, 1, 1},
	// {0, 1, 0, 0, 0},
	// {0, 0, 1, 0, 1},
	// {0, 0, 0, 1, 1},
	// {0, 0, 0, 0, 0},
	//
	// With the following partition.
	//
	//       C1
	// ┌─────────────┐
	// │0, 1, 0, 1, 1│ R1
	// └─────────────┘
	// ┌─────────────┐
	// │0, 1, 0, 0, 0│ R2
	// │0, 0, 1, 0, 1│
	// └─────────────┘
	// ┌─────────────┐
	// │0, 0, 0, 1, 1│ R3
	// │0, 0, 0, 0, 0│
	// └─────────────┘
	//
	// Given the following column refinement {0, 1, 2} {3, 4}, the matrix is left like
	// this:
	//
	//    C1       C2
	// ┌───────┐ ┌────┐
	// │0, 1, 0│ │1, 1│ R1
	// └───────┘ └────┘
	// ┌───────┐ ┌────┐
	// │0, 1, 0│ │0, 0│ R2
	// │0, 0, 1│ │0, 1│
	// └───────┘ └────┘
	// ┌───────┐ ┌────┐
	// │0, 0, 0│ │1, 1│ R3
	// │0, 0, 0│ │0, 0│
	// └───────┘ └────┘
	//
	// Then the size of the affected blocks are the following (from left to right
	// and top to bottom): (1, 2, 2, 1, 0, 2).
	//
	// The following test simulates the previous example.

	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	// Define row parts.
	R1 := NewOrderedBipartitionFromIntSlice([]int{0})
	R2 := NewOrderedBipartitionFromIntSlice([]int{1, 2})
	R3 := NewOrderedBipartitionFromIntSlice([]int{3, 4})

	// Define order. (Only the next values matter here)
	R1.SetNext(R2)
	R2.SetNext(R3)

	// Define column part.
	C := NewOrderedBipartitionFromIntSlice([]int{0, 1, 2, 3, 4})

	// Define a block map.
	sizeMap := NewBlockMap()

	// Define the size of the blocks (R1, C), (R2, C), (R3, C).
	B1 := NewBlockFromPartitions(R1, C)
	B2 := NewBlockFromPartitions(R2, C)
	B3 := NewBlockFromPartitions(R3, C)

	B1.SetSize(3)
	B1.SetRowBlockSize(0, 3)
	B2.SetSize(3)
	B2.SetRowBlockSize(1, 1)
	B2.SetRowBlockSize(2, 2)
	B3.SetSize(2)
	B3.SetRowBlockSize(3, 2)
	B3.SetRowBlockSize(4, 0)

	// Add the previous blocks to the map.
	sizeMap.Add(R1.GetSet(), C.GetSet(), B1)
	sizeMap.Add(R2.GetSet(), C.GetSet(), B2)
	sizeMap.Add(R3.GetSet(), C.GetSet(), B3)

	// Define column refinement.
	C1 := set.NewIntSetFromValues([]int{0, 1, 2})
	C2 := set.NewIntSetFromValues([]int{3, 4})

	// Update the blocks affected by the refinement, this should produce the
	// following blocks:
	// (R1, C1) - size: 1,
	// (R1, C2) - size: 2,
	// (R2, C1) - size: 2,
	// (R2, C2) - size: 1,
	// (R3, C1) - size: 0,
	// (R3, C2) - size: 2,
	updateAffectedBlocksColumns(m, C1, C2, B1, sizeMap)

	// Rows and columns partitions.
	rowPartitions := [][]int{
		{0}, {1, 2}, {3, 4},
	}

	colPartitions := [][]int{
		{0, 1, 2}, {3, 4},
	}

	// Sizes of the produces blocks.
	expectedSizes := []int{1, 2, 2, 1, 0, 2}

	// Expected row blocks of each produced block.
	expectedRowBlocks := [][][]int{
		{{1}, {2}},
		{{1, 1}, {0, 1}},
		{{0, 0}, {2, 0}},
	}

	// Build every posible block and check if its size is correct.
	k := 0
	for i, rp := range rowPartitions {
		for j, cp := range colPartitions {

			// Create integer sets to contain the indexes.
			Ri := set.NewIntSetFromValues(rp)
			Cj := set.NewIntSetFromValues(cp)

			// If the block was not added, return error.
			if !sizeMap.Contains(Ri, Cj) {
				t.Errorf("Expected map to contain block (%v, %v)", Ri, Cj)
			}

			// Get the block defined by (Ri, Cj).
			B := sizeMap.Get(Ri, Cj)

			// Check if block's size is the expected.
			if B.GetSize() != expectedSizes[k] {
				t.Errorf("Expected block (%v, %v) size to be %v but got %v",
					Ri, Cj, expectedSizes[k], B.GetSize())
			}

			// Check if row blocks size are correct.
			for l, r := range rp {
				if B.GetRowBlockSize(r) != expectedRowBlocks[i][j][l] {
					t.Errorf("Expected block (%v, %v) size to be %v but got %v",
						r, Cj, expectedRowBlocks[i][j][l], B.GetRowBlockSize(r))
				}
			}

			k++
		}
	}
}

// TestUpdateAffectedBlocksRows tests if the function updates correctly the
// size of the affected blocks by a column refinement.
func TestUpdateAffectedBlocksRows(t *testing.T) {
	// Let's suppose we have the following matrix:
	//
	// {0, 1, 0, 1, 1},
	// {0, 1, 0, 0, 0},
	// {0, 0, 1, 0, 1},
	// {0, 0, 0, 1, 1},
	// {0, 0, 0, 0, 0},
	//
	// With the following partition.
	//
	//     C1    C2    C3
	//    ┌────┐┌────┐┌─┐
	//    │0, 1││1, 1││1│ r0
	//    │0, 1││0, 0││0│ r1
	// R1 │0, 0││1, 0││1│ r2
	//    │0, 0││0, 1││1│ r3
	//    │0, 0││0, 0││0│ r4
	//    └────┘└────┘└─┘
	//
	// Let's suppose the following row partition is given:
	// {0, 1, 4}, {2, 3}
	//
	// Then the row refinement would produce the following blocks:
	//
	//     C1    C2    C3
	//    ┌────┐┌────┐┌─┐
	// R1 │0, 1││1, 1││1│ r0
	//    │0, 1││0, 0││0│ r1
	//    │0, 0││0, 0││0│ r4
	//    └────┘└────┘└─┘
	//    ┌────┐┌────┐┌─┐
	// R2 │0, 0││1, 0││1│ r2
	//    │0, 0││0, 1││1│ r3
	//    └────┘└────┘└─┘
	//
	// Then the size of the affected blocks are the following (from left to right
	// and top to bottom): (2, 2, 1, 0, 2, 2).
	//
	// The following test simulates the previous example.

	// Arbitrary matrix.
	m := [][]byte{
		{0, 1, 0, 1, 1},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	// Define column parts.
	C1 := NewOrderedBipartitionFromIntSlice([]int{0, 1})
	C2 := NewOrderedBipartitionFromIntSlice([]int{2, 3})
	C3 := NewOrderedBipartitionFromIntSlice([]int{4})

	// Define order. (Only the next values matter here).
	C1.SetNext(C2)
	C2.SetNext(C3)

	// Define row part.
	R := NewOrderedBipartitionFromIntSlice([]int{0, 1, 2, 3, 4})

	// Define a block map.
	sizeMap := NewBlockMap()

	// Define the of the blocks (R, C1), (R, C2), (R, C3)
	B1 := NewBlockFromPartitions(R, C1)
	B2 := NewBlockFromPartitions(R, C2)
	B3 := NewBlockFromPartitions(R, C3)

	// Define the size of the previous blocks.
	B1.SetSize(2)
	B1.SetRowBlockSize(0, 1)
	B1.SetRowBlockSize(1, 1)
	B1.SetRowBlockSize(2, 0)
	B1.SetRowBlockSize(3, 0)
	B1.SetRowBlockSize(4, 0)

	B2.SetSize(4)
	B2.SetRowBlockSize(0, 2)
	B2.SetRowBlockSize(1, 0)
	B2.SetRowBlockSize(2, 1)
	B2.SetRowBlockSize(3, 1)
	B2.SetRowBlockSize(4, 0)

	B3.SetSize(3)
	B3.SetRowBlockSize(0, 1)
	B3.SetRowBlockSize(1, 0)
	B3.SetRowBlockSize(2, 1)
	B3.SetRowBlockSize(3, 1)
	B3.SetRowBlockSize(4, 0)

	// Add the previous blocks to the map.
	sizeMap.Add(R.GetSet(), C1.GetSet(), B1)
	sizeMap.Add(R.GetSet(), C2.GetSet(), B2)
	sizeMap.Add(R.GetSet(), C3.GetSet(), B3)

	// Define row refinement.
	R1 := set.NewIntSetFromValues([]int{0, 1, 4})
	R2 := set.NewIntSetFromValues([]int{2, 3})

	// Update the blocks affected by the refinement.
	updateAffectedBlocksRows(m, R1, R2, B1, sizeMap)

	// Rows and columns partitions.
	rowPartitions := [][]int{
		{0, 1, 4}, {2, 3},
	}

	colPartitions := [][]int{
		{0, 1}, {2, 3}, {4},
	}

	// Sizes of the produces blocks.
	expectedSizes := []int{2, 2, 1, 0, 2, 2}

	// Expected row blocks of each produced block.
	expectedRowBlocks := [][][]int{
		{{1, 1, 0}, {2, 0, 0}, {1, 0, 0}},
		{{0, 0}, {1, 1}, {1, 1}},
	}

	// Build every posible block and check if its size is correct.
	k := 0
	for i, rp := range rowPartitions {
		for j, cp := range colPartitions {

			// Create integer sets to contain the indexes.
			Ri := set.NewIntSetFromValues(rp)
			Cj := set.NewIntSetFromValues(cp)

			// If the block was not added, return error.
			if !sizeMap.Contains(Ri, Cj) {
				t.Errorf("Expected map to contain block (%v, %v)", Ri, Cj)
			}

			// Get the block defined by (Ri, Cj).
			B := sizeMap.Get(Ri, Cj)

			// Check if block's size is the expected.
			if B.GetSize() != expectedSizes[k] {
				t.Errorf("Expected block (%v, %v) size to be %v but got %v",
					Ri, Cj, expectedSizes[k], B.GetSize())
			}

			// Check if row blocks size are correct.
			for l, r := range rp {
				if B.GetRowBlockSize(r) != expectedRowBlocks[i][j][l] {
					t.Errorf("Expected block (%v, %v) size to be %v but got %v",
						r, Cj, expectedRowBlocks[i][j][l], B.GetRowBlockSize(r))
				}
			}

			k++
		}
	}

}
