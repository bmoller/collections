/*
Package collections provides basic data structures for a variety of purposes.

Implementations are not guaranteed to be stable, but all functionality is exposed via interfaces for a stable API.
Some interfaces have multiple implementations; consideration should be given to how the structure will be used when making a choice.
*/

package collections

import (
	"errors"
	"fmt"
)

// Collection

/*
A Collection represents an arbitrary group of values.
*/
type Collection[T comparable] interface {
	Empty() bool // Indicates if the Collection is empty
	Size() int   // Returns the number of values in the Collection
}

// Iterable

/*
An Iterable returns an Iterator to navigate its elements.
*/
type Iterable[T comparable] interface {
	Iterator() Iterator[T]
}

/*
An Iterator can be used to loop through all of the elements of the returning Iterable.
*/
type Iterator[T comparable] func() (T, error)

/*
ErrNoMoreItems indicates that an iterator has returned all of its items.
Callers should be prepared for this error and treat it as an expected state.
*/
var ErrNoMoreItems = errors.New("no more items to return")

// List

/*
A List is an ordered Collection of items.
Values added to the list retain the order in which they are added, and can be accessed by index.
New elements can be inserted at an arbitrary index or added to the end of the list.
Elements can be removed individually by index, or all at once with a call to Clear.

A List can also return a subset of its values via a call to SubList.
Similar to a slice, a SubList is created by referencing a range of indexes of the originating list.
However, a SubList is not another view into the same values, but instead is a complete copy of the elements in the range specified.
*/
type List[T comparable] interface {
	Collection[T]

	Add(T)
	Clear()
	Get(int) (T, error)
	Insert(int, T) error
	Remove(int) (T, error)
	SubList(int, int) (List[T], error)
}

/*
ErrEmptyList indicates that the list is empty and the method called is not available.
*/
var ErrEmptyList = errors.New("list is empty")

/*
ErrIndexOutOfRange is returned when a method is called with an index beyond the list's size.
*/
type ErrIndexOutOfRange struct {
	Index int // Index requested in the method call
	Size  int // Size of the list
}

func (e ErrIndexOutOfRange) Error() string {
	return fmt.Sprintf("index %d is invalid for list of length %d", e.Index, e.Size)
}

/*
ErrInvalidRange is returned if a range specifies an index less than 0, or if the End index is less than the Start index.
*/
type ErrInvalidRange struct {
	End   int // End index of the requested range
	Start int // Start index of the requested range
}

func (e ErrInvalidRange) Error() string {
	return fmt.Sprintf("invalid range with start %d and end %d; valid ranges follow 0 <= start <= end", e.Start, e.End)
}

// LinkedList

/*
A ListNode stores a single element of a LinkedList.
In addition to the element's value the node also stores references to the next and previous nodes in the list, allowing for traversal in either direction.
In general, a call to Next or Previous will return a pointer to a concrete implementation of the ListNode interface.
A LinkedList also stores references to several of its key nodes, such as the head and tail.
*/
type ListNode[T comparable] interface {
	Next() ListNode[T]
	Previous() ListNode[T]
	Value() T
}

/*
A LinkedList implements the List interface with an extra set of methods to increase versatility for traversal and element management.
The list is composed of a group of ListNodes that are members of the LinkedList.
References are always maintained to the first and last nodes in the list for rapid addition of new items.

Once a reference to any node is obtained, the list can be traversed to any other node by repeated calls to Next or Previous.
However, nodes do not necessarily know their positions in the larger list, though the order in which they are added is retained.

LinkedLists can add elements directly before or after an existing node.
Existing nodes can also be directly removed, in which case the previous and next nodes (if any) are linked to each other.

As a general rule, LinkedLists are slower than Lists for any index-based operations as the nodes must be traversed to reach the required element.
*/
type LinkedList[T comparable] interface {
	List[T]

	GetNode(int) (ListNode[T], error)
	Head() ListNode[T]
	InsertAfter(ListNode[T], T) (ListNode[T], error)
	InsertBefore(ListNode[T], T) (ListNode[T], error)
	RemoveNode(ListNode[T]) error
	Tail() ListNode[T]
}

/*
ErrNodeIsNotElement is returned if a method is called with a reference to a node that is not in the LinkedList.
*/
var ErrNodeIsNotElement = errors.New("node is not an element of this list")

/*
ErrWrongNodeType indicates that a referenced node is from an incompatible concrete implementation of LinkedList.
Most methods of nodes and lists reference interface types in their signatures, but rely on specific plumbing in their implementation.
*/
var ErrWrongNodeType = errors.New("node is from an incompatible list implementation")

// Queue

/*
A Queue represents a First In, First Out data structure.
New elements are added to the end of the queue and will be returned after all preceding items.
Values are returned and removed from the Queue via Pop, or a value can be retrieved without removal via Peek.
*/
type Queue[T comparable] interface {
	Collection[T]

	Peek() (T, error)
	Pop() (T, error)
	Push(T)
}

/*
ErrEmptyQueue is returned when Peek or Pop is called on a Queue with no elements.
This should be considered a part of normal operations and callers should expect to handle the error.
*/
var ErrEmptyQueue = errors.New("queue is empty")

// Set

/*
A Set guarantees that exactly one instance of a value is present when added.
If the value is already a member of the Set then Add is a no-op.
The Set can be queried for a specific value and can remove values.
The Remove method never returns an error; it only ensures that the element is not present in the Set.
An indeterminate value can also be removed and returned with a call to Pop.
*/
type Set[T comparable] interface {
	Collection[T]
	Iterable[T]

	Add(T)
	Contains(T) bool
	Pop() (T, error)
	Remove(T)
}

/*
ErrEmptySet is returned when Peek or Pop are called on an empty Set.
*/
var ErrEmptySet = errors.New("stack is empty")

// Stack

/*
A Stack is an ordered group of items from which the last added is the first returned.
The common analogy is a stack of clean plates at a buffet or cafeteria; when one is removed, another rises to take its place.
A new element is added to the top of the Stack (first for retrieval) with a call to Push.
Peek and Pop return the next value from the Stack, with Peek retaining the value on the Stack and Pop removing it.
*/
type Stack[T comparable] interface {
	Collection[T]

	Peek() (T, error)
	Pop() (T, error)
	Push(T)
}

/*
ErrEmptyStack is returned when Peek or Pop are called on an empty Stack.
*/
var ErrEmptyStack = errors.New("stack is empty")
