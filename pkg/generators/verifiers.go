package generators

// IsClique receives a graph and a collection (subset) of vertices of the graph,
// and verifies whether every two of these vertices are adjacent.
func IsClique(g *StaticGraph, vertices []int) bool {
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
func IsStable(g *StaticGraph, vertices []int) bool {
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
