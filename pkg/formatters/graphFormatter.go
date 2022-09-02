package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
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
	vGroups := sliceutils.DivideByteSlice(vExtended, 6)

	// Convert the bits into its binary number.
	vInts := sliceutils.ByteMatrixToIntSlice(vGroups)

	// Sum 63 to each int so that they can be printable.
	ASCII := sliceutils.IntSliceSumToEach(vInts, 63)

	return ASCII
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
