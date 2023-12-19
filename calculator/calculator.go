package calculator

func Add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	if(b == 0) {
		return 0
	} else {
		return a / b
	}
}





