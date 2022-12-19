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

	// Expected column refinement of the first row.
	expectedC0 := set.NewIntSetFromValues([]int{0, 2})

	expectedC1 := set.NewIntSetFromValues([]int{1, 3, 4})

	// Obtained column refinement of the first row.
	obtainedC1, obtainedC0 := columnRefinement(m, 0, C, false)

	// Check that the column refinement is correct.
	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 0, C, true)

	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the second row.
	expectedC0 = set.NewIntSetFromValues([]int{0, 2, 3, 4})

	expectedC1 = set.NewIntSetFromValues([]int{1})

	// Obtained column refinement of the second row.
	obtainedC1, obtainedC0 = columnRefinement(m, 1, C, false)

	// Check that the column refinement is correct.
	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 1, C, true)

	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the third row.
	expectedC0 = set.NewIntSetFromValues([]int{0, 1, 3})

	expectedC1 = set.NewIntSetFromValues([]int{2, 4})

	// Obtained column refinement of the third row.
	obtainedC1, obtainedC0 = columnRefinement(m, 2, C, false)

	// Check that the column refinement is correct.
	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 2, C, true)

	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the fourth row.
	expectedC0 = set.NewIntSetFromValues([]int{0, 1, 2})

	expectedC1 = set.NewIntSetFromValues([]int{3, 4})

	// Obtained column refinement of the fourth row.
	obtainedC1, obtainedC0 = columnRefinement(m, 3, C, false)

	// Check that the column refinement is correct.
	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 3, C, true)

	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the fifth row.
	expectedC0 = set.NewIntSetFromValues([]int{0, 1, 2, 3, 4})

	expectedC1 = set.NewIntSet()

	// Obtained column refinement of the fourth row.
	obtainedC1, obtainedC0 = columnRefinement(m, 4, C, false)

	// Check that the column refinement is correct.
	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 4, C, true)

	if !expectedC0.Equals(obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !expectedC1.Equals(obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
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

	// Expected row refinement of the first column.
	expectedR0 := set.NewIntSetFromValues([]int{0, 1, 2, 3, 4})

	expectedR1 := set.NewIntSet()

	// Obtained row refinement of the first column.
	obtainedR1, obtainedR0 := rowRefinement(m, 0, R, false)

	// Check that the row refinement is correct.
	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 0, R, true)

	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the second column.
	expectedR0 = set.NewIntSetFromValues([]int{2, 3, 4})

	expectedR1 = set.NewIntSetFromValues([]int{0, 1})

	// Obtained row refinement of the second column.
	obtainedR1, obtainedR0 = rowRefinement(m, 1, R, false)

	// Check that the row refinement is correct.
	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 1, R, true)

	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the third column.
	expectedR0 = set.NewIntSetFromValues([]int{0, 1, 3, 4})

	expectedR1 = set.NewIntSetFromValues([]int{2})

	// Obtained row refinement of the third column.
	obtainedR1, obtainedR0 = rowRefinement(m, 2, R, false)

	// Check that the row refinement is correct.
	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 2, R, true)

	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the fourth column.
	expectedR0 = set.NewIntSetFromValues([]int{1, 2, 4})

	expectedR1 = set.NewIntSetFromValues([]int{0, 3})

	// Obtained row refinement of the fourth column.
	obtainedR1, obtainedR0 = rowRefinement(m, 3, R, false)

	// Check that the row refinement is correct.
	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 3, R, true)

	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the fifth column.
	expectedR0 = set.NewIntSetFromValues([]int{1, 4})

	expectedR1 = set.NewIntSetFromValues([]int{0, 2, 3})

	// Obtained row refinement of the fifth column.
	obtainedR1, obtainedR0 = rowRefinement(m, 4, R, false)

	// Check that the row refinement is correct.
	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 4, R, true)

	if !expectedR0.Equals(obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !expectedR1.Equals(obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
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

	// Columns and rows indexes.
	rowIndexes := []int{0, 1, 2, 3, 4}
	colIndexes := []int{0, 1, 2, 3, 4}

	// Columns and rows indexes sets.
	rowIndexesSet := set.NewIntSetFromValues(rowIndexes)
	colIndexesSet := set.NewIntSetFromValues(colIndexes)

	// Columns and row parts.
	rowPart := NewOrderedBipartitionFromIntSet(rowIndexesSet)
	colPart := NewOrderedBipartitionFromIntSet(colIndexesSet)

	// Block.
	B := NewBlock(rowPart, colPart)

	// Number of non-zero entries in B.
	nonZeroEntries := 8
	nonZeroEntriesRows := []int{3, 1, 2, 2, 0}

	// Calculate block's size.
	calculateSize(m, B)
	sizeB := B.GetSize()

	// Test if block's size has been calculated correctly.
	if nonZeroEntries != sizeB {
		t.Errorf("Expected %v but got %v", nonZeroEntries, sizeB)
	}

	// Test if the row blocks' size has been calculated correctly.
	for i, r := range rowIndexes {
		if nonZeroEntriesRows[i] != B.GetRowBlockSize(r) {
			t.Errorf("Expected %v but got %v",
				nonZeroEntriesRows[i],
				B.GetRowBlockSize(r))
		}
	}
}
