// ©2022 Brandon Moller

package linkedqueue_test

import (
	"errors"
	"testing"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/linkedqueue"
)

func TestQueueEmpty(t *testing.T) {
	queue := linkedqueue.New[int]()
	if !queue.Empty() {
		t.Fatal("expected new queue to be empty")
	}
	queue.Push(1)
	if queue.Empty() {
		t.Fatal("expected queue to not be empty after pushing an item")
	}
}

func TestQueuePeek(t *testing.T) {
	queue := linkedqueue.New[int]()
	if _, err := queue.Peek(); err == nil {
		t.Fatal("expected error from Peek on new queue")
	} else if !errors.Is(err, collections.ErrEmptyQueue) {
		t.Fatalf("expected ErrEmptyQueue but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		queue.Push(i)
	}
	for i := 0; i < 1000; i++ {
		if element, err := queue.Peek(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
		queue.Pop()
	}
}

func TestQueuePop(t *testing.T) {
	queue := linkedqueue.New[int]()
	for i := 0; i < 1000; i++ {
		queue.Push(i)
	}
	for i := 0; i < 1000; i++ {
		if element, err := queue.Pop(); err != nil {
			t.Fatalf("expected element from pop but received error: %s", err)
		} else if element != i {
			t.Fatalf("expected element %d but got %d", i, element)
		}
	}
	if _, err := queue.Pop(); err == nil || !errors.Is(err, collections.ErrEmptyQueue) {
		t.Fatalf("expected ErrEmptyQueue from Pop on empty stack but got: %s", err)
	}
}

func TestQueueSize(t *testing.T) {
	queue := linkedqueue.New[int]()
	for i := 1; i < 1001; i++ {
		queue.Push(i)
		if queue.Size() != i {
			t.Fatalf("expected queue size %d but got %d", i, queue.Size())
		}
	}
}
