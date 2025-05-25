package operations

import (
	"errors"
	"math"
)

func Add(op1, op2 int64) (int64, error) {
	return op1 + op2, nil
}

func Subtract(op1, op2 int64) (int64, error) {
	return op1 - op2, nil
}

func Multiply(op1, op2 int64) (int64, error) {
	return op1 * op2, nil
}

func Divide(op1, op2 int64) (int64, error) {
	if op2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return op1 / op2, nil
}

func Power(op1, op2 int64) (int64, error) {
	if op2 < 0 {
		return 0, errors.New("power with negative exponent")
	}
	res := math.Pow(float64(op1), float64(op2))
	return int64(res), nil
}

func Logarithm(op1, op2 int64) (int64, error) {
	if op2 < 2 || op1 <= 0 {
		return 0, errors.New("invalid values for logarithmic operations")
	}
	res := math.Log(float64(op1)) / math.Log(float64(op2))
	return int64(res), nil
}

func SquareRoot(op int64) (int64, error) {
	if op < 0 {
		return 0, errors.New("cannot calculate the square root of a negative number")
	}
	res := math.Sqrt(float64(op))
	return int64(res), nil
}

func TernaryOperator(op1, op2, op3 int64) (int64, error) {
	if int64(op1) != 0 {
		return op2, nil
	} else {
		return op3, nil
	}
}
