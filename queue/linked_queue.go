package queue

type queueNode[T any] struct {
	data T
	next *queueNode[T]
}

func newQueueNode[T any](data T) *queueNode[T] {
	node := new(queueNode[T])
	node.data = data
	return node
}

type linkedQueue[T any] struct {
	first *queueNode[T]
	last  *queueNode[T]
}

func NewLinkedQueue[T any]() Queue[T] {
	return new(linkedQueue[T])
}

func (q *linkedQueue[T]) IsEmpty() bool {
	return q.first == nil
}

func (q *linkedQueue[T]) Peek() T {
	if q.IsEmpty() {
		panic("The queue is empty")
	}
	return q.first.data
}

func (q *linkedQueue[T]) Enqueue(data T) {
	node := newQueueNode(data)
	if q.first == nil {
		q.first = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}
}

func (q *linkedQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("The queue is empty")
	}
	element := q.first.data
	q.first = q.first.next
	if q.first == nil {
		q.last = nil
	}
	return element
}
