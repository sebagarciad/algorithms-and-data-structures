package calculator

import (
	"errors"

	TDAPila "github.com/sebagarciad/algorithms-and-data-structures/data_structures/stack"
	Op "github.com/sebagarciad/algorithms-and-data-structures/rpn_calculator/operations"
)

const (
	_UNARY_OPS   = 1
	_BINARY_OPS  = 2
	_TERNARY_OPS = 3
)

type Calculations struct{}

// validatePop checks for input errors to ensure the operation is valid.
// It receives the operand stack and the required number of operands for each operation type,
// and returns a slice with the popped operands, or an error if operands are missing.
func validatePop(stack TDAPila.Pila[int64], numOperands int) ([]int64, error) {
	operands := make([]int64, numOperands)
	for i := numOperands - 1; i > -1; i-- {
		if stack.EstaVacia() {
			return nil, errors.New("missing operands")
		}
		operands[i] = stack.Desapilar()
	}
	return operands, nil
}

func (calc *Calculations) UnaryOperations(stack TDAPila.Pila[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _UNARY_OPS)
	if err != nil {
		return 0, err
	}
	return Op.RaizCuadrada(operands[0])
}

func (calc *Calculations) BinaryOperations(stack TDAPila.Pila[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _BINARY_OPS)
	if err != nil {
		return 0, err
	}
	op1 := operands[0]
	op2 := operands[1]

	switch operator {
	case "+":
		return Op.Suma(op1, op2)

	case "-":
		return Op.Resta(op1, op2)

	case "*":
		return Op.Multiplicacion(op1, op2)

	case "/":
		return Op.Division(op1, op2)

	case "^":
		return Op.Potencia(op1, op2)

	case "log":
		return Op.Logaritmo(op1, op2)

	default:
		return 0, errors.New("invalid operator")
	}
}

func (calc *Calculations) TernaryOperations(stack TDAPila.Pila[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _TERNARY_OPS)
	if err != nil {
		return 0, err
	}
	op1 := operands[0]
	op2 := operands[1]
	op3 := operands[2]

	return Op.OperadorTernario(op1, op2, op3)
}
