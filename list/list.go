package list

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

func (l *sliceList[T]) ensureCapacity(count int) {
	if l.size+count >= cap(l.data) {
		newData := make([]T, (l.size+count)*growthFactor)
		copy(newData, l.data)
		l.data = newData
	}
}

func (l *sliceList[T]) Add(item T) T {
	l.ensureCapacity(1)
	l.data[l.size] = item
	l.size++

	return l.data[l.size-1]
}

func (l *sliceList[T]) AddAll(items []T) int {
	oldSize := l.size
	l.ensureCapacity(len(items))
	l.size += copy(l.data[l.size:], items)

	return l.size - oldSize
}

func (l *sliceList[T]) Clear() {
	l.size = 0
}

func (l *sliceList[T]) Contains(item T) bool {
	for i := 0; i < l.size; i++ {
		if l.data[i] == item {
			return true
		}
	}

	return false
}

func (l *sliceList[T]) ContainsAll(items []T) bool {
	for i := 0; i < len(items); i++ {
		if !l.Contains(items[i]) {
			return false
		}
	}

	return true
}

func (l *sliceList[T]) Empty() bool {
	return l.size == 0
}

func (l *sliceList[T]) Equals(l2 collections.List[T]) bool {
	if l.size != l2.Size() {
		return false
	}

	for i := 0; i < l.size; i++ {
		if element, _ := l2.Get(i); l.data[i] != element {
			return false
		}
	}

	return true
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

	return
}

func (l *sliceList[T]) IndexOf(item T) int {
	for i := 0; i < l.size; i++ {
		if l.data[i] == item {
			return i
		}
	}

	return -1
}

func (l *sliceList[T]) Insert(index int, item T) (element T, err error) {
	// to avoid copying elements twice we intentionally don't use ensureCapacity
	if index < 0 || index > l.size {
		err = collections.ErrInvalidIndex{
			RequestedIndex: index,
			Size:           l.size,
		}
	} else {
		newSize := l.size
		if l.size+1 >= cap(l.data) {
			newSize = l.size * growthFactor
		}
		newData := make([]T, newSize)
		copy(newData, l.data[:index])
		newData[index] = item
		copy(newData[index+1:], l.data[index:])
		l.data = newData
		element = l.data[index]
	}

	return
}

func (l *sliceList[T]) LastIndexOf(item T) int {
	if l.size == 0 {
		return -1
	}

	for i := l.size - 1; i > -1; i-- {
		if l.data[i] == item {
			return i
		}
	}

	return -1
}

func (l *sliceList[T]) Remove(item T) (index int) {
	index = -1
	if l.size == 0 {
		return
	}

	for i := 0; i < l.size; i++ {
		if l.data[i] == item {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	for i := index; i < l.size-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.size--

	return
}
