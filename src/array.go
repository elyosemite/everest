package src

import "fmt"

func Arrays() {
	// The capacity of an array is defined at creation time.

	coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(coral)
	fmt.Println(primes)
}
