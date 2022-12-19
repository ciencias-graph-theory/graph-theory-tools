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
		// Obtain the expected refinement of the row.
		expectedC0 := set.NewIntSetFromValues(expectedRefinement[0])
		expectedC1 := set.NewIntSetFromValues(expectedRefinement[1])

		// Obtained column refinement of the row.
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

	// Expected row refinements for each row.
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

	// Test for each row, that the obtained column refinement is correct.
	for col, expectedRefinement := range expectedRowRefinement {
		// Obtain the expected refinement of the row.
		expectedR0 := set.NewIntSetFromValues(expectedRefinement[0])
		expectedR1 := set.NewIntSetFromValues(expectedRefinement[1])

		// Obtained column refinement of the first row.
		obtainedR1, obtainedR0 := rowRefinement(m, col, R, false)

		// Check that the column refinement is correct.
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
			// Columns and rows indexes sets.
			rowIndexesSet := set.NewIntSetFromValues(rowIndexes)
			colIndexesSet := set.NewIntSetFromValues(colIndexes)

			// Columns and row parts.
			rowPart := NewOrderedBipartitionFromIntSet(rowIndexesSet)
			colPart := NewOrderedBipartitionFromIntSet(colIndexesSet)

			// Block.
			B := NewBlock(rowPart, colPart)

			// Calculate block's size.
			calculateSize(m, B)
			sizeB := B.GetSize()

			// Test if block's size has been calculated correctly.
			if nonZeroEntries[ri][cj] != sizeB {
				t.Errorf("Expected %v but got %v", nonZeroEntries[ri][cj], sizeB)
			}

			// Test if the row blocks' size has been calculated correctly.
			for i, r := range rowIndexes {
				if nonZeroEntriesRows[ri][cj][i] != B.GetRowBlockSize(r) {
					t.Errorf("Expected %v but got %v",
						nonZeroEntriesRows[ri][cj][i],
						B.GetRowBlockSize(r))
				}
			}
		}
	}
}
