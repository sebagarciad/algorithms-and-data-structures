package stack_test

import (
	"testing"

	ADTStack "github.com/sebagarciad/algorithms-and-data-structures/stack"

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

func TestEmptyStack(t *testing.T) {
	// Tests on stack of integers
	stack := ADTStack.NewStack[int]()

	require.True(t, stack.IsEmpty(), "IsEmpty should return True for a newly created stack")

	// Test that Pop action on a newly created stack is invalid
	require.Panics(t, func() { stack.Pop() }, "Cannot pop a newly created stack")

	// Test that Peek action on a newly created stack is invalid
	require.Panics(t, func() { stack.Peek() }, "A newly created stack cannot have a top")

	stack.Push(_INT1)
	require.False(t, stack.IsEmpty(), "After pushing an element, IsEmpty should return False")

	stack.Pop()
	require.True(t, stack.IsEmpty(), "After popping all elements, IsEmpty should return True")

	// Tests on stack of strings
	stackString := ADTStack.NewStack[string]()
	require.True(t, stackString.IsEmpty(), "IsEmpty should return True for a newly created stack")

	require.Panics(t, func() { stackString.Pop() }, "Cannot pop a newly created stack")

	require.Panics(t, func() { stackString.Peek() }, "A newly created stack has no top")
}

func TestPeek(t *testing.T) {
	stack := ADTStack.NewStack[int]()

	stack.Push(_INT1)
	stack.Push(_INT2)
	stack.Push(_INT3)
	require.EqualValues(t, _INT3, stack.Peek(), "Top should be 100")

	stack.Pop()
	require.EqualValues(t, _INT2, stack.Peek(), "Top should be 89")

	stack.Pop()
	require.EqualValues(t, _INT1, stack.Peek(), "Top should be 34")

	stack.Pop()
	require.Panics(t, func() { stack.Peek() }, "Cannot peek the top of an empty stack")
}

func TestPush(t *testing.T) {
	// Tests with stack of integers
	stack := ADTStack.NewStack[int]()

	stack.Push(_INT1)
	require.False(t, stack.IsEmpty(), "After pushing an element, IsEmpty should return False")
	stack.Pop()

	stack.Push(_INT1)
	stack.Push(_INT2)
	stack.Push(_INT3)
	require.False(t, stack.IsEmpty())
	require.EqualValues(t, _INT3, stack.Peek(), "Top should be 100")

	stack.Pop()
	require.EqualValues(t, _INT2, stack.Peek(), "Top should be 89")

	stack.Pop()
	require.EqualValues(t, _INT1, stack.Peek(), "Top should be 34")
	stack.Pop()
	require.True(t, stack.IsEmpty())

	// Volume test
	for i := range _INT_VOL {
		stack.Push(i)
		require.EqualValues(t, i, stack.Peek(), "The stack top should be correct for each Push call")
	}
	require.False(t, stack.IsEmpty())
	require.Equal(t, _INT_VOL-1, stack.Peek(), "With many elements, the stack should work well and not take too long")

	for j := 0; j < _INT_VOL; j++ {
		stack.Pop()
	}
	require.True(t, stack.IsEmpty())

	for i := range _INT1 {
		stack.Push(i)
	}
	for i := _INT1; i < _INT_VOL; i++ {
		stack.Push(i)
	}
	require.False(t, stack.IsEmpty())
	require.Equal(t, _INT_VOL-1, stack.Peek(), "The stack should work with few or many elements")

	for j := 0; j < _INT_VOL; j++ {
		stack.Pop()
	}
	require.True(t, stack.IsEmpty())

	// Tests with stack of strings
	stackString := ADTStack.NewStack[string]()

	stackString.Push("This stack")
	stackString.Push("is not empty")
	require.False(t, stackString.IsEmpty(), "The stack should be able to push strings without problems")

	require.Equal(t, "is not empty", stackString.Pop(), "Should return 'is not empty'")
	require.Equal(t, "This stack", stackString.Pop(), "Should return 'This stack'")
	require.True(t, stackString.IsEmpty(), "IsEmpty should return True after popping a stack of strings")

	// Tests with stack of slices
	stackSlice := ADTStack.NewStack[[]int]()

	var slice1 []int = []int{1, 2, 3}
	var slice2 []int = []int{4, 5, 6, 7}
	stackSlice.Push(slice1)
	stackSlice.Push(slice2)
	require.False(t, stackSlice.IsEmpty(), "The stack should work well with slices")

	require.Equal(t, []int{4, 5, 6, 7}, stackSlice.Pop())
	require.Equal(t, []int{1, 2, 3}, stackSlice.Pop())
	require.True(t, stackSlice.IsEmpty(), "IsEmpty should return True after popping a stack of slices")
}

func TestPop(t *testing.T) {
	// Tests with stack of integers
	stack := ADTStack.NewStack[int]()

	stack.Push(_INT1)
	require.False(t, stack.IsEmpty())

	element := stack.Pop()
	require.True(t, stack.IsEmpty(), "After popping all elements, IsEmpty should return True")
	require.EqualValues(t, _INT1, element, "Should return 34")

	stack.Push(_INT1)
	stack.Push(_INT2)
	stack.Push(_INT3)
	require.False(t, stack.IsEmpty())

	require.EqualValues(t, _INT3, stack.Pop(), "Should return 100")
	require.EqualValues(t, _INT2, stack.Pop(), "Should return 89")
	require.EqualValues(t, _INT1, stack.Pop(), "Should return 34")
	require.True(t, stack.IsEmpty())

	// Volume test
	for i := range _INT_VOL {
		stack.Push(i)
	}
	require.False(t, stack.IsEmpty())
	for i := _INT_VOL - 1; i > -1; i-- {
		top := stack.Pop()
		require.EqualValues(t, i, top, "The stack should respond quickly and maintain the stack invariant when using many elements")
	}
	require.True(t, stack.IsEmpty(), "Should return True when all elements are popped")

	for i := range _INT1 {
		stack.Push(i)
	}
	for i := _INT1; i < _INT_VOL; i++ {
		stack.Push(i)
	}
	require.False(t, stack.IsEmpty())
	for i := _INT_VOL - 1; i > -1; i-- {
		top := stack.Pop()
		require.EqualValues(t, i, top, "Should resize well and maintain the stack invariant")
	}
	require.True(t, stack.IsEmpty())

	// Test that Pop action on an empty stack is invalid
	require.Panics(t, func() { stack.Pop() }, "Cannot pop an empty stack")

	// Test that Peek action on an empty stack is invalid
	require.Panics(t, func() { stack.Peek() }, "An empty stack has no top")

	// Tests with stack of strings
	stackStrings := ADTStack.NewStack[string]()

	stackStrings.Push(_STR1)
	stackStrings.Push(_STR2)
	stackStrings.Push(_STR3)
	stackStrings.Push(_STR4)
	require.False(t, stackStrings.IsEmpty())

	require.Equal(t, _STR4, stackStrings.Pop(), "Should return the string 'you?'")
	require.Equal(t, _STR3, stackStrings.Pop(), "Should return the string 'are'")
	require.Equal(t, _STR2, stackStrings.Pop(), "Should return the string 'how'")
	require.Equal(t, _STR1, stackStrings.Pop(), "Should return the string 'Hello'")
	require.True(t, stackStrings.IsEmpty())

	// Tests with stack of floats
	stackFloat := ADTStack.NewStack[float64]()

	stackFloat.Push(_FLOAT1)
	stackFloat.Push(_FLOAT2)
	stackFloat.Push(_FLOAT3)
	stackFloat.Push(_FLOAT4)
	require.False(t, stackFloat.IsEmpty())

	require.Equal(t, _FLOAT4, stackFloat.Pop(), "Should return the float 12457.532")
	require.Equal(t, _FLOAT3, stackFloat.Pop(), "Should return the float 12.4578")
	require.Equal(t, _FLOAT2, stackFloat.Pop(), "Should return the float 5.4568489655")
	require.Equal(t, _FLOAT1, stackFloat.Pop(), "Should return the float 9.21564")
	require.True(t, stackFloat.IsEmpty(), "After popping all elements, IsEmpty should return True")
}
