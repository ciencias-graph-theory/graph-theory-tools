// Provides simple integer set implementation.
package set

type IntSet struct {
	setMap map[int]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{
		setMap: make(map[int]struct{}),
	}
}

func (s *IntSet) Add(item int) bool {
	if _, contains := s.setMap[item]; contains {
		return true
	} else {
		s.setMap[item] = struct{}{}
		return false
	}
}

func (s *IntSet) Contains(item int) bool {
	if _, contains := s.setMap[item]; contains {
		return true
	} else {
		return false
	}
}

func (s *IntSet) Remove(item int) bool {
	if _, contains := s.setMap[item]; contains {
		delete(s.setMap, item)
		return true
	} else {
		return false
	}
}

func (s *IntSet) Items() map[int]struct{} {
	return s.setMap
}
