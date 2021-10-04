package main

import (
	"fmt"
	"strings"
	"src/math"
)

/*func Sum(a int, b int) int {
	return a+b
}*/

func main() {
	var firstNumber int
	var secondNumber int
	var name string

	fmt.Printf("Type your name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	fmt.Printf("Type two numbers: ")
	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)

	result := math.Sum(firstNumber, secondNumber)
	
	fmt.Println(result)
	fmt.Printf("Your name is %v\n", name)

	fmt.Printf("Type of \"result\" variable is %T\n", result)

	fmt.Printf("The sum is %v\n", result)

	fmt.Printf("Type of \"result\" variable is %T\n", result)
}