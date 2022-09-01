package generators

import (
	"testing"

	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

// TestIsClique runs over all the possible subsets of vertices with size greater
// or equal than 2 from a complete graph of order 6 and verifies if each one is
// a clique. Next, it verifies if an almost-complete graph is a clique.
func TestIsClique(t *testing.T) {
	a := [][]byte{
		{0, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
	}
	vertices := make([]int, len(a))
	for i := 0; i < len(vertices); i++ {
		vertices[i] = i
	}
	subsets := [][]int{}
	subsets = append(subsets, []int{})
	for _, v := range vertices {
		for _, set := range subsets {
			contains := false
			for _, i := range set {
				if i == v {
					contains = true
				}
			}
			if !contains {
				subsets = append(subsets, append(set, v))
			}
		}
	}
	subsets = append(subsets, vertices)
	g := graph.NewFromMatrix(a)
	for _, set := range subsets {
		if len(set) > 1 && !IsClique(g, set) {
			t.Errorf(
				"Expected %v, but got %v",
				true,
				IsClique(g, set),
			)
		}
	}
	b := [][]byte{
		{0, 1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 1},
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 0},
	}
	g = graph.NewFromMatrix(b)
	if IsClique(g, vertices) {
		t.Errorf(
			"Expected %v, but got %v",
			false,
			IsClique(g, vertices),
		)
	}
}

// TestIsStable calls IsStable with a graph where each pair of vertices is
// non-adjacent except for the ones that contain vertices 1 & 3. It runs over
// all possible subsets of vertices with size greater or equal than 2.
func TestIsStable(t *testing.T) {
	a := [][]byte{
		{0, 1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1, 1},
		{0, 1, 0, 1, 0, 0},
		{1, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 0},
	}
	vertices := make([]int, len(a))
	for i := 0; i < len(vertices); i++ {
		vertices[i] = i
	}
	subsets := [][]int{}
	subsets = append(subsets, []int{})
	for _, v := range vertices {
		for _, set := range subsets {
			contains := false
			for _, i := range set {
				if i == v {
					contains = true
				}
			}
			if !contains {
				subsets = append(subsets, append(set, v))
			}
		}
	}
	subsets = append(subsets, vertices)
	g := graph.NewFromMatrix(a)
	for _, set := range subsets {
		if len(set) < 2 {
			continue
		}
		stableSet := true
		for _, v := range set {
			if v == 1 || v == 3 {
				stableSet = false
				break
			}
		}
		if stableSet && !IsStable(g, set) {
			t.Errorf(
				"Expected %v, but got %v",
				true,
				IsStable(g, set),
			)
		} else if !stableSet && IsStable(g, set) {
			t.Errorf(
				"Expected %v, but got %v",
				false,
				IsStable(g, set),
			)
		}
	}
}
