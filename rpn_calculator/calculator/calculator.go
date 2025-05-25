package calculator

// ADTStack "github.com/sebagarciad/algorithms-and-data-structures/data_structures/stack"

type Calculator interface {
	BinaryOperations(ADTStack.Pila[int64], string) (int64, error)

	UnaryOperations(ADTStack.Pila[int64], string) (int64, error)

	TernaryOperations(ADTStack.Pila[int64], string) (int64, error)
}
