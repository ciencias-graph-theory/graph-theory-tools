package filters

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/formatters"
	"github.com/ciencias-graph-theory/graph-theory-tools/pkg/graph"
)

type Graph = graph.Graph
type StaticGraph = graph.StaticGraph
type StaticDigraph = graph.StaticDigraph

// Given a slice of formatted graph strings and a boolean function (condition),
// return a slice of the graphs that satisfy the condition given by the boolean
// function.
func filterFormats(strs []string, condition func(*StaticGraph) bool, formatting func(string) *StaticGraph) []*StaticGraph {
	// Slice to keep the graphs that satisfy the condition.
	var sat []*StaticGraph

	for _, s := range strs {
		// Convert the string into a graph using the formatting function.
		g := formatting(s)

		// If the graph satisfies the condition, add it to the slice.
		if condition(g) {
			sat = append(sat, g)
		}
	}

	return sat
}

// Given graph6 strings, return as a slice the ones that satisfy the given
// condition.
func FilterGraph6(strs []string, condition func(*StaticGraph) bool) []*StaticGraph {
	return filterFormats(strs, condition, formatters.FromGraph6)
}

// Given loop6 strings, return as a slice the ones that satisfy the given
// condition.
func FilterLoop6(strs []string, condition func(*StaticGraph) bool) []*StaticGraph {
	return filterFormats(strs, condition, formatters.FromLoop6)
}

// Given sparse6 strings, return as a slice the ones that satisfy the given
// condition.
func FilterSparse6(strs []string, condition func(*StaticGraph) bool) []*StaticGraph {
	return filterFormats(strs, condition, formatters.FromSparse6)
}

// Given digraph6 strings, return as a slice the ones that satisfy the given
// condition.
func FilterDigraph6(strs []string, condition func(*StaticDigraph) bool) []*StaticDigraph {
	// Slice to keep the digraphs that satisfy the condition.
	var sat []*StaticDigraph

	for _, s := range strs {
		// Convert the string into a digraph using the formatting function.
		g := formatters.FromDigraph6(s)

		// If the digraph satisfies the condition, add it to the slice.
		if condition(g) {
			sat = append(sat, g)
		}
	}

	return sat
}
