package linkedlist_test

import (
	"errors"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/linkedlist"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
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

	if node == nil {
		t.Fatal("expected a valid node pointer but got nil")
	} else if node != list.Tail() {
		t.Fatal("final node does not match list tail")
	} else if next := node.Next(); next != nil {
		t.Errorf("type: %T", next)
		t.Errorf("next node: %v", next)
		t.Fatal("expected tail node to have nil pointer to next node")
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
