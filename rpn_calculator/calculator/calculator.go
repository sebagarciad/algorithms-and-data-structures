package calculator

import (
	ADTStack "data_structures/stack"
)

type Calculator interface {
	BinaryOperations(ADTStack.Stack[int64], string) (int64, error)

	UnaryOperations(ADTStack.Stack[int64], string) (int64, error)

	TernaryOperations(ADTStack.Stack[int64], string) (int64, error)
}
