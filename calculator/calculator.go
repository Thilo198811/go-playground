package calculator

import (
	"errors"
)

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
