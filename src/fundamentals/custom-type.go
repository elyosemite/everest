package main

import "fmt"

type hotdog int

var my_own_hotdog hotdog = 10

func main() {
	my_other_hotdog := 10

	fmt.Printf("%v\n", my_other_hotdog)

	my_other_hotdog = int(my_own_hotdog)

	fmt.Printf("%v\n", my_other_hotdog)
}