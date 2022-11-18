// ©2022 Brandon Moller

package slicelist_test

import (
	"errors"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/slicelist"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}

func TestListNewFromItems(t *testing.T) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}

	list := slicelist.NewFromItems(input)
	if list.Size() != 1000 {
		t.Fatalf("expected list size %d but got %d", 1000, list.Size())
	}
	for i := 0; i < 1000; i++ {
		if element, err := list.Get(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
	}
}

func TestListNewWithSize(t *testing.T) {
	list := slicelist.NewWithSize[int](1001)
	for i := 0; i < 1000; i++ {
		list.Add(i)
	}
	if list.Size() != 1000 {
		t.Fatalf("expected list size %d but got %d", 1000, list.Size())
	}
	for i := 0; i < 1000; i++ {
		if element, err := list.Get(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
	}

	if err := list.Insert(500, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestListAdd(t *testing.T) {
	for i := 1; i < 1001; i++ {
		list := slicelist.New[int]()
		for j := 0; j < i; j++ {
			list.Add(j)
		}
	}
}

func TestListClear(t *testing.T) {
	for i := 1; i < 1001; i++ {
		list := slicelist.New[int]()
		for j := 0; j < i; j++ {
			list.Add(j)
		}
		list.Clear()
		if list.Size() != 0 || !list.Empty() {
			t.Fatal("expected list to be empty after Clear")
		}
	}
}

func TestListEmpty(t *testing.T) {
	list := slicelist.New[int]()
	if !list.Empty() {
		t.Fatal("expected new list to be empty")
	}
	list.Add(0)
	if list.Empty() {
		t.Fatal("expected list to not be empty after adding element")
	}
}

func TestListGet(t *testing.T) {
	list := slicelist.New[int]()
	if _, err := list.Get(0); err == nil {
		t.Fatal("expected error from Get on an empty list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		list.Add(i)
	}
	for i := 0; i < 1000; i++ {
		if element, err := list.Get(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
	}

	indexError := new(collections.ErrIndexOutOfRange)
	if _, err := list.Get(1000); err == nil {
		t.Fatal("expected error from Get for element out of range")
	} else if !errors.As(err, indexError) {
		t.Fatalf("expected ErrIndexOutOfRange but got: %s", err)
	}

	list.Clear()
	if _, err := list.Get(0); err == nil {
		t.Fatal("expected error from Get after Clear on list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but received: %s", err)
	}

	list.Add(0)
	if _, err := list.Get(500); err == nil {
		t.Fatal("expected error from Get for element out of range")
	} else if !errors.As(err, indexError) {
		t.Fatalf("expected ErrIndexOutOfRange but got: %s", err)
	}
}

func TestListInsert(t *testing.T) {
	input := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		input[i] = i
	}
	expectedOutput := make([]int, 1004)
	copy(expectedOutput, input[:200])
	expectedOutput[200] = 0
	copy(expectedOutput[201:], input[200:400])
	expectedOutput[401] = 0
	copy(expectedOutput[402:], input[400:600])
	expectedOutput[602] = 0
	copy(expectedOutput[603:], input[600:800])
	expectedOutput[803] = 0
	copy(expectedOutput[804:], input[800:])

	list := slicelist.New[int]()
	for _, value := range input {
		list.Add(value)
	}

	if err := list.Insert(-1, 0); err == nil {
		t.Fatal("expected error on Insert at negative index")
	}
	if err := list.Insert(9999, 0); err == nil {
		t.Fatal("expected error on Insert outside of range")
	}

	if err := list.Insert(200, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := list.Insert(401, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := list.Insert(602, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := list.Insert(803, 0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if list.Size() != 1004 {
		t.Fatalf("expected list size %d but got %d", 1004, list.Size())
	}
	for i := 0; i < 1004; i++ {
		if element, err := list.Get(i); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != expectedOutput[i] {
			t.Fatalf("expected element with value %d but got %d", expectedOutput[i], element)
		}
	}
}

func TestListRemove(t *testing.T) {
	list := slicelist.New[int]()

	if _, err := list.Remove(0); err == nil {
		t.Fatal("expected error on call to Remove on empty list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		list.Add(i)
	}

	indexErr := new(collections.ErrIndexOutOfRange)
	if _, err := list.Remove(9999); err == nil {
		t.Fatal("expected error on call to Remove index out of range")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		if element, err := list.Remove(0); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element with value %d but got %d", i, element)
		}
	}

	if _, err := list.Remove(0); err == nil {
		t.Fatal("expected error on call to Remove on empty list")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList but got: %s", err)
	}

	for i := 0; i < 1000; i++ {
		list.Add(i)
	}

	for i := 999; i > 0; i-- {
		if _, err := list.Remove(rand.Intn(i)); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}
	if _, err := list.Remove(0); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestListSize(t *testing.T) {
	for i := 1; i < 1001; i++ {
		list := slicelist.New[int]()
		for j := 0; j < i; j++ {
			list.Add(j)
		}
		if list.Size() != i {
			t.Fatalf("expected list size %d but got %d", i, list.Size())
		}
	}
}

func TestListSubList(t *testing.T) {
	l1 := slicelist.New[int]()
	if _, err := l1.SubList(0, 0); err == nil {
		t.Fatal("expected error from SubList on new List")
	} else if !errors.Is(err, collections.ErrEmptyList) {
		t.Fatalf("expected ErrEmptyList, got %T", err)
	}

	for i := 0; i < 1000; i++ {
		l1.Add(i)
	}
	l2, err := l1.SubList(250, 750)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else if l2.Size() != 500 {
		t.Fatalf("expected sublist size of %d but got %d", 500, l2.Size())
	}
	for i := 250; i < 750; i++ {
		if element, err := l2.Get(i - 250); err != nil {
			t.Fatalf("unexpected error: %s", err)
		} else if element != i {
			t.Fatalf("expected element %d but not %d", i, element)
		}
	}

	rangeErr := new(collections.ErrInvalidRange)
	if _, err := l1.SubList(-1, 10); err == nil {
		t.Fatal("expected error from SubList with start less than 0")
	} else if !errors.As(err, rangeErr) {
		t.Fatalf("expected ErrInvalidRange, got %T", err)
	}
	if _, err := l1.SubList(10, 0); err == nil {
		t.Fatal("expected error from SubList with end less than start")
	} else if !errors.As(err, rangeErr) {
		t.Fatalf("expected ErrInvalidRange, got %T", err)
	}

	indexErr := new(collections.ErrIndexOutOfRange)
	if _, err := l1.SubList(1000, 1005); err == nil {
		t.Fatal("expected error from SubList with start == size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
	if _, err := l1.SubList(0, 1005); err == nil {
		t.Fatal("expected error from SubList with end > size")
	} else if !errors.As(err, indexErr) {
		t.Fatalf("expected ErrIndexOutOfRange, got %T", err)
	}
}
