package linked_list

type List[T any] interface {
	// IsEmpty returns true if the list has no elements, false otherwise.
	IsEmpty() bool

	// InsertFirst inserts a new element at the beginning of the list.
	InsertFirst(T)

	// InsertLast inserts a new element at the end of the list.
	InsertLast(T)

	// DeleteFirst removes the first element from the list.
	// If the list has elements, it removes the first one and returns its value.
	// If it's empty, it panics with the message "The list is empty".
	DeleteFirst() T

	// SeeFirst gets the value of the first element in the list.
	// If it's empty, it panics with the message "The list is empty".
	SeeFirst() T

	// SeeLast gets the value of the last element in the list.
	// If it's empty, it panics with the message "The list is empty".
	SeeLast() T

	// Length returns the number of elements in the list.
	Length() int

	// Iterate traverses the elements of the list and executes the "visit" function on each of them.
	// If "visit" returns false, the iteration stops.
	Iterate(visit func(T) bool)

	// Iterator returns an iterator that allows traversing the list.
	Iterator() ListIterator[T]
}

type ListIterator[T any] interface {
	// SeeCurrent returns the value of the current element of the iterator.
	// If the iterator has finished traversing the list, it panics with the message "The iterator has finished iterating".
	SeeCurrent() T

	// HasNext returns true if the iterator has a next element, false otherwise.
	HasNext() bool

	// Next advances the iterator to the next element.
	// If the iterator has finished traversing the list, it panics with the message "The iterator has finished iterating".
	Next()

	// Insert inserts a new element at the current position of the iterator.
	// The new element is placed before the current element.
	// The iterator moves to the position of the new element.
	Insert(T)

	// Delete removes the current element and returns its value.
	// After deleting, the iterator moves to the next element.
	// If the iterator has finished traversing the list, it panics with the message "The iterator has finished iterating".
	Delete() T
}
