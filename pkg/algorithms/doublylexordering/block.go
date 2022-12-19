package doublylexordering

// Let M be a square (0,1)-matrix with indexed rows and columns, R = (R_1, ...,
// R_n) and C = (C_1, ..., C_n) a partition of the indexed rows and columns
// respectively. A Block is defined as an ordered pair, B = (R_i, C_j).
type Block struct {
	// Row part of the block. Contains row's indexes.
	R *OrderedBipartition

	// Column part of the block. Contains column's indexes.
	C *OrderedBipartition

	// The size of a block is the amount of non-zero entries in sub-matrix of M
	// defined by the rows and columns of B.
	size int

	// A row block is simply a block conformed by a single row, e.g. (r, C_j). A
	// row block's size then is the amount of non-zero entries defined by the
	// columns of C_j in the row. The following map contains the sizes of all the
	// row blocks in Ri.
	rowBlocksSizes map[int]int
}

// NewBlock initializes a block.
func NewBlock(Ri, Cj *OrderedBipartition) *Block {
	return &Block{
		R:              Ri,
		C:              Cj,
		size:           0,
		rowBlocksSizes: nil,
	}
}

// GetRowPart returns the row part of B. If B = (R_i, C_j) then the function
// returns R_i.
func (B *Block) GetRowPart() *OrderedBipartition {
	return B.R
}

// GetColumnPart returns the column part of B. If B = (R_i, C_j) then the function
// returns C_j.
func (B *Block) GetColumnPart() *OrderedBipartition {
	return B.C
}

// Sets the block's size.
func (B *Block) SetSize(s int) {
	B.size = s
}

// SetRowBlocksSizes sets the size for the row blocks.
func (B *Block) SetRowBlocksSizes(sizes map[int]int) {
	B.rowBlocksSizes = sizes
}

// GetRowBlocksSizes returns the map which contains the sizes of all the row
// blocks of B.
func (B *Block) GetRowBlockSize(row int) int {
	return B.rowBlocksSizes[row]
}
