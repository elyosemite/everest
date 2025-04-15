package src

import "fmt"

func Datatypes() {
	fmt.Println("Printing varaible value")

	a := 1
	b := "Yuri Melo"
	c := true
	d := `Yuri dos Santos Melo`
	e := 3.1415

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)

	fmt.Println("Show me datatypes")
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
}
