package src

import "fmt"

const favColor string = "blue"

func Loops() {
	i := 0
	rvariable := []string{"Yuri Melo", "Pedro Da Silva", "Kim Jong un"}
	mmap := map[int]string{
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

	// Loops
	ForClauseLoop()
	NoRangeClauseLoop()
	RangeClauseLoop()
}

/*
	for [ Initial Statement ] ; [ Condition ] ; [ Post Statement ] {
		[Action]
	}
*/
func ForClauseLoop() {
	fmt.Println("Start. ForClause: ")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("End. ForClause: ")
}

func NoRangeClauseLoop() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}
}

func RangeClauseLoop() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}

	// Without i variable
	for _, shark := range sharks {
		fmt.Println(shark)
	}

	// Iterate over item
	for range sharks {
		sharks = append(sharks, "shark")
	}

	fmt.Printf("%q\n", sharks)

	// Iterate over map
	sammyShark := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	for key, value := range sammyShark {
		fmt.Println(key + ": " + value)
	}
}

func ShowMeSharks() {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

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
