package filters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/formatters"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// Given an array of graph6 strings and a boolean function, return a slice of
// the graphs that satisfy the condition given by the boolean function.
func FilterGraph6(strs []string, f func(*StaticGraph) bool) []*StaticGraph {
	var sat []*StaticGraph

	for _, s := range strs {
		// Convert the string into a graph.
		g := formatters.FromGraph6(s)

		if f(g) {
			sat = append(sat, g)
		}
	}

	return sat
}
