package formatters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// Parses a byte slice to format6.
// This function is known in https://users.cecs.anu.edu.au/~bdm/data/formats.txt
// as R(x).
func parseByteSliceFormat6(v []byte) []int {

	// Extend the previous vector so its length is a multiple of 6.
	vExtended := sliceutils.ExtendByteSlice(v, 6)

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
	if n <= 62 {
		return []int{(n + 63)}
	} else if n <= 258047 {
		nbin := sliceutils.IntToByteSlice(n)
		nASCII := parseByteSliceFormat6(nbin)
		return append([]int{63}, nASCII...)
	} else {
		nbin := sliceutils.IntToByteSlice(n)
		nASCII := parseByteSliceFormat6(nbin)
		return append([]int{63, 63}, nASCII...)
	}
}

// Returns the graph6 format string of a given static graph.
func ToGraph6(graph *StaticGraph) string {
	// Obtain the adjacency matrix.
	matrix, _ := graph.Matrix()

	// Obtain the vector of bits in the upper triangle of the adj. matrix.
	upperTriangle := sliceutils.ByteMatrixUpperTriangle(matrix, false)

	// Parse the edges of the graph to the accepted format6 standart.
	edgesASCII := parseByteSliceFormat6(upperTriangle)

	// Parse the order of the graph to the accepted format6 standart.
	orderASCII := parseOrderFormat6(graph.Order())

	// Concat both edges and orded ASCII values.
	graphASCII := append(orderASCII, edgesASCII...)

	// Retuns the ASCII representation of the graph.
	return sliceutils.IntSliceToASCII(graphASCII)
}
