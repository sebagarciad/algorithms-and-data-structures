package priority_queue

type PriorityQueue[T any] interface {

	// IsEmpty returns true if the queue is empty, false otherwise.
	IsEmpty() bool

	// Enqueue adds an element to the heap.
	Enqueue(T)

	// PeekMax returns the element with the highest priority. If it is empty, it panics with the message
	// "The queue is empty".
	PeekMax() T

	// Dequeue removes and returns the element with the highest priority. If it is empty, it panics with the message
	// "The queue is empty".
	Dequeue() T

	// Size returns the number of elements in the priority queue.
	Size() int
}
