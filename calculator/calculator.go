package calculator

import (
	"errors"
)

func Calc(a int, b int, operation string) (int, error) {
	if operation == "+" {
		return Add(a, b), nil
	} else if operation == "-" {
		return Sub(a, b), nil
	} else if operation == "*" {
		return Mul(a, b), nil
	} else if operation == "/" {
		return Div(a, b)
	}
	return 0, errors.New("Unsupported Operation")
}

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by Zero")
	} else {
		return a / b, nil
	}
}
