// ©2022 Brandon Moller

package linkedlist_test

import (
	"errors"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/linkedlist"
)

type badNode[T comparable] struct {
	next     *badNode[T]
	previous *badNode[T]
	value    T
}

func (b *badNode[T]) Next() collections.ListNode[T] {
	return b.next
}

func (b *badNode[T]) Previous() collections.ListNode[T] {
	return b.previous
}

func (b *badNode[T]) Value() T {
	return b.value
}

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}

func TestLinkedListBackward(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}

	node := list.Tail()
	for i := 0; i < 999; i++ {
		node = node.Previous()
	}

	switch {
	case node == nil:
		t.Fatal("expected a valid node pointer but got nil")
	case node != list.Head():
		t.Fatal("final node does not match list head")
	case !reflect.ValueOf(node.Previous()).IsNil():
		// we reflect here because nodes work around pointers but signatures reference the interface
		t.Fatal("expected head node to have nil pointer to previous node")
	}
}

func TestLinkedListForward(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}

	node := list.Head()
	for i := 0; i < 999; i++ {
		node = node.Next()
	}

	switch {
	case node == nil:
		t.Fatal("expected a valid node pointer but got nil")
	case node != list.Tail():
		t.Fatal("final node does not match list tail")
	case !reflect.ValueOf(node.Next()).IsNil():
		// we reflect here because nodes work around pointers but signatures reference the interface
		t.Fatal("expected tail node to have nil pointer to next node")
	}
}

func TestLinkedListClear(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}
	list.Clear()

	switch {
	case list.Size() != 0:
		t.Fatalf("expected list size %d but got %d", 0, list.Size())
	case !reflect.ValueOf(list.Head()).IsNil():
		t.Fatal("expected list to have nil head pointer")
	case !reflect.ValueOf(list.Tail()).IsNil():
		t.Fatal("expected list to have nil tail pointer")
	}
}

func TestLinkedListEmpty(t *testing.T) {
	list := linkedlist.New[int]()
	if !list.Empty() {
		t.Fatal("expected new list to be empty")
	}
	list.Add(0)
	if list.Empty() {
		t.Fatal("expected list to not be empty afterr adding element")
	}
}

func TestLinkedListGet(t *testing.T) {
	list := linkedlist.New[int]()
	if _, err := list.Get(0); err == nil || !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		list.Add(i)
	}
	for i := 0; i < 1000; i++ {
		if element, err := list.Get(i); err != nil {
			t.Fatalf("got unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element to have value %d but got %d", i, element)
		}
	}

	errIndex := new(collections.ErrIndexOutOfRange)
	if _, err := list.Get(9999); err == nil {
		t.Fatal("expected error from call to Get out of range")
	} else if !errors.As(err, errIndex) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
	if _, err := list.Get(-1); err == nil {
		t.Fatal("expected error from call to Get with index < 0")
	} else if !errors.As(err, errIndex) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
}

func TestLinkedListGetNode(t *testing.T) {
	list := linkedlist.New[int]()
	if _, err := list.GetNode(0); err == nil {
		t.Fatal("expected error from call to GetNode on new list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 249; j++ {
			list.Add(rand.Int())
			list.Add(0)
		}
	}

	for _, index := range []int{249, 499, 749, 999} {
		if node, err := list.GetNode(index); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if node.Value() != 0 {
			t.Fatalf("expected node with value %d but got %d", 0, node.Value())
		}
	}

	indexErr := new(collections.ErrIndexOutOfRange)
	if _, err := list.GetNode(9999); err == nil {
		t.Fatal("expected error from call to Get with index > size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrInvalidIndex but got: %s", err)
	}
	if _, err := list.GetNode(-1); err == nil {
		t.Fatal("expected error from call to Get with index < 0")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrInvalidIndex but got: %s", err)
	}
}

func TestLinkedListHead(t *testing.T) {
	list := linkedlist.New[int]()
	if node := list.Head(); !reflect.ValueOf(node).IsNil() {
		t.Fatal("expected head node of new list to be nil")
	}

	list.Add(0)
	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}

	if value := list.Head().Value(); value != 0 {
		t.Fatalf("expected head node to have value %d but got %d", 0, value)
	}
}

func TestLinkedListInsert(t *testing.T) {
	list := linkedlist.New[int]()
	if err := list.Insert(0, 0); err == nil {
		t.Fatal("expected error from Insert on an empty List")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList, got %T", err)
	}

	list.Add(999)
	for i := 998; i > -1; i-- {
		if err := list.Insert(0, i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
	node := list.Head()
	for i := 0; i < 1000; i++ {
		if node.Value() != i {
			t.Fatalf("expected node with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}

	indexErr := new(collections.ErrIndexOutOfRange)
	if err := list.Insert(-1, 0); err == nil {
		t.Fatal("expected error from Insert with index < 0")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
	if err := list.Insert(9999, 0); err == nil {
		t.Fatal("expected error from Insert with index > size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}

	for i := 1; i < 101; i++ {
		if err := list.Insert(rand.Intn(1000), rand.Int()); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if list.Size() != 1000+i {
			t.Fatalf("expected size %d, got %d", 1000+i, list.Size())
		}
	}
}

func TestLinkedListInsertAfter(t *testing.T) {
	l1 := linkedlist.New[int]()
	for i := 0; i < 100; i++ {
		l1.Add(i)
	}

	node := l1.Head()
	for i := 0; i < 50; i++ {
		node = node.Next()
	}
	if newNode, err := l1.InsertAfter(node, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if newNode.Value() != 0 {
		t.Fatalf("expected element with value %d, got %d", 0, newNode.Value())
	}
	if l1.Size() != 101 {
		t.Fatalf("expected list size %d, got %d", 101, l1.Size())
	}

	node = l1.Head()
	for i := 0; i < 51; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}
	if node.Value() != 0 {
		t.Fatalf("expected element with value %d, got %d", 0, node.Value())
	}
	node = node.Next()
	for i := 51; i < 100; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}

	l2 := linkedlist.New[int]()
	l2.Add(0)
	node = l2.Head()
	if _, err := l1.InsertAfter(node, 0); err == nil {
		t.Fatal("expected error from InsertAfter with foreign node")
	} else if !errors.Is(err, collections.ErrNodeIsNotElement) {
		t.Fatalf("expected ErrNodeIsNotElement, got %T", err)
	}

	node = l1.Tail()
	if newNode, err := l1.InsertAfter(node, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if newNode != l1.Tail() {
		t.Fatal("expected inserted node to be the new tail")
	}

	bad := &badNode[int]{
		value: 0,
	}
	if _, err := l1.InsertAfter(bad, 0); err == nil {
		t.Fatal("expected error from InsertAfter with bad node type")
	} else if !errors.Is(err, collections.ErrWrongNodeType) {
		t.Fatalf("expected ErrWrongNodeType, got %T", err)
	}
}

func TestLinkedListInsertBefore(t *testing.T) {
	l1 := linkedlist.New[int]()
	for i := 0; i < 100; i++ {
		l1.Add(i)
	}

	node := l1.Head()
	for i := 0; i < 50; i++ {
		node = node.Next()
	}
	if newNode, err := l1.InsertBefore(node, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if newNode.Value() != 0 {
		t.Fatalf("expected element with value %d, got %d", 0, newNode.Value())
	}
	if l1.Size() != 101 {
		t.Fatalf("expected list size %d, got %d", 101, l1.Size())
	}

	node = l1.Head()
	for i := 0; i < 50; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}
	if node.Value() != 0 {
		t.Fatalf("expected element with value %d, got %d", 0, node.Value())
	}
	node = node.Next()
	for i := 50; i < 100; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}

	l2 := linkedlist.New[int]()
	l2.Add(0)
	node = l2.Head()
	if _, err := l1.InsertBefore(node, 0); err == nil {
		t.Fatal("expected error from InsertAfter with foreign node")
	} else if !errors.Is(err, collections.ErrNodeIsNotElement) {
		t.Fatalf("expected ErrNodeIsNotElement, got %T", err)
	}

	node = l1.Head()
	if newNode, err := l1.InsertBefore(node, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if newNode != l1.Head() {
		t.Fatal("expected inserted node to be the new head")
	}

	bad := &badNode[int]{
		value: 0,
	}
	if _, err := l1.InsertBefore(bad, 0); err == nil {
		t.Fatal("expected error from InsertBefore with bad node type")
	} else if !errors.Is(err, collections.ErrWrongNodeType) {
		t.Fatalf("expected ErrWrongNodeType, got %T", err)
	}
}

func TestLinkedListRemove(t *testing.T) {
	list := linkedlist.New[int]()
	if _, err := list.Remove(0); err == nil {
		t.Fatal("expected error from call to Remove on empty list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		list.Add(i)
	}

	indexErr := new(collections.ErrIndexOutOfRange)
	if _, err := list.Remove(9999); err == nil {
		t.Fatal("expected error from call to Remove with out of range index")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrInvalidIndex but got: %s", err)
	}

	for i := 999; i > -1; i-- {
		if element, err := list.Remove(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
	}

	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}
	for i := 0; i < 999; i++ {
		if _, err := list.Remove(rand.Intn(list.Size())); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
}

func TestLinkedListRemoveNode(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}

	bad := &badNode[int]{}
	if err := list.RemoveNode(bad); err == nil {
		t.Fatal("expected error from removing wrong node type")
	} else if !errors.Is(err, collections.ErrWrongNodeType) {
		t.Fatalf("expected ErrWrongNodeType, got %T", err)
	}

	l2 := linkedlist.New[int]()
	l2.Add(0)
	if err := list.RemoveNode(l2.Head()); err == nil {
		t.Fatal("expected error when trying to remove non-member node")
	} else if !errors.Is(err, collections.ErrNodeIsNotElement) {
		t.Fatalf("expected ErrNodeIsNotElement, got %T", err)
	}

	node := list.Head()
	for i := 0; i < 50; i++ {
		node = node.Next()
	}
	if err := list.RemoveNode(node); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	node = list.Head()
	for i := 0; i < 50; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}
	for i := 51; i < 100; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}

	if err := list.RemoveNode(list.Head()); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := list.RemoveNode(list.Tail()); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	node = list.Head()
	for i := 1; i < 50; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}
	for i := 51; i < 99; i++ {
		if node.Value() != i {
			t.Fatalf("expected element with value %d, got %d", i, node.Value())
		}
		node = node.Next()
	}
	if list.Tail().Value() != 98 {
		t.Fatalf("expected element with value %d, got %d", 98, list.Tail().Value())
	}
}

func TestLinkedListSize(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 1; i < 1001; i++ {
		list.Add(i)
		if list.Size() != i {
			t.Fatalf("expected list size %d but got %d", i, list.Size())
		}
	}
}

func TestLinkedListSubList(t *testing.T) {
	list := linkedlist.New[int]()
	for i := 0; i < 100; i++ {
		list.Add(i)
	}

	invalidErr := new(collections.ErrInvalidRange)
	indexErr := new(collections.ErrIndexOutOfRange)
	if _, err := list.SubList(-1, 50); err == nil {
		t.Fatal("expected error from SubList with negative start index")
	} else if !errors.As(err, invalidErr) {
		t.Fatalf("expected ErrInvalidRange, got %T", err)
	}
	if _, err := list.SubList(20, 10); err == nil {
		t.Fatal("expected error from SubList with end < start")
	} else if !errors.As(err, invalidErr) {
		t.Fatalf("expected ErrInvalidRange, got %T", err)
	}
	if _, err := list.SubList(110, 120); err == nil {
		t.Fatal("expected error from SubList with start > size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
	if _, err := list.SubList(50, 120); err == nil {
		t.Fatal("expected error from SubList with end > size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}

	subList, err := list.SubList(11, 30)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if subList.Size() != 19 {
		t.Fatalf("expected SubList with size %d, got %d", 19, subList.Size())
	}
	for i := 11; i < 30; i++ {
		if element, err := subList.Get(i - 11); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d, got %d", i-11, element)
		}
	}

	if _, ok := subList.(collections.LinkedList[int]); !ok {
		t.Fatal("expected to be able to cast List to LinkedList")
	}
}

func TestLinkedListTail(t *testing.T) {
	list := linkedlist.New[int]()
	if node := list.Tail(); !reflect.ValueOf(node).IsNil() {
		t.Fatal("expected tail node of new list to be nil")
	}

	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}
	list.Add(0)

	if value := list.Tail().Value(); value != 0 {
		t.Fatalf("expected tail node to have value %d but got %d", 0, value)
	}
}
