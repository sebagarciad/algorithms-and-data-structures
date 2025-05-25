package priority_queue_test

import (
	"math/rand"
	"strings"
	"testing"

	TDAHeap "github.com/sebagarciad/algorithms-and-data-structures/priority_queue"

	"github.com/stretchr/testify/require"
)

func cmpInt(a, b int) int {
	return a - b
}

func cmpStr(a, b string) int {
	return strings.Compare(a, b)
}

func TestEmptyHeap(t *testing.T) {
	heap := TDAHeap.NewHeap[int](cmpInt)
	require.True(t, heap.IsEmpty(), "Checks that the heap is empty when created")
	require.Panics(t, func() { heap.PeekMax() }, "Panics when trying to see the max of an empty heap")
	require.Panics(t, func() { heap.Dequeue() }, "Panics when trying to dequeue an element from an empty heap")
	require.Equal(t, 0, heap.Size(), "A newly created heap should have count 0")
}

func TestHeapOneElement(t *testing.T) {
	heap := TDAHeap.NewHeap[int](cmpInt)

	heap.Enqueue(5)
	require.False(t, heap.IsEmpty(), "The heap cannot be empty after enqueuing an element")
	require.Equal(t, 1, heap.Size(), "Should return 1")
	require.Equal(t, 5, heap.PeekMax(), "The max should be 5")

	heap.Dequeue()
	require.True(t, heap.IsEmpty(), "Should return true")
	require.Equal(t, 0, heap.Size(), "After dequeuing the only element, count should be 0")
	require.Panics(t, func() { heap.PeekMax() }, "Should panic when using PeekMax on an empty heap")
}

func TestHeapEnqueueAndDequeueMultipleElements(t *testing.T) {
	heap := TDAHeap.NewHeap[int](cmpInt)

	heap.Enqueue(10)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 1, heap.Size(), "Heap count should be 1")
	require.Equal(t, 10, heap.PeekMax(), "Max should be 10")

	heap.Enqueue(5)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 2, heap.Size(), "Heap count should be 2")
	require.Equal(t, 10, heap.PeekMax(), "Max should be 10")

	heap.Enqueue(14)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 3, heap.Size(), "Heap count should be 3")
	require.Equal(t, 14, heap.PeekMax(), "Max should be 14")

	heap.Enqueue(12)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 4, heap.Size(), "Heap count should be 4")
	require.Equal(t, 14, heap.PeekMax(), "Max should be 14")

	heap.Enqueue(20)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 5, heap.Size(), "Heap count should be 5")
	require.Equal(t, 20, heap.PeekMax(), "Max should be 20")

	heap.Enqueue(7)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 6, heap.Size(), "Heap count should be 6")
	require.Equal(t, 20, heap.PeekMax(), "Max should be 20")

	heap.Enqueue(1)

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 7, heap.Size(), "Heap count should be 7")
	require.Equal(t, 20, heap.PeekMax(), "Max should be 20")

	require.Equal(t, 20, heap.Dequeue(), "Should return 20")
	require.Equal(t, 14, heap.PeekMax(), "New max should be 14")
	require.Equal(t, 6, heap.Size(), "After dequeuing, count should be 6")

	require.Equal(t, 14, heap.Dequeue(), "Dequeued value should be 14")
	require.Equal(t, 12, heap.PeekMax(), "New max should be 12")
	require.Equal(t, 5, heap.Size(), "Count should be 5")

	require.Equal(t, 12, heap.Dequeue(), "Dequeued value should be 12")
	require.Equal(t, 10, heap.PeekMax(), "New max should be 10")
	require.Equal(t, 4, heap.Size(), "Count should be 4")

	require.Equal(t, 10, heap.Dequeue(), "Dequeued value should be 10")
	require.Equal(t, 7, heap.PeekMax(), "New max should be 7")
	require.Equal(t, 3, heap.Size(), "Count should be 3")

	require.Equal(t, 7, heap.Dequeue(), "Dequeued value should be 7")
	require.Equal(t, 5, heap.PeekMax(), "New max should be 5")
	require.Equal(t, 2, heap.Size(), "Count should be 2")

	require.Equal(t, 5, heap.Dequeue(), "Dequeued value should be 5")
	require.Equal(t, 1, heap.PeekMax(), "New max should be 1")
	require.Equal(t, 1, heap.Size(), "Count should be 1")

	require.Equal(t, 1, heap.Dequeue(), "Dequeued value should be 1")
	require.Equal(t, 0, heap.Size(), "Count should be 0")
	require.True(t, heap.IsEmpty(), "Should return true after dequeuing all elements")
	require.Panics(t, func() { heap.Dequeue() }, "Should panic: no more elements to dequeue")
}

func TestHeapWithStrings(t *testing.T) {
	heap := TDAHeap.NewHeap[string](cmpStr)

	heap.Enqueue("Algorithms")

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 1, heap.Size(), "Heap count should be 1")
	require.Equal(t, "Algorithms", heap.PeekMax(), "Max should be 'Algorithms'")

	heap.Enqueue("Y")

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 2, heap.Size(), "Heap count should be 2")
	require.Equal(t, "Y", heap.PeekMax(), "Max should be 'Y'")

	heap.Enqueue("Structures")

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 3, heap.Size(), "Heap count should be 3")
	require.Equal(t, "Y", heap.PeekMax(), "Max should be 'Y'")

	heap.Enqueue("Of")

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 4, heap.Size(), "Heap count should be 4")
	require.Equal(t, "Y", heap.PeekMax(), "Max should be 'Y'")

	heap.Enqueue("Data")

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 5, heap.Size(), "Heap count should be 5")
	require.Equal(t, "Y", heap.PeekMax(), "Max should be 'Y'")

	require.Equal(t, "Y", heap.Dequeue(), "Should return 'Y'")
	require.Equal(t, "Structures", heap.PeekMax(), "New max should be 'Structures'")
	require.Equal(t, 4, heap.Size(), "After dequeuing, count should be 4")

	require.Equal(t, "Structures", heap.Dequeue(), "Dequeued value should be 'Structures'")
	require.Equal(t, "Of", heap.PeekMax(), "New max should be 'Of'")
	require.Equal(t, 3, heap.Size(), "Count should be 3")

	require.Equal(t, "Of", heap.Dequeue(), "Dequeued value should be 'Of'")
	require.Equal(t, "Data", heap.PeekMax(), "New max should be 'Data'")
	require.Equal(t, 2, heap.Size(), "Count should be 2")

	require.Equal(t, "Data", heap.Dequeue(), "Dequeued value should be 'Data'")
	require.Equal(t, 1, heap.Size(), "Count should be 1")

	require.Equal(t, "Algorithms", heap.Dequeue(), "Dequeued value should be 'Algorithms'")
	require.Equal(t, 0, heap.Size(), "Count should be 0")
	require.True(t, heap.IsEmpty(), "Should return true after dequeuing all elements")
	require.Panics(t, func() { heap.Dequeue() }, "Should panic: no more elements to dequeue")
}

func TestHeapVolumeWithoutArray(t *testing.T) {
	heap := TDAHeap.NewHeap[int](cmpInt)

	for i := 0; i < 100000; i++ {
		heap.Enqueue(i)
		require.False(t, heap.IsEmpty(), "Should return false")
		require.Equal(t, i+1, heap.Size(), "Heap count should be '%d'", i+1)
		require.Equal(t, i, heap.PeekMax(), "Heap max should be '%d'", i)
	}

	require.False(t, heap.IsEmpty(), "Should return false")
	require.Equal(t, 100000, heap.Size(), "Heap count should be 100000")

	for i := 99999; i >= 0; i-- {
		require.Equal(t, i, heap.PeekMax(), "Heap max should be '%d'", i)
		require.Equal(t, i, heap.Dequeue(), "Heap should respond quickly to a volume test")
		require.Equal(t, i, heap.Size(), "Heap count should be '%d'", i)
	}

	require.True(t, heap.IsEmpty(), "Should return true")
	require.Equal(t, 0, heap.Size(), "Heap count should be 100000")
}

func TestHeapArrEmpty(t *testing.T) {
	heap := TDAHeap.NewHeapFromArray[int](nil, cmpInt)
	require.True(t, heap.IsEmpty(), "Checks that the heap created from an empty array is empty")
	require.Equal(t, 0, heap.Size(), "Heap count should be 0")
	require.Panics(t, func() { heap.PeekMax() }, "Should panic when trying to see the max of an empty heap")
}

func TestHeapArrOneElement(t *testing.T) {
	heap := TDAHeap.NewHeapFromArray([]int{42}, cmpInt)
	require.False(t, heap.IsEmpty(), "Heap cannot be empty after creating with one element")
	require.Equal(t, 1, heap.Size(), "Heap count should be 1")
	require.Equal(t, 42, heap.PeekMax(), "Max should be 42")
}

func TestHeapArrMultipleElements(t *testing.T) {
	elements := []int{5, 3, 8, 1, 2, 7}
	heap := TDAHeap.NewHeapFromArray(elements, cmpInt)

	require.False(t, heap.IsEmpty(), "Heap cannot be empty after creating with multiple elements")
	require.Equal(t, 6, heap.Size(), "Heap count should be 6")
	require.Equal(t, 8, heap.PeekMax(), "Max should be 8")
}

func TestHeapArrWithOperations(t *testing.T) {
	elements := []int{5, 3, 8, 1, 2, 7}
	heap := TDAHeap.NewHeapFromArray(elements, cmpInt)

	require.False(t, heap.IsEmpty(), "Heap cannot be empty after creating with multiple elements")
	require.Equal(t, 6, heap.Size(), "Heap count should be 6")
	require.Equal(t, 8, heap.PeekMax(), "Max should be 8")

	require.Equal(t, 8, heap.Dequeue(), "Dequeued element should be 8")
	require.Equal(t, 7, heap.PeekMax(), "New max should be 7")
	require.Equal(t, 5, heap.Size(), "Heap count should be 5")

	heap.Enqueue(10)
	require.Equal(t, 6, heap.Size(), "Heap count should be 6 after enqueuing a new element")
	require.Equal(t, 10, heap.PeekMax(), "New max should be 10")

	expected := []int{10, 7, 5, 3, 2, 1}
	for _, element := range expected {
		require.Equal(t, element, heap.Dequeue(), "Heap should respond in the correct order when dequeuing")
	}

	require.True(t, heap.IsEmpty(), "Should return true")
	require.Equal(t, 0, heap.Size(), "Heap count should be 0")
}

func TestHeapArrVolume(t *testing.T) {
	elements := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		elements[i] = i
	}

	heap := TDAHeap.NewHeapFromArray(elements, cmpInt)

	require.False(t, heap.IsEmpty(), "Heap cannot be empty after creating with multiple elements")
	require.Equal(t, 100000, heap.Size(), "Heap count should be 100000")

	for i := 99999; i >= 0; i-- {
		require.Equal(t, i, heap.PeekMax(), "Heap max should be '%d'", i)
		require.Equal(t, i, heap.Dequeue(), "Heap should respond quickly to a volume test")
		require.Equal(t, i, heap.Size(), "Heap count should be '%d'", i)
	}

	require.True(t, heap.IsEmpty(), "Should return true")
	require.Equal(t, 0, heap.Size(), "Heap count should be 0")
}

func TestHeapSortEmptyArray(t *testing.T) {
	elements := []int{}
	expected := []int{}

	TDAHeap.HeapSort(elements, cmpInt)
	require.Equal(t, expected, elements, "HeapSort does not break with an empty array")
}

func TestHeapSortOneElementArray(t *testing.T) {
	elements := []int{4}
	expected := []int{4}

	TDAHeap.HeapSort(elements, cmpInt)
	require.Equal(t, expected, elements, "HeapSort works correctly with a one-element array")
}

func TestHeapSortMultipleElementsArray(t *testing.T) {
	elements := []int{4, 10, 3, 5, 1}
	expected := []int{1, 3, 4, 5, 10}

	TDAHeap.HeapSort(elements, cmpInt)
	require.Equal(t, expected, elements, "The sorted array should be [1, 3, 4, 5, 10]")
}

func TestHeapSortStrings(t *testing.T) {
	elements := []string{"algorithms", "and", "structures", "of", "data"}
	expected := []string{"algorithms", "data", "of", "structures", "and"}

	TDAHeap.HeapSort(elements, cmpStr)
	require.Equal(t, expected, elements, "The sorted array should be ['algorithms', 'data', 'of', 'structures', 'and']")
}

func TestHeapSortVolume(t *testing.T) {
	elements := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		elements[i] = i
	}

	rand.Shuffle(len(elements), func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})

	expected := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		expected[i] = i
	}

	TDAHeap.HeapSort(elements, cmpInt)
	require.Equal(t, expected, elements, "HeapSort works correctly with a large amount of unordered elements")
}
