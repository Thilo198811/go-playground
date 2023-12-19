package main

import (
	"fmt"
	"log"

	"example.com/calculator"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(calculator.Add(2, 2))
	fmt.Println(quote.Go())

	res, err := calculator.Div(2, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
