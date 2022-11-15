package slicelist

import "github.com/bmoller/collections"

const (
	growthFactor int = 2
	initialSize  int = 100
)

type sliceList[T comparable] struct {
	data []T
	size int
}

/*
NewList creates and returns a new List backed by a Go slice.
*/
func New[T comparable]() collections.List[T] {
	return &sliceList[T]{
		data: make([]T, initialSize),
	}
}

func NewFromItems[T comparable](items []T) collections.List[T] {
	return &sliceList[T]{
		data: items,
		size: len(items),
	}
}

func NewWithSize[T comparable](size int) collections.List[T] {
	return &sliceList[T]{
		data: make([]T, size),
	}
}

func (l *sliceList[T]) Add(item T) {
	if l.size+1 == cap(l.data) {
		newData := make([]T, (l.size+1)*growthFactor)
		for i := 0; i < l.size; i++ {
			newData[i] = l.data[i]
		}
		l.data = newData
	}
	l.data[l.size] = item
	l.size++
}

func (l *sliceList[T]) Clear() {
	l.size = 0
}

func (l *sliceList[T]) Empty() bool {
	return l.size == 0
}

func (l *sliceList[T]) Get(index int) (item T, err error) {
	switch {
	case l.size == 0:
		err = collections.ErrEmptyList
	case index >= l.size:
		err = collections.ErrInvalidIndex{
			RequestedIndex: index,
			Size:           l.size,
		}
	default:
		item = l.data[index]
	}

	return item, err
}

func (l *sliceList[T]) Insert(index int, item T) error {
	if index < 0 || index > l.size {
		return collections.ErrInvalidIndex{
			RequestedIndex: index,
			Size:           l.size,
		}
	}

	if newSize := l.size + 1; newSize == cap(l.data) {
		newData := make([]T, newSize*growthFactor)
		for i := 0; i < index; i++ {
			newData[i] = l.data[i]
		}
		newData[index] = item
		for i := index; i < l.size; i++ {
			newData[i+1] = l.data[i]
		}
	} else {
		for i := l.size; i > index; i-- {
			l.data[i] = l.data[i-1]
		}
		l.data[index] = item
	}
	l.size++

	return nil
}

func (l *sliceList[T]) Remove(index int) (element T, err error) {
	if l.size == 0 {
		return element, collections.ErrEmptyList
	} else if index > l.size-1 {
		return element, collections.ErrInvalidIndex{
			RequestedIndex: index,
			Size:           l.size,
		}
	}

	element = l.data[index]
	for i := index; i < l.size-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.size--

	return element, err
}

func (l *sliceList[T]) Size() int {
	return l.size
}

func (l *sliceList[T]) SubList(start, end int) collections.List[T] {
	size := end - start
	data := make([]T, size*growthFactor)
	for i := start; i < end; i++ {
		data[i-start] = l.data[i]
	}
	return &sliceList[T]{
		data: data,
		size: size,
	}
}
