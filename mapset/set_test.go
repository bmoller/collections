package mapset_test

import (
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/bmoller/collections"
	"github.com/bmoller/collections/mapset"
)

func TestSetAdd1000(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000; i++ {
		testSet.Add(i)
	}

	if testSet.Size() != 1000 {
		t.Fatalf("expected size %d but got %d", 1000, testSet.Size())
	}
}

func TestSetAdd1000000(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000000; i++ {
		testSet.Add(i)
	}
	if testSet.Size() != 1000000 {
		t.Fatalf("expected size %d but got %d", 1000000, testSet.Size())
	}
}

func TestSetContains1000(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 1000; i++ {
		if !testSet.Contains(i) {
			t.Fatal("set does not contain an added element")
		}
	}
}

func TestSetEmpty(t *testing.T) {
	testSet := mapset.New[int]()
	if !testSet.Empty() {
		t.Fatal("new mapset does not report as empty")
	}
	testSet.Add(0)
	if testSet.Empty() {
		t.Fatal("set reports as empty after adding an element")
	}
	for i := 1; i < 1000; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 1000; i++ {
		testSet.Pop()
	}
	if !testSet.Empty() {
		t.Fatal("set does not report as empty after popping all elements")
	}
}

func TestSetIterator(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000; i++ {
		testSet.Add(i)
	}
	itr := testSet.Iterator()
	for i := 0; i < 1000; i++ {
		if _, err := itr(); err != nil {
			t.Fatal("expected an element but received error")
		}
	}
	if _, err := itr(); err == nil || !errors.Is(err, collections.ErrNoMoreItems) {
		t.Fatalf("exhausted iterator should return ErrNoMoreItems but got %v", err)
	}
}

func TestSetPop(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 1000; i++ {
		if _, err := testSet.Pop(); err != nil {
			t.Fatalf("failed to pop element from mapset on call %d", i)
		}
	}
	if _, err := testSet.Pop(); err == nil {
		t.Fatal("shouldn't be able to pop from an empty mapset")
	}
}

func TestSetRemove(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 0; i < 1000; i++ {
		testSet.Add(i)
	}
	for i := 0; i < 1000; i++ {
		testSet.Remove(i)
	}
	if testSet.Size() != 0 {
		t.Fatal("expected empty set after removing all elements")
	}
}

func TestSetSize(t *testing.T) {
	testSet := mapset.New[int]()
	for i := 1; i < 1000; i++ {
		testSet.Add(i)
		if testSet.Size() != i {
			t.Fatalf("expected size %d but got size %d", i, testSet.Size())
		}
	}
}

func TestSetUnion(t *testing.T) {
	itemsA := []int{1, 2, 3, 4, 5}
	itemsB := []int{6, 7, 8, 9, 10}

	a := mapset.New[int]()
	for _, item := range itemsA {
		a.Add(item)
	}
	b := mapset.New[int]()
	for _, item := range itemsB {
		b.Add(item)
	}
	c := mapset.Union(a, b)
	if c.Size() != len(itemsA)+len(itemsB) {
		t.Fatalf("expected result set size %d but got %d", len(itemsA)+len(itemsB), c.Size())
	}
	for _, item := range itemsA {
		if !c.Contains(item) {
			t.Fatalf("result set missing expected element %d", item)
		}
	}
	for _, item := range itemsB {
		if !c.Contains(item) {
			t.Fatalf("result set missing expected element %d", item)
		}
	}
}

func TestSetIntersection(t *testing.T) {
	itemsA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemsB := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	commonItems := []int{6, 7, 8, 9, 10}

	a := mapset.New[int]()
	for _, item := range itemsA {
		a.Add(item)
	}
	b := mapset.New[int]()
	for _, item := range itemsB {
		b.Add(item)
	}
	c := mapset.Intersection(a, b)
	if c.Size() != 5 {
		t.Fatalf("expected result set size %d but got %d", 5, c.Size())
	}
	for _, item := range commonItems {
		if !c.Contains(item) {
			t.Fatalf("result set missing expected element %d", item)
		}
	}
}

func TestSetDifference(t *testing.T) {
	itemsA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemsB := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	uniqueItems := []int{1, 2, 3, 4, 5}

	a := mapset.New[int]()
	for _, item := range itemsA {
		a.Add(item)
	}
	b := mapset.New[int]()
	for _, item := range itemsB {
		b.Add(item)
	}
	c := mapset.Difference(a, b)
	if c.Size() != 5 {
		t.Fatalf("expected result set size %d but got %d", 5, c.Size())
	}
	for _, item := range uniqueItems {
		if !c.Contains(item) {
			t.Fatalf("result set missing expected element %d", item)
		}
	}
}

func TestSetIsSubset(t *testing.T) {
	itemsA := []int{1, 2, 3, 4}
	itemsB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a := mapset.New[int]()
	for _, item := range itemsA {
		a.Add(item)
	}
	b := mapset.New[int]()
	for _, item := range itemsB {
		b.Add(item)
	}
	if !mapset.IsSubset(a, b) {
		t.Fatal("expected a to be subset of b")
	}
	if mapset.IsSubset(b, a) {
		t.Fatal("expected b to not be subset of a")
	}
}

// benchmarks

func BenchmarkAddRandInt1000(b *testing.B) {
	rand.Seed(time.Now().UTC().Unix())
	testSet := mapset.New[int]()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			testSet.Add(rand.Int())
		}
	}
}

func BenchmarkAddRandInt1000000(b *testing.B) {
	rand.Seed(time.Now().UTC().Unix())
	testSet := mapset.New[int]()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			testSet.Add(rand.Int())
		}
	}
}
