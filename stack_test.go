package collections_test

import (
	"errors"
	"testing"

	"github.com/bmoller/collections"
)

func TestEmptyStack(t *testing.T) {
	stack := collections.NewStack[int]()

	if _, err := stack.Peek(); err == nil {
		t.Fatal("Peek on an empty stack should return an error")
	} else if !errors.Is(err, collections.ErrEmptyStack) {
		t.Fatalf("expected error %T but got %T", collections.ErrEmptyStack, err)
	}
}

func TestStackSearch(t *testing.T) {
	stack := collections.NewStack[int]()

	for _, i := range []int{-1, 9384, 2, 43, -341325, 5892, 132491324} {
		stack.Push(i)
	}
	index := stack.Search(132491324)
	if index != 0 {
		t.Fatalf("expected element %d to have index %d but got %d", 132491324, 0, index)
	}
	index = stack.Search(47)
	if index != -1 {
		t.Fatalf("expected to not find element %d but got index %d", 47, index)
	}
}
