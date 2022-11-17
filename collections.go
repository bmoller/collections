package collections

import (
	"errors"
	"fmt"
)

// Collection

type Collection[T comparable] interface {
	Empty() bool
	Size() int
}

// Iterator

type Iterator[T comparable] func() (T, error)

var ErrNoMoreItems = errors.New("no more items to return")

// List

type List[T comparable] interface {
	Collection[T]

	Add(T)
	Clear()
	Get(int) (T, error)
	Insert(int, T) error
	Remove(int) (T, error)
	SubList(int, int) (List[T], error)
}

var ErrEmptyList = errors.New("list is empty")

type ErrIndexOutOfRange struct {
	Index int
	Size  int
}

func (e ErrIndexOutOfRange) Error() string {
	return fmt.Sprintf("index %d is invalid for list of length %d", e.Index, e.Size)
}

type ErrInvalidRange struct {
	End   int
	Start int
}

func (e ErrInvalidRange) Error() string {
	return fmt.Sprintf("invalid range with start %d and end %d; valid ranges follow 0 <= start <= end", e.Start, e.End)
}

// LinkedList

type ListNode[T comparable] interface {
	Next() ListNode[T]
	Previous() ListNode[T]
	Value() T
}

type LinkedList[T comparable] interface {
	List[T]

	GetNode(int) (ListNode[T], error)
	Head() ListNode[T]
	InsertAfter(ListNode[T], T) (ListNode[T], error)
	InsertBefore(ListNode[T], T) (ListNode[T], error)
	RemoveNode(ListNode[T]) error
	Tail() ListNode[T]
}

var ErrNodeIsNotElement = errors.New("node is not an element of this list")

var ErrWrongNodeType = errors.New("node is from an incompatible list implementation")

// Queue

type Queue[T comparable] interface {
	Collection[T]

	Peek() (T, error)
	Pop() (T, error)
	Push(T)
}

var ErrEmptyQueue = errors.New("queue is empty")

// Set

type Set[T comparable] interface {
	Collection[T]

	Add(T)
	Contains(T) bool
	Iterator() Iterator[T]
	Pop() (T, error)
	Remove(T)
}

// ErrEmptySet is returned when Peek or Pop are called on an empty Set.
var ErrEmptySet = errors.New("stack is empty")

// Stack

type Stack[T comparable] interface {
	Collection[T]

	Peek() (T, error)
	Pop() (T, error)
	Push(T)
}

// ErrEmptyStack is returned when Peek or Pop are called on an empty Stack.
var ErrEmptyStack = errors.New("stack is empty")
