package queue

type Queue[T any] interface {

	// IsEmpty returns true if the queue has no enqueued elements, false otherwise.
	IsEmpty() bool

	// Peek returns the value at the front of the queue. If it is empty, it panics with the message
	// "The queue is empty".
	Peek() T

	// Enqueue adds a new element to the end of the queue.
	Enqueue(T)

	// Dequeue removes the first element from the queue. If the queue has elements, it removes the first one
	// and returns its value. If it is empty, it panics with the message "The queue is empty".
	Dequeue() T
}
