package stack

const _INITIAL_SIZE int = 2

/* Definition of the stack struct provided by the course. */
type dynamicStack[T any] struct {
	data  []T
	count int
}

// Creates and returns a dynamic stack of the Stack interface of type any, with an initial capacity of 2 elements.
func NewStack[T any]() Stack[T] {
	stack := new(dynamicStack[T])
	stack.data = make([]T, _INITIAL_SIZE)
	return stack
}

func (stack *dynamicStack[T]) IsEmpty() bool {
	return stack.count == 0
}

func (stack *dynamicStack[T]) Peek() T {
	if stack.IsEmpty() {
		panic("The stack is empty")
	}
	return stack.data[stack.count-1]
}

// Pushes a new element. If the number of elements is equal to the stack's capacity, it resizes
// to double the current size.
func (stack *dynamicStack[T]) Push(element T) {
	if stack.count == cap(stack.data) {
		stack.resize(2 * cap(stack.data))
	}
	stack.data[stack.count] = element
	stack.count++
}

// Pops an element and returns it. If the number of elements is equal to or less than a quarter of the capacity
// the stack is resized to half the current size.
func (stack *dynamicStack[T]) Pop() T {
	if stack.IsEmpty() {
		panic("The stack is empty")
	}
	top := stack.data[stack.count-1]
	stack.count--
	if stack.count <= cap(stack.data)/4 {
		stack.resize(cap(stack.data) / 2)
	}
	return top
}

// Creates a new slice and copies the elements from the previous slice to it. If the new capacity falls below the
// initial capacity, the initial capacity is restored.
func (stack *dynamicStack[T]) resize(newCap int) {
	if newCap < _INITIAL_SIZE {
		newCap = _INITIAL_SIZE
	}
	newData := make([]T, newCap)
	copy(newData, stack.data)
	stack.data = newData
}
