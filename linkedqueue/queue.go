package linkedqueue

import (
	"github.com/bmoller/collections"
)

type queueNode[T comparable] struct {
	next  *queueNode[T]
	value T
}

type linkedQueue[T comparable] struct {
	head *queueNode[T]
	size int
	tail *queueNode[T]
}

func New[T comparable]() collections.Queue[T] {
	return new(linkedQueue[T])
}

func (q *linkedQueue[T]) Empty() bool {
	return q.size == 0
}

func (q *linkedQueue[T]) Peek() (element T, err error) {
	if q.size == 0 {
		return element, collections.ErrEmptyQueue
	}

	return q.head.value, nil
}

func (q *linkedQueue[T]) Pop() (element T, err error) {
	if q.size == 0 {
		return element, collections.ErrEmptyQueue
	}

	element = q.head.value
	q.head = q.head.next
	q.size--

	return element, nil
}

func (q *linkedQueue[T]) Push(item T) {
	element := &queueNode[T]{
		value: item,
	}

	switch q.size {
	case 0:
		q.head = element
		q.tail = element
	default:
		q.tail.next = element
		q.tail = element
	}
	q.size++
}

func (q *linkedQueue[T]) Size() int {
	return q.size
}
