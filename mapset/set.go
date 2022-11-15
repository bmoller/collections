package mapset

import "github.com/bmoller/collections"

type mapSet[T comparable] struct {
	data map[T]bool
}

func New[T comparable]() collections.Set[T] {
	return &mapSet[T]{
		data: make(map[T]bool),
	}
}

func (s *mapSet[T]) Add(item T) {
	s.data[item] = true
}

func (s *mapSet[T]) Contains(item T) bool {
	return s.data[item]
}

func (s *mapSet[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *mapSet[T]) Iterator() collections.Iterator[T] {
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

func (s *mapSet[T]) Pop() (element T, err error) {
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

func (s *mapSet[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *mapSet[T]) Size() int {
	return len(s.data)
}

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

	return &mapSet[T]{
		data: result,
	}
}

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

	return &mapSet[T]{
		data: result,
	}
}

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

	return &mapSet[T]{
		data: result,
	}
}

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
