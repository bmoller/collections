// ©2022 Brandon Moller

/*
Package linkedlist is an implementation of [collections.LinkedList] backed by individual list nodes.
The list is doubly-linked and can be traversed in either direction from any node in the list.
For methods with nodes as their parameters, lists verify that the nodes are members of the receiving list.
*/
package linkedlist

import "github.com/bmoller/collections"

type listNode[T comparable] struct {
	elementOf *linkedList[T]
	next      *listNode[T]
	previous  *listNode[T]
	value     T
}

func (n *listNode[T]) Next() collections.ListNode[T] {
	return n.next
}

func (n *listNode[T]) Previous() collections.ListNode[T] {
	return n.previous
}

func (n *listNode[T]) Value() T {
	return n.value
}

type linkedList[T comparable] struct {
	head *listNode[T]
	size int
	tail *listNode[T]
}

func New[T comparable]() collections.LinkedList[T] {
	return new(linkedList[T])
}

func (l *linkedList[T]) Add(item T) {
	node := &listNode[T]{
		elementOf: l,
		value:     item,
	}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		oldTail := l.tail
		oldTail.next = node
		node.previous = oldTail
		l.tail = node
	}
	l.size++
}

func (l *linkedList[T]) Clear() {
	l.head, l.tail = nil, nil
	l.size = 0
}

func (l *linkedList[T]) Empty() bool {
	return l.size == 0
}

func (l *linkedList[T]) Get(index int) (element T, err error) {
	if l.size == 0 {
		return element, collections.ErrEmptyList
	} else if index >= l.size || index < 0 {
		return element, collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
	}

	node := (collections.ListNode[T])(l.head)
	for i := 0; i < index; i++ {
		node = node.Next()
	}

	return node.Value(), err
}

func (l *linkedList[T]) GetNode(index int) (collections.ListNode[T], error) {
	if l.size == 0 {
		return nil, collections.ErrEmptyList
	} else if index >= l.size || index < 0 {
		return nil, collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
	}

	var node collections.ListNode[T] = l.head
	for i := 0; i < index; i++ {
		node = node.Next()
	}

	return node, nil
}

func (l *linkedList[T]) Head() collections.ListNode[T] {
	return l.head
}

func (l *linkedList[T]) Insert(index int, item T) error {
	switch {
	case l.size == 0:
		return collections.ErrEmptyList
	case index > l.size, index < 0:
		err := collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
		return err
	case index == 0:
		node := &listNode[T]{
			elementOf: l,
			next:      l.head,
			value:     item,
		}
		l.head.previous = node
		l.head = node
	default:
		current := l.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		node := &listNode[T]{
			elementOf: l,
			next:      current.next,
			previous:  current,
			value:     item,
		}
		current.next.previous = node
		current.next = node
	}
	l.size++

	return nil
}

func (l *linkedList[T]) InsertAfter(node collections.ListNode[T], item T) (collections.ListNode[T], error) {
	typedNode, ok := node.(*listNode[T])
	if !ok {
		return nil, collections.ErrWrongNodeType
	} else if typedNode.elementOf != l {
		return nil, collections.ErrNodeIsNotElement
	}

	newNode := &listNode[T]{
		elementOf: l,
		previous:  typedNode,
		next:      typedNode.next,
		value:     item,
	}
	if l.tail == typedNode {
		l.tail = newNode
	} else {
		typedNode.next.previous = newNode
	}
	typedNode.next = newNode
	l.size++

	return newNode, nil
}

func (l *linkedList[T]) InsertBefore(node collections.ListNode[T], item T) (collections.ListNode[T], error) {
	typedNode, ok := node.(*listNode[T])
	if !ok {
		return nil, collections.ErrWrongNodeType
	} else if typedNode.elementOf != l {
		return nil, collections.ErrNodeIsNotElement
	}

	newNode := &listNode[T]{
		elementOf: l,
		previous:  typedNode.previous,
		next:      typedNode,
		value:     item,
	}
	if l.head == typedNode {
		l.head = newNode
	} else {
		typedNode.previous.next = newNode
	}
	typedNode.previous = newNode
	l.size++

	return newNode, nil
}

func (l *linkedList[T]) Remove(index int) (element T, err error) {
	if l.size == 0 {
		return element, collections.ErrEmptyList
	} else if index >= l.size || index < 0 {
		err := collections.ErrIndexOutOfRange{
			Index: index,
			Size:  l.size,
		}
		return element, err
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	if current.next != nil {
		current.next.previous = current.previous
	}
	if current.previous != nil {
		current.previous.next = current.next
	}
	if l.head == current {
		l.head = current.next
	}
	l.size--

	return current.value, nil
}

func (l *linkedList[T]) RemoveNode(node collections.ListNode[T]) error {
	typedNode, ok := node.(*listNode[T])
	if !ok {
		return collections.ErrWrongNodeType
	} else if typedNode.elementOf != l {
		return collections.ErrNodeIsNotElement
	}

	if typedNode.next != nil {
		typedNode.next.previous = typedNode.previous
	}
	if typedNode.previous != nil {
		typedNode.previous.next = typedNode.next
	}
	if l.head == typedNode {
		l.head = typedNode.next
	}
	if l.tail == typedNode {
		l.tail = typedNode.previous
	}

	return nil
}

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) SubList(start int, end int) (collections.List[T], error) {
	switch {
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

	current := l.head
	for i := 0; i < start; i++ {
		current = current.next
	}

	list := new(linkedList[T])
	for i := start; i < end; i++ {
		list.Add(current.value)
		current = current.next
	}

	return list, nil
}

func (l *linkedList[T]) Tail() collections.ListNode[T] {
	return l.tail
}
