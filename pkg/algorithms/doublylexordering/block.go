package doublylexordering

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
)

// Let M be a square (0,1)-matrix with indexed rows and columns, R = (R_1, ...,
// R_n) and C = (C_1, ..., C_n) a partition of the indexed rows and columns
// respectively. A Block is defined as an ordered pair, B = (R_i, C_j).
type Block struct {
	// Row part of the block. Contains row's indexes.
	R *IntSet

	// Column part of the block. Contains column's indexes.
	C *IntSet

	// The size of a block is the amount of non-zero entries in sub-matrix of M
	// defined by the rows and columns of B.
	size int

	// A row block is simply a block conformed by a single row, e.g. (r, C_j). A
	// row block's size then is the amount of non-zero entries defined by the
	// columns of C_j in the row. The following map contains the sizes of all the
	// row blocks in Ri.
	rowBlocksSizes map[int]int
}

// NewBlockFromIntSet initializes a block given two IntSets.
func NewBlockFromIntSets(Ri, Cj *IntSet) *Block {
	return &Block{
		R:              Ri,
		C:              Cj,
		size:           0,
		rowBlocksSizes: make(map[int]int),
	}
}

// NewBlockFromIndexes initializes a block given rows and columns indexes.
func NewBlockFromIndexes(rowIndexes, colIndexes []int) *Block {
	Ri := set.NewIntSetFromValues(rowIndexes)
	Cj := set.NewIntSetFromValues(colIndexes)

	return NewBlockFromIntSets(Ri, Cj)
}

// GetRows returns the row part of B. If B = (R_i, C_j) then the function
// returns R_i.
func (B *Block) GetRows() *IntSet {
	return B.R
}

// GetColumns returns the column part of B. If B = (R_i, C_j) then the function
// returns C_j.
func (B *Block) GetColumns() *IntSet {
	return B.C
}

// Sets the block's size.
func (B *Block) SetSize(s int) {
	B.size = s
}

// Get the block's size.
func (B *Block) GetSize() int {
	return B.size
}

// SetRowBlockMap sets the size for the row blocks.
func (B *Block) SetRowBlockMap(sizes map[int]int) {
	B.rowBlocksSizes = sizes
}

// SetRowBlockSize sets the size of the row block (r, Cj) of B = (Ri, Cj).
func (B *Block) SetRowBlockSize(row, size int) {
	B.rowBlocksSizes[row] = size
}

// GetRowBlockSize returns the size of the block (r, Cj) where r is the given
// row index.
func (B *Block) GetRowBlockSize(row int) int {
	return B.rowBlocksSizes[row]
}

// GetRowBlockMap returns the map which contains the sizes of all the row
// blocks of B.
func (B *Block) GetRowBlockMap() map[int]int {
	return B.rowBlocksSizes
}

// IsConstant returns if a block is constant. A block B = (Ri, Cj) is constant
// if and only if size(B) = 0 or size(B) = |Ri||Cj|.
func (B *Block) IsConstant() bool {
	numRows := B.GetRows().Cardinality()
	numCols := B.GetColumns().Cardinality()
	sizeB := B.GetSize()

	return (sizeB == 0) || (sizeB == (numRows * numCols))
}
