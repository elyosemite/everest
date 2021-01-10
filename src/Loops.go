package main

import "fmt"

func main() {
	i := 0

	for i := 0; i < 4; i++ {
		fmt.Println("I'm back-end Developer")
	}

	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	/*for {
		fmt.Println("I am infinity loop")
	}*/
}