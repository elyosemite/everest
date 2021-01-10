package main

import "fmt"

func main() {
	i := 0
	rvariable := []string{"Yuri Melo", "Pedro Da Silva", "Kim Jong un"}
	mmap := map[int]string {
		22: "Yuri Melo",
		33: "GFG",
		44: "Igor Melo",
	}
	// Using channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()

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

	for i, j := range "Yuri dos Santos Melo" {
		fmt.Printf("The index number of %U is %d\n", j, i)
	}

	// Loop in maps
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	// Loop in channel
	for i := range chnl {
		fmt.Println(i)
	}

}