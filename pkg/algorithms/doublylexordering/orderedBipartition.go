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

	// Previous partition in the order.
	previous *OrderedBipartition

	// Next partition in the order.
	next *OrderedBipartition

	// Indicates if its at the start of the ordered partition.
	start bool
}

// NewOrderedBipartitionFromIntSet initializes an ordered bipartition with the
// given set S.
func NewOrderedBipartitionFromIntSet(S *IntSet) *OrderedBipartition {
	return &OrderedBipartition{
		set:            S,
		leftPartition:  nil,
		rightPartition: nil,
		hasPartition:   false,
		previous:       nil,
		next:           nil,
		start:          false,
	}
}

// NewOrderedBipartitionFromIntSlice initializes an ordered bipartition with the
// given values.
func NewOrderedBipartitionFromIntSlice(s []int) *OrderedBipartition {
	return &OrderedBipartition{
		set:            set.NewIntSetFromValues(s),
		leftPartition:  nil,
		rightPartition: nil,
		hasPartition:   false,
		previous:       nil,
		next:           nil,
		start:          false,
	}
}

// GetSet returns the set contained in the ordered bipartition.
func (O *OrderedBipartition) GetSet() *IntSet {
	return O.set
}

// SetPartition defines a bipartition for the ordered bipartition and respects
// the previous order. Let P be the original partition and Pl and Pr the left
// and right partition respectively. Then P's previous part is now Pl's previous
// part, P's next part is now Pr's next part, Pl's next part is Pr and Pr's
// previous part is Pl.
func (O *OrderedBipartition) SetPartition(left, right *OrderedBipartition) {
	// Define bipartition.
	O.leftPartition = left
	O.rightPartition = right
	O.hasPartition = true

	// Preserve the existing order.
	previousPart := O.GetPrevious()
	nextPart := O.GetNext()

	previousPart.SetNext(left)
	left.SetPrevious(previousPart)
	left.SetNext(right)
	right.SetPrevious(left)
	right.SetNext(nextPart)
	nextPart.SetPrevious(right)

	// If the original part was at the start, then the left partition is now at
	// the start.
	if O.IsStart() {
		left.SetStart()
	}
}

// HasPartition returns if the ordered bipartition has a defined bipartition.
func (O *OrderedBipartition) HasPartition() bool {
	return O.hasPartition
}

// GetPrevious returns the part that goes before in the order, in other words,
// the previous part.
func (O *OrderedBipartition) GetPrevious() *OrderedBipartition {
	return O.previous
}

// GetNext returns the part that goes next in the order.
func (O *OrderedBipartition) GetNext() *OrderedBipartition {
	return O.next
}

// SetPrevious defines the part that goes before in the order, in other words,
// it defines the previous part.
func (O *OrderedBipartition) SetPrevious(p *OrderedBipartition) {
	O.previous = p
}

// SetNext defines the part that goes next in the order.
func (O *OrderedBipartition) SetNext(n *OrderedBipartition) {
	O.next = n
}

// SetStart indicates that this part is at the start of the order.
func (O *OrderedBipartition) SetStart() {
	O.start = true
}

// IsStart returns if the part is at the start of the order.
func (O *OrderedBipartition) IsStart() bool {
	return O.start
}

// GetOrderedPartitionAsSlice returns the ordered partition as an int slice.
func GetOrderedPartitionAsSlice(start *OrderedBipartition) []int {
	// Define a slice to save the order.
	var order []int

	// Traverse all the parts.
	current := start

	for current != nil {
		// Append all the values in the set of the current part to the slice.
		for _, val := range current.GetSet().GetValues() {
			order = append(order, val)
		}

		// Move to the next part.
		current = current.GetNext()
	}

	return order
}
