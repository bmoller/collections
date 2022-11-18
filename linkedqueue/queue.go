// ©2022 Brandon Moller

/*
Package linkedqueue provides an implementation of [collections.Queue] backed by individual node instances.
Each element added to the queue is stored in a node, with a pointer to the next node.
The queue maintains references to the next node to return and the tail for fast Pop and Push operations.
*/
package linkedqueue

import (
	"github.com/bmoller/collections"
)

type queueNode[T comparable] struct {
	next  *queueNode[T]
	value T
}

type queue[T comparable] struct {
	head *queueNode[T]
	size int
	tail *queueNode[T]
}

func New[T comparable]() collections.Queue[T] {
	return new(queue[T])
}

func (q *queue[T]) Empty() bool {
	return q.size == 0
}

func (q *queue[T]) Peek() (element T, err error) {
	if q.size == 0 {
		return element, collections.ErrEmptyQueue
	}

	return q.head.value, nil
}

func (q *queue[T]) Pop() (element T, err error) {
	if q.size == 0 {
		return element, collections.ErrEmptyQueue
	}

	element = q.head.value
	q.head = q.head.next
	q.size--

	return element, nil
}

func (q *queue[T]) Push(item T) {
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

func (q *queue[T]) Size() int {
	return q.size
}
