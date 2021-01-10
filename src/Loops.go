package main

import "fmt"

func main() {
	i := 0
	rvariable := []string{"Yuri Melo", "Pedro Da Silva", "Kim Jong un"}

	for i := 0; i < 4; i++ {
		fmt.Println("I'm back-end Developer")
	}

	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	// INfinity loop
	/*for {
		fmt.Println("I am infinity loop")
	}*/

	for i, j := range rvariable {
		fmt.Println(i, j)
	}

}