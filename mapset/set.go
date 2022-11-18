/*
Package mapset is a map-backed implementation of [collections.Set].

All index and capacity operations are handled by the backing map, so performance should match the performance of a map of the same size.
No order of elements is guaranteed, even between successive calls to Pop.
*/

package mapset

import "github.com/bmoller/collections"

type set[T comparable] struct {
	data map[T]bool
}

func New[T comparable]() collections.Set[T] {
	return &set[T]{
		data: make(map[T]bool),
	}
}

func (s *set[T]) Add(item T) {
	s.data[item] = true
}

func (s *set[T]) Contains(item T) bool {
	return s.data[item]
}

func (s *set[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *set[T]) Iterator() collections.Iterator[T] {
	var (
		err  error
		next T
		i    int
		set  []T
		size int = len(s.data)
	)
	set = make([]T, size)
	for element := range s.data {
		set[i] = element
		i++
	}
	i = 0

	return func() (T, error) {
		if i == size {
			err = collections.ErrNoMoreItems
		} else {
			next = set[i]
			i++
		}
		return next, err
	}
}

func (s *set[T]) Pop() (element T, err error) {
	if len(s.data) == 0 {
		err = collections.ErrEmptySet
	} else {
		for key := range s.data {
			element = key
			delete(s.data, key)
			break
		}
	}

	return
}

func (s *set[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *set[T]) Size() int {
	return len(s.data)
}

/*
Union returns the result of a set union between a and b, as a new Set.
A union includes all elements from both parent sets.
*/
func Union[T comparable](a, b collections.Set[T]) collections.Set[T] {
	result := make(map[T]bool)
	itr := a.Iterator()
	element, err := itr()
	for err == nil {
		result[element] = true
		element, err = itr()
	}
	itr = b.Iterator()
	element, err = itr()
	for err == nil {
		result[element] = true
		element, err = itr()
	}

	return &set[T]{
		data: result,
	}
}

/*
Intersection returns the set intersection of a and b, as a new Set.
An intersection includes only those items which are common to both parent sets.
*/
func Intersection[T comparable](a, b collections.Set[T]) collections.Set[T] {
	result := make(map[T]bool)
	itr := a.Iterator()
	element, err := itr()
	for err == nil {
		if b.Contains(element) {
			result[element] = true
		}
		element, err = itr()
	}

	return &set[T]{
		data: result,
	}
}

/*
Difference determines the set difference between a and b, and returns it as a new Set.
Difference includes only those elements which are unique to either parent Set.
*/
func Difference[T comparable](a, b collections.Set[T]) collections.Set[T] {
	result := make(map[T]bool)
	itr := a.Iterator()
	element, err := itr()
	for err == nil {
		if !b.Contains(element) {
			result[element] = true
		}
		element, err = itr()
	}

	return &set[T]{
		data: result,
	}
}

/*
IsSubset checks if Set a is a subset of Set b.
The Set a is a subset if all of its elements are also in Set b.
*/
func IsSubset[T comparable](a, b collections.Set[T]) bool {
	itr := a.Iterator()
	element, err := itr()
	for err == nil {
		if !b.Contains(element) {
			return false
		}
		element, err = itr()
	}

	return true
}
