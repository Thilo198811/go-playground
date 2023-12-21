package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/calculator"
)

type calculation struct {
	A         int    `json:"a"`
	B         int    `json:"b"`
	Operation string `json:"operation"`
	Result    int    `json:"result"`
}

var calculatorTrace = []calculation{
	{A: 1, B: 3, Operation: "*"},
	{A: 4, B: 3, Operation: "+"},
	{A: 10, B: 5, Operation: "-"},
}

func getTrace(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, calculatorTrace)
}

func getResults(c *gin.Context) {
	var calculatorResults = []calculation{}

	for _, c := range calculatorTrace {
		result, err := calculator.Calc(c.A, c.B, c.Operation)
		if err != nil {
			log.Fatal(err)
		}
		c.Result = result
		calculatorResults = append(calculatorResults, c)
	}
	c.IndentedJSON(http.StatusOK, calculatorResults)
}

func addCalculation(c *gin.Context) {
	var newCalc calculation
	c.BindJSON(&newCalc)
	calculatorTrace = append(calculatorTrace, newCalc)
	c.IndentedJSON(http.StatusOK, newCalc)
}

func main() {
	router := gin.Default()
	router.GET("/trace", getTrace)
	router.GET("/results", getResults)
	router.POST("/calc", addCalculation)
	router.Run("localhost:8080")
}
