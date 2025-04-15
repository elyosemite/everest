package src

import (
	"fmt"
	"os"
)

func HelloWorld() {
	var s, sep string
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
