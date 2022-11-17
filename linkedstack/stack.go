package linkedstack

import "github.com/bmoller/collections"

type node[T comparable] struct {
	previous *node[T]
	value    T
}

type stack[T comparable] struct {
	size int
	top  *node[T]
}

func New[T comparable]() collections.Stack[T] {
	return new(stack[T])
}

func (s *stack[T]) Empty() bool {
	return s.size == 0
}

func (s *stack[T]) Peek() (element T, err error) {
	if s.size == 0 {
		return element, collections.ErrEmptyStack
	}

	return s.top.value, nil
}

func (s *stack[T]) Pop() (element T, err error) {
	if s.size == 0 {
		return element, collections.ErrEmptyStack
	}

	element = s.top.value
	s.top = s.top.previous
	s.size--

	return element, nil
}

func (s *stack[T]) Push(item T) {
	top := &node[T]{
		previous: s.top,
		value:    item,
	}
	s.top = top
	s.size++
}

func (s *stack[T]) Size() int {
	return s.size
}
