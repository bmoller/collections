package linkedstack_test

import (
	"errors"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/linkedstack"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}

func TestStackEmpty(t *testing.T) {
	stack := linkedstack.New[int]()
	if !stack.Empty() {
		t.Fatal("expected new Stack to be empty")
	}
	stack.Push(0)
	if stack.Empty() {
		t.Fatal("expected Stack not to be empty after pushing an element")
	}
}

func TestStackPeek(t *testing.T) {
	stack := linkedstack.New[int]()
	if _, err := stack.Peek(); err == nil {
		t.Fatal("expected error from Peek on a new Stack")
	} else if !errors.Is(err, collections.ErrEmptyStack) {
		t.Fatalf("expected ErrEmptyStack, got %T", err)
	}

	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	for i := 999; i > -1; i-- {
		switch element, err := stack.Peek(); {
		case err != nil:
			t.Fatalf("unexpected error: %s", err)
		case element != i:
			t.Fatalf("expected element with value %d, got %d", i, element)
		case stack.Size() != i+1:
			t.Fatalf("expected Stack size %d, got %d", i+1, stack.Size())
		}
		stack.Pop()
	}
}

func TestStackPop(t *testing.T) {
	stack := linkedstack.New[int]()
	if _, err := stack.Pop(); err == nil {
		t.Fatal("expected error from Pop on a new Stack")
	} else if !errors.Is(err, collections.ErrEmptyStack) {
		t.Fatalf("expected ErrEmptyStack, got %T", err)
	}

	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	for i := 999; i > -1; i-- {
		switch element, err := stack.Pop(); {
		case err != nil:
			t.Fatalf("unexpected error: %s", err)
		case element != i:
			t.Fatalf("expected element with value %d, got %d", i, element)
		case stack.Size() != i:
			t.Fatalf("expected Stack size %d, got %d", i, stack.Size())
		}
	}
}

func TestStackPush(t *testing.T) {
	stack := linkedstack.New[int]()
	for i := 0; i < 1000; i++ {
		stack.Push(rand.Int())
	}
}

func TestStackSize(t *testing.T) {
	stack := linkedstack.New[int]()
	if stack.Size() != 0 {
		t.Fatalf("expected new Stack to have size %d, got %d", 0, stack.Size())
	}
	for i := 1; i < 1001; i++ {
		stack.Push(i)
		if stack.Size() != i {
			t.Fatalf("expected size %d, got %d", i, stack.Size())
		}
	}
}
