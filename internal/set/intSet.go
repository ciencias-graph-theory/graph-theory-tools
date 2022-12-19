package set

// An IntSet is a set of integers.
type IntSet struct {
	// There is no built-in data structure to represent a set. To get around this we
	// will define a set as a boolean map where all the elements inside the map will
	// always point to an empty struct (empty struct take 0 bytes on memory).
	setMap map[int]struct{}
}

// NewIntSet initializes an empty integer set.
func NewIntSet() *IntSet {
	return &IntSet{
		setMap: make(map[int]struct{}),
	}
}

// NewIntSetFromValues returns a set that contains the given values.
func NewIntSetFromValues(values []int) *IntSet {
	// Define an empty "set".
	tempSet := make(map[int]struct{})

	// Add the values to the "set". As we are using a map, then there is no
	// repetition of elements.
	for _, v := range values {
		tempSet[v] = struct{}{}
	}

	// Return the "set" as an IntSet.
	return &IntSet{
		setMap: tempSet,
	}
}

// Add adds an integer to the set, returns whether the set contains the given
// integer.
func (s *IntSet) Add(item int) bool {
	if _, contains := s.setMap[item]; contains {
		return true
	} else {
		s.setMap[item] = struct{}{}
		return false
	}
}

// Contains returns whether the set contains the given integer.
func (s *IntSet) Contains(item int) bool {
	if _, contains := s.setMap[item]; contains {
		return true
	} else {
		return false
	}
}

// Remove deletes an integer from the set, returns whether the set contained
// the given integer.
func (s *IntSet) Remove(item int) bool {
	if _, contains := s.setMap[item]; contains {
		delete(s.setMap, item)
		return true
	} else {
		return false
	}
}

// Cardinality returns the amount of elements in the set.
func (s *IntSet) Cardinality() int {
	return len(s.setMap)
}

// IsEmpty tests whether the set has no items.
func (s *IntSet) IsEmpty() bool {
	return len(s.setMap) == 0
}

// GetValues returns the values in the set. The order of this values is not
// deterministic; executing this function two times may not yield the same
// result.
func (s *IntSet) GetValues() []int {
	vals := make([]int, s.Cardinality(), s.Cardinality())

	k := 0
	for v, _ := range s.setMap {
		vals[k] = v
		k++
	}

	return vals
}

// Returns whether or not two integer sets are equal.
func (s *IntSet) Equals(x *IntSet) bool {
	// If the sets differ in length, return false.
	if s.Cardinality() != x.Cardinality() {
		return false
	}

	// Check that all the values of s are in x.
	svals := s.GetValues()

	for _, val := range svals {
		// If x doesn't contain a value of s, return false.
		if !x.Contains(val) {
			return false
		}
	}

	// If all the values of s are in x, and they have the same length, they are
	// equal.
	return true
}
