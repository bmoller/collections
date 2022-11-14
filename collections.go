package collections

import (
	"errors"
	"fmt"
)

// Collection

type Collection[T comparable] interface {
	Contains(T) bool
	ContainsAll([]T) bool
	Empty() bool
	Equals(Collection[T]) bool
	Size() int
	Slice() []T
}

// List

type List[T comparable] interface {
	Collection[T]

	Get(int) (T, error)
	IndexOf(T) int
	Insert(int, T) (T, error)
	LastIndexOf(T) int
	Remove(int) T
	ReplaceAll(func(T) T)
	Sort(func(T, T))
	SubList(int, int) List[T]
}

var ErrEmptyList = errors.New("list is empty")

type ErrInvalidIndex struct {
	RequestedIndex int
	Size           int
}

func (e ErrInvalidIndex) Error() string {
	return fmt.Sprintf("index %d is invalid for list of length %d", e.RequestedIndex, e.Size)
}

// Set

type Set[T comparable] interface {
	Collection[T]

	Add(T) bool
	Has(T) bool
	Pop() (T, error)
	Remove(T) bool
}

var ErrEmptySet = errors.New("stack is empty")

// Stack

type Stack[T comparable] interface {
	Collection[T]

	Peek() (T, error)
	Pop() (T, error)
	Push(T) T
	Search(T) int
}

// ErrEmptyStack is returned when either Peek() or Pop() is called on an empty Stack.
var ErrEmptyStack = errors.New("stack is empty")
