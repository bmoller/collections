package collections

type LinkedList[T comparable] interface {
	Append(T)
	Head() *ListNode[T]
	Size() int
}

type ListNode[T comparable] interface {
	Next() *ListNode[T]
	Previous() *ListNode[T]
	Value() T
}
