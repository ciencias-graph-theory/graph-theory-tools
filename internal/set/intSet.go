// Provides simple integer set implementation.
package set

// An IntSet is a set of integers, represented by a map with integer keys
// and empty struct values (empty structs take 0 bytes on memory).
type IntSet struct {
	setMap map[int]struct{}
}

// NewIntSet initializes an empty integer set.
func NewIntSet() *IntSet {
	return &IntSet{
		setMap: make(map[int]struct{}),
	}
}

// Add adds an integer to the set, returns whether the set contains
// the given integer.
func (s *IntSet) Add(item int) bool {
	if _, contains := s.setMap[item]; contains {
		return true
	} else {
		s.setMap[item] = struct{}{}
		return false
	}
}

// Contains tests whether the set contains the given integer.
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

// IsEmpty tests whether the set has no items.
func (s *IntSet) IsEmpty() bool {
	return len(s.setMap) == 0
}

// Items returns the underlying integer -> struct map.
func (s *IntSet) Items() map[int]struct{} {
	return s.setMap
}
