package main

// реализуйте быстрое множество
type IntSet struct {
	Set map[int]bool
}

func MakeIntSet() IntSet {

	return IntSet{Set: make(map[int]bool)}
}

func (s IntSet) Contains(elem int) bool {
	return s.Set[elem]
}

func (s IntSet) Add(elem int) bool {
	if !s.Contains(elem) {
		s.Set[elem] = true
		return true
	}
	return false
}
