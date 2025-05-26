package calculator

import (
	"errors"

	Op "rpn_calculator/operations"

	ADTStack "data_structures/stack"
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
func validatePop(stack ADTStack.Stack[int64], numOperands int) ([]int64, error) {
	operands := make([]int64, numOperands)
	for i := numOperands - 1; i > -1; i-- {
		if stack.IsEmpty() {
			return nil, errors.New("missing operands")
		}
		operands[i] = stack.Pop()
	}
	return operands, nil
}

func (calc *Calculations) UnaryOperations(stack ADTStack.Stack[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _UNARY_OPS)
	if err != nil {
		return 0, err
	}
	return Op.SquareRoot(operands[0])
}

func (calc *Calculations) BinaryOperations(stack ADTStack.Stack[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _BINARY_OPS)
	if err != nil {
		return 0, err
	}
	op1 := operands[0]
	op2 := operands[1]

	switch operator {
	case "+":
		return Op.Add(op1, op2)

	case "-":
		return Op.Subtract(op1, op2)

	case "*":
		return Op.Multiply(op1, op2)

	case "/":
		return Op.Divide(op1, op2)

	case "^":
		return Op.Power(op1, op2)

	case "log":
		return Op.Logarithm(op1, op2)

	default:
		return 0, errors.New("invalid operator")
	}
}

func (calc *Calculations) TernaryOperations(stack ADTStack.Stack[int64], operator string) (int64, error) {
	operands, err := validatePop(stack, _TERNARY_OPS)
	if err != nil {
		return 0, err
	}
	op1 := operands[0]
	op2 := operands[1]
	op3 := operands[2]

	return Op.TernaryOperator(op1, op2, op3)
}
