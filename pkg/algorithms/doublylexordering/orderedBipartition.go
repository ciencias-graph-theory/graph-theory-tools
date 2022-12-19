package doublylexordering

import (
	"github.com/ciencias-graph-theory/graph-theory-tools/internal/set"
)

type IntSet = set.IntSet

// A partition of a set is a grouping of its elements into non-empty subsets. An
// ordered partition of a finite set S is an ordered sequence (S_1, S_2, ...,
// S_k) where {S_1, S_2, ..., S_k} is a partition of F. An ordered bipartition
// is an ordered partition where there are only two parts; the first part is
// called left partition and the second one is the right partition.
type OrderedBipartition struct {
	// The set S.
	set *IntSet

	// The bipartition of S.
	leftPartition  *OrderedBipartition
	rightPartition *OrderedBipartition

	// Tells if S has a partition defined.
	hasPartition bool
}

func NewOrderedBipartitionFromIntSet(S *IntSet) *OrderedBipartition {
	return &OrderedBipartition{
		set:            S,
		leftPartition:  nil,
		rightPartition: nil,
		hasPartition:   false,
	}
}

func (O *OrderedBipartition) GetSet() *IntSet {
	return O.set
}

func (O *OrderedBipartition) SetPartition(left, right *OrderedBipartition) {
	O.leftPartition = left
	O.rightPartition = right
	O.hasPartition = true
}

func (O *OrderedBipartition) HasPartition() bool {
	return O.hasPartition
}
