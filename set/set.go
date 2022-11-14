package collections

import "github.com/bmoller/collections"

type mapSet[T comparable] struct {
	data map[T]bool
}

func New[T comparable]() collections.Set[T] {
	return &mapSet[T]{
		data: make(map[T]bool),
	}
}

func (s *mapSet[T]) Add(item T) (success bool) {
	success = !s.data[item]
	s.data[item] = true

	return
}

func (s *mapSet[T]) Has(item T) bool {
	return s.data[item]
}

func (s *mapSet[T]) Items() (items []T) {
	items = make([]T, len(s.data))

	i := 0
	for item := range s.data {
		items[i] = item
		i++
	}

	return
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

func (s *mapSet[T]) Remove(item T) (success bool) {
	success = s.data[item]
	delete(s.data, item)

	return
}

func (s *mapSet[T]) Size() int {
	return len(s.data)
}
