package doublylexordering

// Type aliases to improve readability.
type matrix = [][]byte

// Let M be a square (0,1)-matrix with indexed rows and columns, R = (R_1, ...,
// R_n) and C = (C_1, ..., C_n) a partition of the indexed rows and columns
// respectively. A Block is defined as an ordered pair, B = (R_i, C_j).
type Block struct {
	// Set of row's indexes.
	Ri *OrderedBipartition

	// Set of column's indexes.
	Cj *OrderedBipartition

	// The size of a block is the amount of non-zero entries in sub-matrix of M
	// defined by the rows and columns of B.
	size int

	// A row block is simply a block conformed by a single row, e.g. (r, C_j). A
	// row block's size then is the amount of non-zero entries defined by the
	// columns of C_j in the row. The following map contains the sizes of all the
	// row blocks in Ri.
	rowBlocksSizes map[int]int
}
