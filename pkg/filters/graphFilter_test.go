package filters

import (
	// "fmt"
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/formatters"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
	"testing"
)

// equalGraphSlices returns whether or not two graph slices are equal.
func equalGraphSlices(A, B []*StaticGraph) bool {
	for i, a := range A {
		b := B[i]

		am, _ := a.Matrix()
		bm, _ := b.Matrix()

		if !sliceutils.EqualByteMatrix(am, bm) {
			return false
		}
	}

	return true
}

// hasLoops returns whether or not the given graph has loops.
func hasLoops(graph *StaticGraph) bool {
	// Obtain adjacency matrix.
	matrix, _ := graph.Matrix()

	for i, _ := range matrix {
		if matrix[i][i] == 1 {
			return true
		}
	}

	return false
}

func isSimple(graph *StaticGraph) bool {
	// Obtain adjacency matrix.
	matrix, _ := graph.Matrix()

	for i, v := range matrix {
		for j, _ := range v {
			// If it has loops, return false.
			if (i == j) && matrix[i][j] > 0 {
				return false
			}

			// If it has multiple edges, return false.
			if matrix[i][j] > 1 {
				return false
			}
		}
	}

	// Otherwise, return true.
	return true
}

// IsK2s returns whether or not the given graph is set of disjoint complete
// graphs of two vertices. We can determine this by the following conditional: A
// graph is a set of disjoint K2s if and only if all vertices have degree 1.
func isK2s(graph *StaticGraph) bool {
	// Obtain the degree sequence of the graph.
	degrees := graph.DegreeSequence()

	for _, d := range degrees {
		if d != 1 {
			return false
		}
	}

	return true
}

// IsCycles returns whether or not the given graph is a cycle or a set of
// disjoint cycles. We can the determine this by the following conditional: A
// graph is a set of disjoint cycles if and only if all vertices have degree 2.
func isCycles(graph *StaticGraph) bool {
	// Obtain the degree sequence of the graph.
	degrees := graph.DegreeSequence()

	for _, d := range degrees {
		if d != 2 {
			return false
		}
	}

	return true
}

func isComplete(graph *StaticGraph) bool {
	// Obtain the graph's order.
	n := graph.Order()

	// Obtain the degree sequence of the graph.
	degrees := graph.DegreeSequence()

	for _, d := range degrees {
		if d != (n - 1) {
			return false
		}
	}

	return true
}

func TestFilterGraph6(t *testing.T) {
	// Graph6 format of C4 and its isomorphism.
	C4 := "Cr"
	isomorphismC4 := "C]"

	// Graph6 format of two disjoint C4's.
	twoC4s := "Gr?GOK"

	// Graph6 format of two K2's.
	twoK2s := "C`"

	// Graph6 format of a complete graph with 4 vertices (K4).
	K4 := "C~"

	// Graph6 format of a complete graph with 3 vertices (K3).
	K3 := "Bw"

	// Graph6 format of three disjoint C3's, which are K3's.
	threeC3s := "HwCW?CB"

	// Graph6 format of four disjoint K2's.
	fourK2s := "G`?G?C"

	// An array containing all the graph6 strings.
	total := []string{C4, isomorphismC4, twoC4s, twoK2s, K4, K3, threeC3s, fourK2s}

	// Convert all the graph6 strings into graphs.
	gC4 := formatters.FromGraph6(C4)
	gIsomorphismC4 := formatters.FromGraph6(isomorphismC4)
	gTwoC4s := formatters.FromGraph6(twoC4s)
	gTwoK2s := formatters.FromGraph6(twoK2s)
	gK4 := formatters.FromGraph6(K4)
	gK3 := formatters.FromGraph6(K3)
	gThreeC3s := formatters.FromGraph6(threeC3s)
	gFourK2s := formatters.FromGraph6(fourK2s)

	// Classifications.
	cycles := []*StaticGraph{gC4, gIsomorphismC4, gTwoC4s, gK3, gThreeC3s}
	k2s := []*StaticGraph{gTwoK2s, gFourK2s}
	completes := []*StaticGraph{gK4, gK3}

	// Obtainded graphs.
	obtainedCycles := FilterGraph6(total, isCycles)
	obtainedK2s := FilterGraph6(total, isK2s)
	obtainedCompletes := FilterGraph6(total, isComplete)

	// Comparisons.
	expected := [][]*StaticGraph{
		cycles,
		k2s,
		completes,
	}

	obtained := [][]*StaticGraph{
		obtainedCycles,
		obtainedK2s,
		obtainedCompletes,
	}

	for i, _ := range expected {
		if !equalGraphSlices(expected[i], obtained[i]) {
			t.Errorf("Filter error: Expected %v but got %v", expected[i], obtained[i])
		}
	}

}

func TestFilterLoop6(t *testing.T) {
	// Loop6 format of C4 and its isomorphism.
	C4 := ";CSW"
	isomorphismC4 := ";CEo "

	// Loop6 format of two disjoint C4's.
	twoC4s := ";GSW??WK"

	// Loop6 format of two K2's.
	twoK2s := ";COG"

	// Loop6 format of a complete graph with 4 vertices (K4).
	K4 := ";CUw"

	// Loop6 format of a complete graph with 3 vertices (K3).
	K3 := ";BU"

	// Loop6 format of three disjoint C3's, which are K3's.
	threeC3s := ";HU?Oo?A?o"

	// Loop6 format of four disjoint K2's.
	fourK2s := ";GOG?O?A"

	// Loop6 format of a graph with a loop.
	loopG := ";Cmw"

	// An array containing all the graph6 strings.
	total := []string{C4, isomorphismC4, twoC4s, twoK2s, K4, K3, threeC3s, fourK2s, loopG}

	// Convert all the graph6 strings into graphs.
	lC4 := formatters.FromLoop6(C4)
	lIsomorphismC4 := formatters.FromLoop6(isomorphismC4)
	lTwoC4s := formatters.FromLoop6(twoC4s)
	lTwoK2s := formatters.FromLoop6(twoK2s)
	lK4 := formatters.FromLoop6(K4)
	lK3 := formatters.FromLoop6(K3)
	lThreeC3s := formatters.FromLoop6(threeC3s)
	lFourK2s := formatters.FromLoop6(fourK2s)
	lLoopG := formatters.FromLoop6(loopG)

	// Classifications.
	cycles := []*StaticGraph{lC4, lIsomorphismC4, lTwoC4s, lK3, lThreeC3s}
	k2s := []*StaticGraph{lTwoK2s, lFourK2s}
	completes := []*StaticGraph{lK4, lK3}
	loops := []*StaticGraph{lLoopG}

	// Obtainded graphs.
	obtainedCycles := FilterLoop6(total, isCycles)
	obtainedK2s := FilterLoop6(total, isK2s)
	obtainedCompletes := FilterLoop6(total, isComplete)
	obtainedLoops := FilterLoop6(total, hasLoops)

	// Comparisons.
	expected := [][]*StaticGraph{
		cycles,
		k2s,
		completes,
		loops,
	}

	obtained := [][]*StaticGraph{
		obtainedCycles,
		obtainedK2s,
		obtainedCompletes,
		obtainedLoops,
	}

	for i, _ := range expected {
		if !equalGraphSlices(expected[i], obtained[i]) {
			t.Errorf("Filter error: Expected %v but got %v", expected[i], obtained[i])
		}
	}

}

func TestFilterSparse6(t *testing.T) {
	// Adj. matrices of different graphs.
	// C4 and its isomorphism.
	matrixC4 := [][]byte{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	matrixIsomorphismC4 := [][]byte{
		{0, 0, 1, 1},
		{0, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
	}

	matrixTwoC4 := [][]byte{
		{0, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
	}

	// Two disjoint K2's.
	matrixTwoK2 := [][]byte{
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
	}

	// Complete graph of 4 vertices (K4).
	matrixK4 := [][]byte{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// Complete graph of 3 vertices (K3).
	matrixK3 := [][]byte{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	// Three disjoint cycles of 3 vertices.
	matrixThreeC3 := [][]byte{
		{0, 1, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 0},
	}

	// Four disjoint K2's.
	matrixFourK2 := [][]byte{
		{0, 1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 0},
	}

	// An arbitrary graph with loops.
	matrixLoop := [][]byte{
		{1, 0, 1, 1},
		{0, 1, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	// An arbitrary graph with multiple edges and loops.
	matrixMultiEdgesLoops := [][]byte{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 2, 0},
	}

	// Convert matrices into graphs.
	C4, _ := graph.NewGraphFromMatrix(matrixC4)
	isomorphismC4, _ := graph.NewGraphFromMatrix(matrixIsomorphismC4)
	twoC4, _ := graph.NewGraphFromMatrix(matrixTwoC4)
	twoK2, _ := graph.NewGraphFromMatrix(matrixTwoK2)
	K4, _ := graph.NewGraphFromMatrix(matrixK4)
	K3, _ := graph.NewGraphFromMatrix(matrixK3)
	threeC3, _ := graph.NewGraphFromMatrix(matrixThreeC3)
	fourK2, _ := graph.NewGraphFromMatrix(matrixFourK2)
	graphLoop, _ := graph.NewGraphFromMatrix(matrixLoop)
	graphMultiEdgesLoops, _ := graph.NewGraphFromMatrix(matrixMultiEdgesLoops)

	// Obtain its sparse6 formats.
	sp6C4 := formatters.ToSparse6(C4)
	sp6IsomorphismC4 := formatters.ToSparse6(isomorphismC4)
	sp6TwoC4 := formatters.ToSparse6(twoC4)
	sp6TwoK2 := formatters.ToSparse6(twoK2)
	sp6K4 := formatters.ToSparse6(K4)
	sp6K3 := formatters.ToSparse6(K3)
	sp6ThreeC3 := formatters.ToSparse6(threeC3)
	sp6FourK2 := formatters.ToSparse6(fourK2)
	sp6Loop := formatters.ToSparse6(graphLoop)
	sp6MultiEdgesLoop := formatters.ToSparse6(graphMultiEdgesLoops)

	// An array containing all the graph6 strings.
	total := []string{
		sp6C4, sp6IsomorphismC4, sp6TwoC4,
		sp6TwoK2, sp6K4, sp6K3,
		sp6ThreeC3, sp6FourK2, sp6Loop,
		sp6MultiEdgesLoop,
	}

	// Classifications.
	cycles := []*StaticGraph{C4, isomorphismC4, twoC4, K3, threeC3}
	k2s := []*StaticGraph{twoK2, fourK2}
	completes := []*StaticGraph{K4, K3}
	loops := []*StaticGraph{graphLoop}
	simples := []*StaticGraph{
		C4, isomorphismC4, twoC4, twoK2,
		K4, K3, threeC3, fourK2,
	}

	// Obtainded graphs.
	obtainedCycles := FilterSparse6(total, isCycles)
	obtainedK2s := FilterSparse6(total, isK2s)
	obtainedCompletes := FilterSparse6(total, isComplete)
	obtainedLoops := FilterSparse6(total, hasLoops)
	obtainedSimples := FilterSparse6(total, isSimple)

	// Comparisons.
	expected := [][]*StaticGraph{
		cycles,
		k2s,
		completes,
		loops,
		simples,
	}

	obtained := [][]*StaticGraph{
		obtainedCycles,
		obtainedK2s,
		obtainedCompletes,
		obtainedLoops,
		obtainedSimples,
	}

	for i, _ := range expected {
		if !equalGraphSlices(expected[i], obtained[i]) {
			t.Errorf("Filter error: Expected %v but got %v", expected[i], obtained[i])
		}
	}

}
