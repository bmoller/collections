package collections

type Iterator[T any] interface {
	/*
	   HasNext indicates if the Iterator has at least one additional element to turn with Next.
	*/
	HasNext() bool
	/*
	   Next advances the Iterator by once and returns the element at that position.
	*/
	Next() T
}

type ListIterator[T any] interface {
	Iterator[T]

	HasPrevious() bool
	NextIndex() int
	Previous() T
	PreviousIndex() int
}
