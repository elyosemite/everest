package main

import "fmt"

func main() {
	// My slice type
	// []int{-3, -2, -1, 0, 1, 2, 3}
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	fmt.Println(seaCreatures)

	seaCreatures = append(seaCreatures, "seahourse")

	fmt.Println(seaCreatures)
}
