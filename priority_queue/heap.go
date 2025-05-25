package priority_queue

const (
	_INITIAL_SIZE        = 15
	_RESIZE_FACTOR       = 2
	_EMPTY_QUEUE_MESSAGE = "The queue is empty"
	_COMPARISON          = 0
)

// ===================== Types ======================
type priorityQueue[T any] struct {
	data []T
	size int
	cmp  func(T, T) int
}

// ================== Auxiliary Functions ===================

func leftChild(index int) int { return 2*index + 1 }

func rightChild(index int) int { return 2*index + 2 }

func parent(index int) int { return (index - 1) / 2 }

func swap[T any](data []T, x, y int) { data[x], data[y] = data[y], data[x] }

func upHeap[T any](data []T, cmp func(T, T) int, index int) {
	if index == 0 {
		return
	}
	parentIdx := parent(index)

	if parentIdx >= 0 && cmp(data[index], data[parentIdx]) > _COMPARISON {
		swap(data, index, parentIdx)
		upHeap(data, cmp, parentIdx)
	}
}

func downHeap[T any](data []T, cmp func(T, T) int, index int, size int) {
	left, right, largest := leftChild(index), rightChild(index), index

	if left < size && cmp(data[left], data[largest]) > 0 {
		largest = left
	}

	if right < size && cmp(data[right], data[largest]) > 0 {
		largest = right
	}

	if largest != index {
		swap(data, index, largest)
		downHeap(data, cmp, largest, size)
	}
}

func heapify[T any](data []T, cmpFunc func(T, T) int) {
	for i := len(data) / 2; i >= 0; i-- {
		downHeap(data, cmpFunc, i, len(data))
	}
}

func (heap *priorityQueue[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	copy(newData, heap.data[:heap.size])
	heap.data = newData
}

// =================== Heap Primitives ====================

func NewHeap[T any](cmpFunc func(T, T) int) PriorityQueue[T] {
	return NewHeapFromArray([]T{}, cmpFunc)
}

func NewHeapFromArray[T any](array []T, cmpFunc func(T, T) int) PriorityQueue[T] {
	newArr := make([]T, max(_INITIAL_SIZE, len(array)))
	copy(newArr, array)
	heapify(newArr, cmpFunc)

	heap := new(priorityQueue[T])
	heap.data = newArr
	heap.cmp = cmpFunc
	heap.size = len(array)
	return heap
}

func (heap *priorityQueue[T]) IsEmpty() bool {
	return heap.size == 0
}

func (heap *priorityQueue[T]) Enqueue(item T) {
	if heap.size == cap(heap.data) {
		heap.resize(heap.size * _RESIZE_FACTOR)
	}
	heap.data[heap.size] = item
	upHeap(heap.data, heap.cmp, heap.size)
	heap.size++
}

func (heap *priorityQueue[T]) PeekMax() T {
	if heap.IsEmpty() {
		panic(_EMPTY_QUEUE_MESSAGE)
	}
	return heap.data[0]
}

func (heap *priorityQueue[T]) Dequeue() T {
	if heap.IsEmpty() {
		panic(_EMPTY_QUEUE_MESSAGE)
	}
	maxElement := heap.data[0]
	heap.size--
	heap.data[0] = heap.data[heap.size]
	downHeap(heap.data, heap.cmp, 0, heap.size)

	if heap.size <= cap(heap.data)/(_RESIZE_FACTOR*_RESIZE_FACTOR) && cap(heap.data) > _INITIAL_SIZE {
		heap.resize(cap(heap.data) / _RESIZE_FACTOR)
	}
	return maxElement
}

func (heap *priorityQueue[T]) Size() int {
	return heap.size
}

// ======================= HeapSort ========================

func HeapSort[T any](elements []T, cmpFunc func(T, T) int) {
	heapify(elements, cmpFunc)

	for i := len(elements) - 1; i >= 0; i-- {
		swap(elements, 0, i)
		downHeap(elements, cmpFunc, 0, i)
	}
}
