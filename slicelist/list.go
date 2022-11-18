/*
Package slicelist provides an array/slice-backed implementation of [collections.List].

Whenever the List grows beyond the bounds of its current backing storage a new slice is created and all elements are copied.
*/

package slicelist

import "github.com/bmoller/collections"

const (
	growthFactor int = 2
	initialSize  int = 100
)

type list[T comparable] struct {
	data []T
	size int
}

func New[T comparable]() collections.List[T] {
	return &list[T]{
		data: make([]T, initialSize),
	}
}

/*
NewFromItems creates a new List with all elements of items as its contents.
The order of items is preserved.
*/
func NewFromItems[T comparable](items []T) collections.List[T] {
	return &list[T]{
		data: items,
		size: len(items),
	}
}

/*
NewWithSize allows the user control over the initial size of the backing slice.
A new List is created and returned with size as its capacity.
*/
func NewWithSize[T comparable](size int) collections.List[T] {
	return &list[T]{
		data: make([]T, size),
	}
}

func (l *list[T]) Add(item T) {
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

func (l *list[T]) Clear() {
	l.size = 0
}

func (l *list[T]) Empty() bool {
	return l.size == 0
}

func (l *list[T]) Get(index int) (item T, err error) {
	switch {
	case l.size == 0:
		err = collections.ErrEmptyList
	case index >= l.size:
		err = collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
	default:
		item = l.data[index]
	}

	return item, err
}

func (l *list[T]) Insert(index int, item T) error {
	if index < 0 || index > l.size {
		return collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
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

func (l *list[T]) Remove(index int) (element T, err error) {
	if l.size == 0 {
		return element, collections.ErrEmptyList
	} else if index > l.size-1 {
		return element, collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
	}

	element = l.data[index]
	for i := index; i < l.size-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.size--

	return element, err
}

func (l *list[T]) Size() int {
	return l.size
}

func (l *list[T]) SubList(start, end int) (collections.List[T], error) {
	switch {
	case l.size == 0:
		return nil, collections.ErrEmptyList
	case start < 0 || end < start:
		return nil, collections.ErrInvalidRange{
			End:   end,
			Start: start,
		}
	case start >= l.size:
		return nil, collections.ErrIndexOutOfRange{
			Index: start,
			Size:  l.size,
		}
	case end > l.size:
		return nil, collections.ErrIndexOutOfRange{
			Index: end,
			Size:  l.size,
		}
	}

	size := end - start
	data := make([]T, size*growthFactor)
	for i := start; i < end; i++ {
		data[i-start] = l.data[i]
	}
	return &list[T]{
		data: data,
		size: size,
	}, nil
}
