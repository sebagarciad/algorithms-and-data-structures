package calculate

import (
	"errors"
	"strconv"

	ADTStack "github.com/sebagarciad/algorithms-and-data-structures/data_structures/stack"
	Calc "github.com/sebagarciad/algorithms-and-data-structures/rpn_calculator/calculator"
)

func Operate(elements []string) (int64, error) {
	stack := ADTStack.NewStack[int64]()
	calc := Calc.Calculations{}

	for _, element := range elements {
		if value, err := strconv.ParseInt(element, 10, 64); err == nil {
			stack.Push(int64(value))
		} else {
			switch element {
			case "sqrt":
				if res, err := calc.UnaryOperations(stack, element); err == nil {
					stack.Push(res)
				} else {
					return 0, err
				}

			case "+", "-", "*", "/", "^", "log":
				if res, err := calc.BinaryOperations(stack, element); err == nil {
					stack.Push(res)
				} else {
					return 0, err
				}

			case "?":
				if res, err := calc.TernaryOperations(stack, element); err == nil {
					stack.Push(res)
				} else {
					return 0, err
				}

			default:
				return 0, errors.New("invalid operator")
			}
		}
	}

	result := stack.Pop()
	if !stack.IsEmpty() {
		return 0, errors.New("extra operands")
	}

	return result, nil
}
