package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
	"math"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// Parses a byte slice to format6. If leftPadding is true then the byte slice is
// extended by appending zeros to the left, otherwise they're appended to the
// right. This function is known in as R(x).
func parseByteSliceFormat6(v []byte, leftPadding bool) []int {

	// Extend the previous vector so its length is a multiple of 6.
	vExtended := sliceutils.ExtendByteSlice(v, 6, leftPadding)

	// Divide the previous vector in groups of 6 bits.
	vGroups, _ := sliceutils.DivideByteSlice(vExtended, 6)

	// Convert the bits into its binary number.
	vInts := sliceutils.ByteMatrixToIntSlice(vGroups)

	// Sum 63 to each int so that they can be printable.
	ASCII := sliceutils.IntSliceSumToEach(vInts, 63)

	return ASCII
}

// The inverse process of obtaining the format 6.
// Returns the original bits that produced the given slice of ints.
func inverseFormat6(v []int) []byte {
	// Decrease each value by 63 to obtain the original values.
	ogVals := sliceutils.IntSliceSumToEach(v, -63)
	numVals := len(v)

	// Obtain the binary representation of each of the original values.
	ogBits := make([][]byte, numVals)

	for i := 0; i < numVals; i++ {
		// Obtain the binary representation.
		binary := sliceutils.IntToByteSlice(ogVals[i])

		// Extend the binary to 6 bits.
		// A left padding is used to not modify the original value.
		extBinary := sliceutils.ExtendByteSlice(binary, 6, true)

		ogBits[i] = extBinary
	}

	// Obtain all the bits as a single slice.
	return sliceutils.ByteMatrixToSlice(ogBits)
}

// Parses the order of a graph to format 6.
// This function is known in https://users.cecs.anu.edu.au/~bdm/data/formats.txt
// as N(x).
func parseOrderFormat6(n int) []int {
	var missingBytes []int

	// If n <= 62, define N(n) as a single byte n+63.
	if n <= 62 {
		return []int{(n + 63)}
	}

	// N(n) is defined as followed if n >= 63.
	// Obtain the binary representation of n as byte slice.
	nbin := sliceutils.IntToByteSlice(n)

	// Parse the order of the graph to the format6.
	nASCII := parseByteSliceFormat6(nbin, true)

	// If 63 <= n <= 258047, define N(n) to be four bytes.
	if n <= 258047 {
		missingBytes = []int{126}

		// In case there are still missing some bytes, append the byte 63.
		for i := len(nASCII); i < 3; i++ {
			missingBytes = append(missingBytes, 63)
		}

		return append(missingBytes, nASCII...)
	} else {
		// If 258048 <= n <= 68719476735, define N(n) to be the eight bytes.
		missingBytes = []int{126, 126}

		// In case there are still missing some bytes, append the byte 63.
		for i := len(nASCII); i < 6; i++ {
			missingBytes = append(missingBytes, 63)
		}

		return append(missingBytes, nASCII...)
	}

}

// Returns the order of a graph given the format6 values.
// This function could be considered the inverse of N(n).
func inverseParseOrderFormat6(vals []int) int {
	// Obtain the order by inverting the process to obtain the format.
	orderBinary := inverseFormat6(vals)

	// Obtain the order value.
	order := sliceutils.ByteSliceToInt(orderBinary)

	return order
}

// Given a slice of ints corresponding to the format6 of a graph, determine the
// bytes that correspond to the order and edges. Return the order of the graph
// and the rest of the int slice.
func determineOrderAndEdges(vals []int) (int, []int) {
	var order int
	var edgeValues []int

	// If first byte is different from 126, then the number of bytes associated to
	// the order is just one.
	if vals[0] != 126 {
		order = inverseParseOrderFormat6([]int{vals[0]})
		edgeValues = vals[1:]
	} else if vals[1] != 126 {
		// If the first byte is 126 and the second byte is different from 126, then
		// the first four bytes are associated with the order of the graph.
		order = inverseParseOrderFormat6(vals[1:4])
		edgeValues = vals[4:]
	} else {
		// Otherwise, the first eight bytes are associated with the order.
		order = inverseParseOrderFormat6(vals[2:8])
		edgeValues = vals[8:]
	}

	return order, edgeValues
}

// Returns the adj. matrix corresponding to the format6 values.
func inverseParseEdgesFormat6(order int, vals []int, diag, sym bool) [][]byte {
	// Create an empty adj. matrix.
	matrix := make([][]byte, order)
	for i := 0; i < order; i++ {
		matrix[i] = make([]byte, order)
	}

	// Obtain the edges bits by inverting the process of the format6.
	edgeBits := inverseFormat6(vals)

	// Travel the bits to build the adj. matrix.
	k := 0
	for j := 0; j < order; j++ {
		for i := 0; i < order; i++ {

			// If i = j and the diagonal is not considered, continue to the next
			// column.
			if (i == j) && !diag {
				break
			}

			// If the matrix is not symmetric then it corresponds to a digraph.
			// Otherwise, only travel the upper triangle of the matrix.
			if (i > j) && sym {
				break
			}

			matrix[j][i] = edgeBits[k]

			// If the matrix is symmetric then a_ij = a_ji.
			if sym {
				matrix[i][j] = edgeBits[k]
			}

			k++
		}
	}

	return matrix
}

// Returns the graph6 format string of a given static graph.
func ToGraph6(graph *StaticGraph) string {
	// Obtain the adjacency matrix.
	matrix, _ := graph.Matrix()

	// Obtain the vector of bits in the upper triangle of the adj. matrix.
	upperTriangle := sliceutils.ByteMatrixUpperTriangle(matrix, false)

	// Parse the edges of the graph to the accepted format6 standart.
	edgesASCII := parseByteSliceFormat6(upperTriangle, false)

	// Parse the order of the graph to the accepted format6 standart.
	orderASCII := parseOrderFormat6(graph.Order())

	// Concat both edges and orded ASCII values.
	graphASCII := append(orderASCII, edgesASCII...)

	// Retuns the ASCII representation of the graph.
	return sliceutils.IntSliceToASCII(graphASCII)
}

func FromGraph6(s string) *StaticGraph {
	// Obtain the ASCII values of the string.
	values := sliceutils.ASCIIToIntSlice(s)

	// Determine the order of the graph and the values corresponding to the edges.
	order, edgeVals := determineOrderAndEdges(values)

	// Obtain the adjacency matrix given the edge values.
	matrix := inverseParseEdgesFormat6(order, edgeVals, false, true)

	// Build and return a graph given a matrix.
	G, _ := graph.NewGraphFromMatrix(matrix)
	return G
}

func ToLoop6(graph *StaticGraph) string {
	// Obtain the adjacency matrix.
	matrix, _ := graph.Matrix()

	// Obtain the vector of bits in the upper triangle of the adj. matrix
	// considering the diagonal.
	upperTriangle := sliceutils.ByteMatrixUpperTriangle(matrix, true)

	// Parse the edges of the graph to the accepted format6 standart.
	edgesASCII := parseByteSliceFormat6(upperTriangle, false)

	// Parse the order of the graph to the accepted format6 standart.
	orderASCII := parseOrderFormat6(graph.Order())

	// Concat both edges and orded ASCII values.
	graphASCII := append(orderASCII, edgesASCII...)

	// Append the loop6 identifier and return the ASCII representation of the
	// graph.
	return ";" + sliceutils.IntSliceToASCII(graphASCII)
}

func FromLoop6(s string) *StaticGraph {
	// Obtain the ASCII values of the string.
	values := sliceutils.ASCIIToIntSlice(s)

	// Ignore the identifier.
	values = values[1:]

	// Determine the order of the graph and the values corresponding to the edges.
	order, edgeVals := determineOrderAndEdges(values)

	// Obtain the adjacency matrix given the edge values. Consider the diagonal.
	matrix := inverseParseEdgesFormat6(order, edgeVals, true, true)

	// Build and return a graph given a matrix.
	G, _ := graph.NewGraphFromMatrix(matrix)
	return G
}

func ToDigraph6(digraph *StaticDigraph) string {
	// Obtain the adjacency matrix.
	matrix, _ := digraph.Matrix()

	// Obtain the vector of bits corresponding to the adj. matrix.
	bits := sliceutils.ByteMatrixToSlice(matrix)

	// Parse the edges of the graph to the accepted format6 standart.
	edgesASCII := parseByteSliceFormat6(bits, false)

	// Parse the order of the graph to the accepted format6 standart.
	orderASCII := parseOrderFormat6(digraph.Order())

	// Concat both edges and orded ASCII values.
	graphASCII := append(orderASCII, edgesASCII...)

	// Append the loop6 identifier and return the ASCII representation of the
	// graph.
	return "&" + sliceutils.IntSliceToASCII(graphASCII)
}

func FromDigraph6(s string) *StaticDigraph {
	// Obtain the ASCII values of the string.
	values := sliceutils.ASCIIToIntSlice(s)

	// Ignore the identifier.
	values = values[1:]

	// Determine the order of the graph and the values corresponding to the edges.
	order, edgeVals := determineOrderAndEdges(values)

	// Obtain the adjacency matrix given the edge values. Consider the diagonal.
	matrix := inverseParseEdgesFormat6(order, edgeVals, true, false)

	// Build and return a graph given a matrix.
	D := graph.NewDigraphFromMatrix(matrix)
	return D
}

// Given a bits slice, build the corresponding block of each edge.
func obtainEdgeBlocks(order int, bits []byte) [][]int {
	// Let k be how many bits are needed to represent the order
	// in binary.
	m := float64(order - 1)
	k := int(math.Ceil(math.Log2(m)))

	// Extend the bits slice so its length is multiple of k + 1.
	// Note: This can also be achieved by removing bits, but by
	// safety we prefer to append them.
	exBits := sliceutils.ExtendByteSlice(bits, k+1, false)

	// Create an empty matrix to store the blocks.
	numBlocks := (len(exBits) / (k + 1))
	bitsBlocks := make([][]byte, numBlocks)
	for i := 0; i < numBlocks; i++ {
		bitsBlocks[i] = make([]byte, k+1)
	}

	// Build the blocks.
	for i := 0; i < numBlocks; i++ {
		bitsBlocks[i] = exBits[(i * (k + 1)) : (i+1)*(k+1)]
	}

	// Create an empty matrix to store the integer equivalent
	// of the blocks.
	blocks := make([][]int, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blocks[i] = make([]int, 2)
	}

	// Convert each block.
	for i := 0; i < numBlocks; i++ {
		// The first element of the pair can only be 1 or 0.
		if bitsBlocks[i][0] == 1 {
			blocks[i][0] = 1
		} else {
			blocks[i][0] = 0
		}

		// The second element of the pair is int representation
		// of the rest of the k bits.
		blocks[i][1] = sliceutils.ByteSliceToInt(bitsBlocks[i][1:])
	}

	return blocks
}

// Build an adjacency matrix of n x n given the blocks of edges
// represented by the Sparse6 format, where n is the order of
// the graph.
func buildFromBlocks(order int, blocks [][]int) [][]byte {
	// Build an empty matrix of size n times n, where n is
	// the order.
	matrix := make([][]byte, order)
	for i := 0; i < order; i++ {
		matrix[i] = make([]byte, order)
	}

	// The letter i represents the current vertex.
	i := 0

	// The letter j represents the vertex i is adjacent with.
	j := 0

	// The limit is the last vertex that can be represented.
	limit := order - 1

	for b := 0; b < len(blocks); b++ {
		if blocks[b][0] == 1 {
			i++

			// If limit has been surpassed, break.
			if i > limit {
				break
			}
		}

		if blocks[b][1] > i {
			i = blocks[b][1]
		} else {
			j = blocks[b][1]

			// The matrix is symmetric.
			matrix[i][j]++
			if i != j {
				// Don't update twice if it is the diagonal.
				matrix[j][i]++
			}
		}
	}

	return matrix
}

// Returns the graph corresponding to the sparse6 string given.
func FromSparse6(s string) *StaticGraph {
	// Obtain the ASCII values of the string.
	values := sliceutils.ASCIIToIntSlice(s)

	// Ignore the identifier.
	values = values[1:]

	// Determine the order of the graph and the values corresponding to the edges.
	order, edgeBytes := determineOrderAndEdges(values)

	// Obtain the blocks corresponding to the edges.
	edgeBits := inverseFormat6(edgeBytes)
	blocks := obtainEdgeBlocks(order, edgeBits)

	// Build the adj. matrix given the blocks.
	matrix := buildFromBlocks(order, blocks)

	// Return the graph with its adj. matrix.
	G, _ := graph.NewGraphFromMatrix(matrix)
	return G
}
