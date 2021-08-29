package main

import (
	"fmt"
)

func Sum(a int, b int) int {
	return a+b
}

func main() {
	var firstNumber int
	var secondNumber int

	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)

	result := Sum(firstNumber, secondNumber)
	
	fmt.Println(result)

	fmt.Printf("Type of \"result\" variable is %T\n", result)
}