package stack

type Stack[T any] interface {

	// IsEmpty returns true if the stack has no elements, false otherwise.
	IsEmpty() bool

	// Peek returns the value at the top of the stack. If the stack has elements, it returns the top value.
	// If it is empty, it panics with the message "The stack is empty".
	Peek() T

	// Push adds a new element to the stack.
	Push(T)

	// Pop removes the top element from the stack. If the stack has elements, it removes and returns the top value.
	// If it is empty, it panics with the message "The stack is empty".
	Pop() T
}
