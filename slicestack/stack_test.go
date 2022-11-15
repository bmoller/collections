package slicestack_test

import (
	"testing"

	"github.com/bmoller/collections/slicestack"
)

func TestStackNewWithSize(t *testing.T) {
	stack := slicestack.NewWithSize[int](1000)

	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	if stack.Size() != 1000 {
		t.Fatalf("expected stack size %d but got size %d", 1000, stack.Size())
	}
	for i := 999; i > -1; i-- {
		if element, err := stack.Peek(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected value %d from Peek but got %d", i, element)
		}
		if element, err := stack.Pop(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected value %d from Pop but got %d", i, element)
		}
		if stack.Size() != i {
			t.Fatalf("expected stack size %d but got %d", i, stack.Size())
		}
	}
	if !stack.Empty() {
		t.Fatal("expected empty stack after popping all elements")
	}
}

func TestStackEmpty(t *testing.T) {
	stack := slicestack.New[int]()
	if !stack.Empty() {
		t.Fatal("expected new Stack to be empty")
	}
	stack.Push(0)
	if stack.Empty() {
		t.Fatal("expected Stack not to be empty after Push")
	}
}

func TestStackPeek(t *testing.T) {
	stack := slicestack.New[int]()

	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	for i := 99; i > -1; i-- {
		if element, err := stack.Peek(); err != nil {
			t.Fatalf("expected element on Peek call %d but received error: %s", i, err)
		} else {
			if element != i {
				t.Fatalf("expected element %d but got %d", i, element)
			}
		}
		stack.Pop()
	}
	if _, err := stack.Peek(); err == nil {
		t.Fatal("Peek on an empty stack should return an error")
	}
}

func TestStackPop(t *testing.T) {
	stack := slicestack.New[int]()

	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	for i := 99; i > -1; i-- {
		if element, err := stack.Pop(); err != nil {
			t.Fatalf("expected element on Pop call %d but received error: %s", i, err)
		} else if element != i {
			t.Fatalf("expected element %d but got %d", i, element)
		} else if stack.Size() != i {
			t.Fatalf("expected size %d but got %d", i, stack.Size())
		}
	}
	if _, err := stack.Pop(); err == nil {
		t.Fatal("Pop on an empty stack should return an error")
	}
}

func TestStackPush(t *testing.T) {
	stack := slicestack.New[int]()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
}

func TestStackSize(t *testing.T) {
	stack := slicestack.New[int]()
	for i := 1; i < 1000; i++ {
		stack.Push(i)
		if stack.Size() != i {
			t.Fatalf("expected stack size %d but got %d", i, stack.Size())
		}
	}
}
