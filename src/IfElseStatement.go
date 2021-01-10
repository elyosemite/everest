package main

import "fmt"

func main() {
	var myAge int = 22;
	var year int = 2021

	if(year == 2021) {
		fmt.Println("My new age is :",(myAge+1))
	} else {
		fmt.Println("Year is different")
	}
}