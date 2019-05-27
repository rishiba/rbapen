package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args have all the command line arguments. As with many languages the
	// first argument always is the program name.

	fmt.Printf("\nProgram name is %s", os.Args[0])
	fmt.Printf("\nNumber of arguments is %d", len(os.Args))

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("\nArgument number %d is : %s", i, os.Args[i])
	}
}
