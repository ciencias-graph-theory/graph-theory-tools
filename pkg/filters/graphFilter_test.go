package filters

import (
	// "fmt"
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/formatters"
	// "github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
	"testing"
)

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

	for i, C := range cycles {
		G := obtainedCycles[i]

		expMatrix, _ := C.Matrix()
		obtMatrix, _ := G.Matrix()

		if !sliceutils.EqualByteMatrix(expMatrix, obtMatrix) {
			t.Errorf("Filter error: Expected %v but got %v", expMatrix, obtMatrix)
		}
	}

	for i, K2 := range k2s {
		G := obtainedK2s[i]

		expMatrix, _ := K2.Matrix()
		obtMatrix, _ := G.Matrix()

		if !sliceutils.EqualByteMatrix(expMatrix, obtMatrix) {
			t.Errorf("Filter error: Expected %v but got %v", expMatrix, obtMatrix)
		}
	}

	for i, Kn := range completes {
		G := obtainedCompletes[i]

		expMatrix, _ := Kn.Matrix()
		obtMatrix, _ := G.Matrix()

		if !sliceutils.EqualByteMatrix(expMatrix, obtMatrix) {
			t.Errorf("Filter error: Expected %v but got %v", expMatrix, obtMatrix)
		}
	}

}
