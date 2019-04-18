package set

import (
	"sort"
	"strconv"
)

// IntSet is a way of handling sets in Golang, with ints
type IntSet struct {
	data map[int]bool
}

// New makes an IntSet
func New(initValues ...int) *IntSet {
	newSet := &IntSet{data: map[int]bool{}}
	for _, val := range initValues {
		newSet.Add(val)
	}
	return newSet
}

// Add adds an int or several to an IntSet
func (iset *IntSet) Add(is ...int) {
	for _, i := range is {
		iset.data[i] = true
	}
}

// Contains returns a bool to say if a set has certain data
func (iset *IntSet) Contains(i int) bool {
	_, ok := iset.data[i]
	return ok
}

// Remove removes a value from a set
func (iset *IntSet) Remove(i int) {
	delete(iset.data, i)
}

// Len returns length of a set
func (iset *IntSet) Len() (total int) {
	for range iset.data {
		total++
	}
	return
}

// Union takes the union of 2 sets
func (iset IntSet) Union(iset2 *IntSet) (resultSet *IntSet) {
	resultSet = New()
	for _, datamap := range []map[int]bool{iset.data, iset2.data} {
		for k := range datamap {
			resultSet.Add(k)
		}
	}
	return
}

// ToSlice returns a slice of all ints in an IntSet
func (iset IntSet) ToSlice() []int {
	var result []int
	for k := range iset.data {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

// Intersect takes the union of 2 sets
func (iset IntSet) Intersect(iset2 *IntSet) (resultSet *IntSet) {
	resultSet = New()
	for k := range iset.data {
		if iset2.Contains(k) {
			resultSet.Add(k)
		}
	}
	return
}

func mapEquals(m1 map[int]bool, m2 map[int]bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
		if v != m2[k] {
			return false
		}
	}
	return true
}

// Equals compares 2 sets to see if they have same contents
func (iset *IntSet) Equals(iset2 *IntSet) bool {
	return mapEquals(iset.data, iset2.data)
}

// Subtract takes one set away from another, returning a new set
func (iset *IntSet) Subtract(iset2 *IntSet) *IntSet {
	result := New()
	for k := range iset.data {
		if !iset2.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

// String to look nice e.g. Set[4, 5, 8]
func (iset *IntSet) String() string {
	stringPrefix := "Set["
	contentString := ""
	stringSuffix := "]"
	ints := iset.ToSlice()
	sort.Ints(ints)
	for i, val := range ints {
		valString := strconv.Itoa(val)
		contentString = contentString + valString
		if i != len(ints)-1 { // only add comma for non-last iteration
			contentString += ", "
		}
	}
	return stringPrefix + contentString + stringSuffix
}
