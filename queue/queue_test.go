package queue_test

import (
	"testing"

	ADTQueue "github.com/sebagarciad/algorithms-and-data-structures/queue"

	"github.com/stretchr/testify/require"
)

const (
	_INT1    int     = 34
	_INT2    int     = 89
	_INT3    int     = 100
	_INT_VOL int     = 10000
	_STR1    string  = "Hello"
	_STR2    string  = "how"
	_STR3    string  = "are"
	_STR4    string  = "you?"
	_FLOAT1  float64 = 9.21564
	_FLOAT2  float64 = 5.4568489655
	_FLOAT3  float64 = 12.4578
	_FLOAT4  float64 = 12457.532
)

func TestEmptyQueue(t *testing.T) {
	// Tests on integer queue
	queue := ADTQueue.NewLinkedQueue[int]()

	require.True(t, queue.IsEmpty(), "IsEmpty should return True for a newly created queue")

	// Test that Dequeue action on a newly created queue is invalid
	require.Panics(t, func() { queue.Dequeue() }, "Cannot Dequeue a newly created queue")

	// Test that Peek action on a newly created queue is invalid
	require.Panics(t, func() { queue.Peek() }, "A newly created queue cannot have a first element")

	queue.Enqueue(_INT1)
	require.False(t, queue.IsEmpty(), "After enqueuing an element, IsEmpty should return False")

	queue.Dequeue()
	require.True(t, queue.IsEmpty(), "After dequeuing all elements, IsEmpty should return True")

	// Tests on string queue
	stringQueue := ADTQueue.NewLinkedQueue[string]()
	require.True(t, stringQueue.IsEmpty(), "IsEmpty should return True for a newly created queue")

	require.Panics(t, func() { stringQueue.Dequeue() }, "Cannot Dequeue a newly created queue")

	require.Panics(t, func() { stringQueue.Peek() }, "A newly created queue does not have a first element")
}

func TestPeek(t *testing.T) {
	queue := ADTQueue.NewLinkedQueue[int]()

	require.Panics(t, func() { queue.Peek() }, "A newly created queue does not have a first element")

	queue.Enqueue(_INT1)
	queue.Enqueue(_INT2)
	queue.Enqueue(_INT3)
	require.EqualValues(t, _INT1, queue.Peek(), "The first should be 34")

	queue.Dequeue()
	require.EqualValues(t, _INT2, queue.Peek(), "The first should be 89")

	queue.Dequeue()
	require.EqualValues(t, _INT3, queue.Peek(), "The first should be 100")

	queue.Dequeue()
	require.Panics(t, func() { queue.Peek() }, "A dequeued queue does not have a first element")
}

func TestEnqueue(t *testing.T) {
	// Tests with integer queue
	queue := ADTQueue.NewLinkedQueue[int]()

	queue.Enqueue(_INT1)
	require.False(t, queue.IsEmpty(), "After enqueuing an element, IsEmpty should return False")
	queue.Dequeue()
	require.True(t, queue.IsEmpty())

	queue.Enqueue(_INT1)
	queue.Enqueue(_INT2)
	queue.Enqueue(_INT3)
	require.False(t, queue.IsEmpty())
	require.EqualValues(t, _INT1, queue.Peek(), "The first should be 34")

	queue.Dequeue()
	require.EqualValues(t, _INT2, queue.Peek(), "The first should be 89")

	queue.Dequeue()
	require.EqualValues(t, _INT3, queue.Peek(), "The first should be 100")
	queue.Dequeue()
	require.True(t, queue.IsEmpty())

	// Volume test
	for i := range _INT_VOL {
		queue.Enqueue(i)
		require.EqualValues(t, 0, queue.Peek(), "The first of the queue should be correct for each Enqueue call")
	}
	require.False(t, queue.IsEmpty())
	require.Equal(t, 0, queue.Peek(), "With many elements, the queue should work well and not take too long")

	for j := 0; j < _INT_VOL; j++ {
		queue.Dequeue()
	}
	require.True(t, queue.IsEmpty())

	// Tests with string queue
	stringQueue := ADTQueue.NewLinkedQueue[string]()

	var str1 string = "This queue"
	var str2 string = "is not empty"
	stringQueue.Enqueue(str1)
	stringQueue.Enqueue(str2)
	require.False(t, stringQueue.IsEmpty(), "The queue should be able to enqueue strings without problems")

	require.Equal(t, str1, stringQueue.Dequeue(), "Should return 'This queue'")
	require.Equal(t, str2, stringQueue.Dequeue(), "Should return 'is not empty'")
	require.True(t, stringQueue.IsEmpty(), "IsEmpty should return True after dequeuing all elements from a string queue")

	// Tests with queue of slices
	sliceQueue := ADTQueue.NewLinkedQueue[[]int]()

	var slice1 []int = []int{1, 2, 3}
	var slice2 []int = []int{4, 5, 6, 7}
	sliceQueue.Enqueue(slice1)
	sliceQueue.Enqueue(slice2)
	require.False(t, sliceQueue.IsEmpty(), "The queue should work well with slices")

	require.Equal(t, []int{1, 2, 3}, sliceQueue.Dequeue())
	require.Equal(t, []int{4, 5, 6, 7}, sliceQueue.Dequeue())
	require.True(t, sliceQueue.IsEmpty(), "IsEmpty should return True after Dequeue on a queue of slices")
}

func TestDequeue(t *testing.T) {
	// Tests with integer queue
	queue := ADTQueue.NewLinkedQueue[int]()

	queue.Enqueue(_INT1)
	require.False(t, queue.IsEmpty(), "IsEmpty should return False after enqueuing an element")

	element := queue.Dequeue()
	require.True(t, queue.IsEmpty(), "After dequeuing all elements, IsEmpty should return True")
	require.EqualValues(t, _INT1, element, "Should return 34")

	queue.Enqueue(_INT1)
	queue.Enqueue(_INT2)
	queue.Enqueue(_INT3)
	require.False(t, queue.IsEmpty())

	require.EqualValues(t, _INT1, queue.Dequeue(), "Should return 34")
	require.EqualValues(t, _INT2, queue.Dequeue(), "Should return 89")
	require.EqualValues(t, _INT3, queue.Dequeue(), "Should return 100")
	require.True(t, queue.IsEmpty(), "After dequeuing all elements, IsEmpty should return True")

	// Volume test
	for i := range _INT_VOL {
		queue.Enqueue(i)
	}
	require.False(t, queue.IsEmpty())
	for i := range _INT_VOL {
		require.EqualValues(t, i, queue.Peek(), "The queue should respond quickly and maintain the queue invariant when using many elements")
		first := queue.Dequeue()
		require.EqualValues(t, i, first)
	}
	require.True(t, queue.IsEmpty(), "Should return True when all elements are dequeued")

	// Test that Dequeue action on a queue with all elements dequeued is invalid
	require.Panics(t, func() { queue.Dequeue() }, "Cannot dequeue an empty queue")

	// Test that Peek action on a queue with all elements dequeued is invalid
	require.Panics(t, func() { queue.Peek() }, "A queue without elements does not have a first element")

	// Tests with string queue
	stringQueue := ADTQueue.NewLinkedQueue[string]()

	stringQueue.Enqueue(_STR1)
	stringQueue.Enqueue(_STR2)
	stringQueue.Enqueue(_STR3)
	stringQueue.Enqueue(_STR4)
	require.False(t, stringQueue.IsEmpty())

	require.Equal(t, _STR1, stringQueue.Dequeue(), "Should return the string 'Hello'")
	require.Equal(t, _STR2, stringQueue.Dequeue(), "Should return the string 'how'")
	require.Equal(t, _STR3, stringQueue.Dequeue(), "Should return the string 'are'")
	require.Equal(t, _STR4, stringQueue.Dequeue(), "Should return the string 'you?'")
	require.True(t, stringQueue.IsEmpty())

	// Tests with float queue
	floatQueue := ADTQueue.NewLinkedQueue[float64]()

	floatQueue.Enqueue(_FLOAT1)
	floatQueue.Enqueue(_FLOAT2)
	floatQueue.Enqueue(_FLOAT3)
	floatQueue.Enqueue(_FLOAT4)
	require.False(t, floatQueue.IsEmpty())

	require.Equal(t, _FLOAT1, floatQueue.Dequeue(), "Should return the float 9.21564")
	require.Equal(t, _FLOAT2, floatQueue.Dequeue(), "Should return the float 5.4568489655")
	require.Equal(t, _FLOAT3, floatQueue.Dequeue(), "Should return the float 12.4578")
	require.Equal(t, _FLOAT4, floatQueue.Dequeue(), "Should return the float 12457.532")
	require.True(t, floatQueue.IsEmpty(), "After dequeuing all elements, IsEmpty should return True")
}
