// ©2022 Brandon Moller

/*
The slicestack package provides a Stack implementation backed by a slice.

The Stack inserts and removes items into and from the slice and tracks the top via an internal pointer
*/
package slicestack

import "github.com/bmoller/collections"

const (
	stackGrowthFactor int = 2   // Length of the backing array is multiplied by this when a replacement is allocated
	stackInitialSize  int = 100 // The initial size of the backing array and slice
)

type stack[T comparable] struct {
	data []T
	size int
}

func New[T comparable]() collections.Stack[T] {
	return &stack[T]{
		data: make([]T, stackInitialSize),
	}
}

/*
NewWithSize creates a new Stack, with support for specifying the size of the backing array.
In some situations it may be advantageous to allocate the entire size needed if the longest possible length is known.
*/
func NewWithSize[T comparable](size int) collections.Stack[T] {
	return &stack[T]{
		data: make([]T, size),
	}
}

func (s *stack[T]) Empty() bool {
	return s.size == 0
}

func (s *stack[T]) Peek() (item T, err error) {
	if s.size == 0 {
		err = collections.ErrEmptyStack
	} else {
		item = s.data[s.size-1]
	}

	return
}

func (s *stack[T]) Pop() (item T, err error) {
	if s.size == 0 {
		err = collections.ErrEmptyStack
	} else {
		item = s.data[s.size-1]
		s.size--
	}

	return
}

func (s *stack[T]) Push(item T) {
	s.data[s.size] = item
	s.size++
	if s.size == cap(s.data) {
		newData := make([]T, len(s.data)*stackGrowthFactor)
		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
	}
}

func (s *stack[T]) Size() int {
	return s.size
}
