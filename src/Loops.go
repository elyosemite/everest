package main

import "fmt"

const favColor string = "blue"

func main() {
	i := 0
	rvariable := []string{"Yuri Melo", "Pedro Da Silva", "Kim Jong un"}
	mmap := map[int]string {
		22: "Yuri Melo",
		33: "GFG",
		44: "Igor Melo",
	}

	for i := 0; i < 4; i++ {
		fmt.Println("I'm back-end Developer")
	}

	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	// Infinity loop
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

	// Show the sharks
	ShowMeSharks()
	ShowMyFavoriteColor()
}

func ShowMeSharks() {
	sharks := []string{ "hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem" }

	for _, shark := range sharks {
		fmt.Println(shark)
	}
}

func ShowMyFavoriteColor() {
	var guess string

	for {
		// Ask the user to guess my favorite color
		fmt.Println("Guess my favorite color: ")

		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// Did they guess the correct color?
		if favColor == guess {
			// They guessed it!
			fmt.Printf("%q is my favorite color!\n", favColor)
			return
		}

		// Wrong! Have them guess again
		fmt.Printf("Sorry, %q is not my favorite color. Guess again.\n", guess)
	}
}