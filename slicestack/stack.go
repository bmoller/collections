package slicestack

import "github.com/bmoller/collections"

const (
	stackGrowthFactor int = 2
	stackInitialSize  int = 100
)

type sliceStack[T comparable] struct {
	data []T
	size int
}

/*
New creates and returns a new instance of the default Stack implementation.
*/
func New[T comparable]() collections.Stack[T] {
	return &sliceStack[T]{
		data: make([]T, stackInitialSize),
	}
}

/*
NewWithSize creates a new Stack, with support for specifying the size of the backing array.
In some situations it may be advantageous to allocate the entire size needed if the longest possible length is known.
*/
func NewWithSize[T comparable](size int) collections.Stack[T] {
	return &sliceStack[T]{
		data: make([]T, size),
	}
}

/*
Empty indicates whether the Stack currently holds any elements.
*/
func (s *sliceStack[T]) Empty() bool {
	return s.size == 0
}

/*
Peek returns the element at the top of the Stack without removing it.
An empty Stack will return ErrEmptyStack.
*/
func (s *sliceStack[T]) Peek() (item T, err error) {
	if s.size == 0 {
		err = collections.ErrEmptyStack
	} else {
		item = s.data[s.size-1]
	}

	return
}

/*
Pop returns the element at the top of the Stack and removes it from the Stack.
An empty Stack will return ErrEmptyStack.
*/
func (s *sliceStack[T]) Pop() (item T, err error) {
	if s.size == 0 {
		err = collections.ErrEmptyStack
	} else {
		item = s.data[s.size-1]
		s.size--
	}

	return
}

/*
Push adds an element to the top of the stack.
*/
func (s *sliceStack[T]) Push(item T) {
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

/*
Size returns the number of elements currently in the Stack.
*/
func (s *sliceStack[T]) Size() int {
	return s.size
}
