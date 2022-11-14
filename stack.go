package collections

import "errors"

const (
	stackGrowthFactor int = 2
	stackInitialSize  int = 100
)

type Stack[T comparable] interface {
	Empty() bool
	Peek() (T, error)
	Pop() (T, error)
	Push(T) T
	Search(T) int
}

// ErrEmptyStack is returned when either Peek() or Pop() is called on an empty Stack.
var ErrEmptyStack = errors.New("stack is empty")

type sliceStack[T comparable] struct {
	data []T
	size int
}

/*
NewStack creates and returns a new instance of the default Stack implementation.
*/
func NewStack[T comparable]() Stack[T] {
	return &sliceStack[T]{
		data: make([]T, stackInitialSize),
	}
}

/*
Empty indicates whether the Stack currently holds any elements.
*/
func (s *sliceStack[T]) Empty() bool {
	return s.size == 0
}

/*
Peek returns the element at the top of the Stack, but does not remove it.
The size of the Stack does not change, and a subsequent call to Peek() or Pop() will return the same element.

An empty Stack will return ErrEmptyStack.
*/
func (s *sliceStack[T]) Peek() (item T, err error) {
	if s.size == 0 {
		err = ErrEmptyStack
	} else {
		item = s.data[s.size-1]
	}

	return
}

/*
Pop returns the element at the top of the Stack and removes it from the Stack.
The size of the Stack decreases by 1, and the next call to Peek() or Pop() will return a different element.

An empty Stack will return ErrEmptyStack.
*/
func (s *sliceStack[T]) Pop() (item T, err error) {
	item, err = s.Peek()
	if err == nil {
		s.size--
	}

	return
}

/*
Push adds an element to the top of the stack.
If no subsequent elements are pushed onto the Stack, the next call to Peek() or Pop() will return this element.
The item is returned as verification that the Stack holds the correct element at its top.
*/
func (s *sliceStack[T]) Push(item T) T {
	s.data[s.size] = item
	s.size++
	if s.size == cap(s.data) {
		newData := make([]T, len(s.data)*stackGrowthFactor)
		copy(newData, s.data)
		s.data = newData
	}

	return s.data[s.size-1]
}

/*
Search scans the Stack from the top for item and returns its index.
The index is 0-based, with the top element having index 0.
If item is not found in the Stack, the value -1 is returned.
Seach can safely be called on an empty Stack and will not return an error.
*/
func (s *sliceStack[T]) Search(item T) (index int) {
	index = -1

	if s.size == 0 {
		return index
	} else {
		for i := 0; i < s.size; i++ {
			if s.data[s.size-1-i] == item {
				index = i
				break
			}
		}
	}

	return index
}
