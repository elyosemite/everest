package main

import "fmt"

func main() {
	interpreted_string := "Hello, how are you?\nI'm Software Developer."
	literal_string := `"Hello, how are you?\nI'm Software Developer."`

	fmt.Println(interpreted_string)
	fmt.Println(literal_string)
}