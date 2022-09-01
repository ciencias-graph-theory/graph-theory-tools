package generators

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/sliceutils"
)

// IsClique receives a graph and a collection (subset) of vertices of the graph,
// and verifies whether every two of these vertices are adjacent.
func IsClique(g Graph, vertices []int) bool {
	for _, v := range vertices {
		s := g.Neighbours(v)
		for _, n := range vertices {
			if n != v && !s.Contains(n) {
				return false
			}
		}
	}
	return true
}

// IsStable receives a graph and a collection (subset) of vertices of the graph,
// and verifies whether every two of these vertices are non-adjacent.
func IsStable(g Graph, vertices []int) bool {
	for _, v := range vertices {
		s := g.Neighbours(v)
		for _, n := range vertices {
			if n != v && s.Contains(n) {
				return false
			}
		}
	}
	return true
}

// AreFullyAdjacent receives a graph and a two collections (subsets) of vertices
// of the graph, and verifies whether they are disjoint, and whether every
// vertex in one of the collections is adjacent to every vertex in the other one.
func AreFullyAdjacent(g Graph, x, y []int) bool {
	if !sliceutils.DisjointIntSlice(x, y) {
		return false
	}
	for _, v := range x {
		s := g.Neighbours(v)
		for _, w := range y {
			if !s.Contains(w) {
				return false
			}
		}
	}
	return true
}

// AreFullyNonAdjacent receives a graph and a two collections (subsets) of
// vertices of the graph, and verifies whether they are disjoint, and whether
// every  vertex in one of the collections is non adjacent to every vertex in
// the other one.
func AreFullyNonAdjacent(g Graph, x, y []int) bool {
	if !sliceutils.DisjointIntSlice(x, y) {
		return false
	}
	for _, v := range x {
		s := g.Neighbours(v)
		for _, w := range y {
			if s.Contains(w) {
				return false
			}
		}
	}
	return true
}
