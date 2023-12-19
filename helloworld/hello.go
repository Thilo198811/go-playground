package main

import (
	"fmt"
	"rsc.io/quote"
	"example.com/calculator"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(calculator.Add(2,2))
	fmt.Println(quote.Go())
}
