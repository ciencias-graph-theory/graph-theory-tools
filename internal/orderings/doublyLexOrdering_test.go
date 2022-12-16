package orderings

import (
	"testing"
)

// This functions tests whether two num sets are equal or not.
func equalNumSets(m1, m2 map[int]bool) bool {
	if len(m1) != len(m2) {
		return false
	}

	for val, _ := range m1 {
		if !m2[val] {
			return false
		}
	}

	return true
}

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
	C := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	// Expected column refinement of the first row.
	expectedC0 := map[int]bool{
		0: true,
		2: true,
	}

	expectedC1 := map[int]bool{
		1: true,
		3: true,
		4: true,
	}

	// Obtained column refinement of the first row.
	obtainedC1, obtainedC0 := columnRefinement(m, 0, C, false)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 0, C, true)

	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the second row.
	expectedC0 = map[int]bool{
		0: true,
		2: true,
		3: true,
		4: true,
	}

	expectedC1 = map[int]bool{
		1: true,
	}

	// Obtained column refinement of the second row.
	obtainedC1, obtainedC0 = columnRefinement(m, 1, C, false)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	obtainedC0, obtainedC1 = columnRefinement(m, 1, C, true)

	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the third row.
	expectedC0 = map[int]bool{
		0: true,
		1: true,
		3: true,
	}

	expectedC1 = map[int]bool{
		2: true,
		4: true,
	}

	// Obtained column refinement of the third row.
	obtainedC1, obtainedC0 = columnRefinement(m, 2, C, false)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 2, C, true)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the fourth row.
	expectedC0 = map[int]bool{
		0: true,
		1: true,
		2: true,
	}

	expectedC1 = map[int]bool{
		3: true,
		4: true,
	}

	// Obtained column refinement of the fourth row.
	obtainedC1, obtainedC0 = columnRefinement(m, 3, C, false)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 3, C, true)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Expected column refinement of the fourth row.
	expectedC0 = map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	expectedC1 = map[int]bool{}

	// Obtained column refinement of the fourth row.
	obtainedC1, obtainedC0 = columnRefinement(m, 4, C, false)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
		t.Errorf("Expected %v but got %v", expectedC1, obtainedC1)
	}

	// Check that the inverse order flag works.
	obtainedC0, obtainedC1 = columnRefinement(m, 4, C, true)

	// Check that the column refinement is correct.
	if !equalNumSets(expectedC0, obtainedC0) {
		t.Errorf("Expected %v but got %v", expectedC0, obtainedC0)
	}

	if !equalNumSets(expectedC1, obtainedC1) {
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
	R := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	// Expected row refinement of the first column.
	expectedR0 := map[int]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	expectedR1 := map[int]bool{}

	// Obtained row refinement of the first column.
	obtainedR1, obtainedR0 := rowRefinement(m, 0, R, false)

	// Check that the row refinement is correct.
	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 0, R, true)

	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the second column.
	expectedR0 = map[int]bool{
		2: true,
		3: true,
		4: true,
	}

	expectedR1 = map[int]bool{
		0: true,
		1: true,
	}

	// Obtained row refinement of the second column.
	obtainedR1, obtainedR0 = rowRefinement(m, 1, R, false)

	// Check that the row refinement is correct.
	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 1, R, true)

	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the third column.
	expectedR0 = map[int]bool{
		0: true,
		1: true,
		3: true,
		4: true,
	}

	expectedR1 = map[int]bool{
		2: true,
	}

	// Obtained row refinement of the third column.
	obtainedR1, obtainedR0 = rowRefinement(m, 2, R, false)

	// Check that the row refinement is correct.
	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 2, R, true)

	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the fourth column.
	expectedR0 = map[int]bool{
		1: true,
		2: true,
		4: true,
	}

	expectedR1 = map[int]bool{
		0: true,
		3: true,
	}

	// Obtained row refinement of the fourth column.
	obtainedR1, obtainedR0 = rowRefinement(m, 3, R, false)

	// Check that the row refinement is correct.
	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 3, R, true)

	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Expected row refinement of the fifth column.
	expectedR0 = map[int]bool{
		1: true,
		4: true,
	}

	expectedR1 = map[int]bool{
		0: true,
		2: true,
		3: true,
	}

	// Obtained row refinement of the fifth column.
	obtainedR1, obtainedR0 = rowRefinement(m, 4, R, false)

	// Check that the row refinement is correct.
	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}

	// Check that the inverse order flag works.
	obtainedR0, obtainedR1 = rowRefinement(m, 4, R, true)

	if !equalNumSets(expectedR0, obtainedR0) {
		t.Errorf("Expected %v but got %v", expectedR0, obtainedR0)
	}

	if !equalNumSets(expectedR1, obtainedR1) {
		t.Errorf("Expected %v but got %v", expectedR1, obtainedR1)
	}
}
