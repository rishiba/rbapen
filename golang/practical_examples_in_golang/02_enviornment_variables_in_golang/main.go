package main

import (
	"fmt"
	"os"
)

func main() {

	envVariables := []string{"ENV1", "ENV2"}

	for _, variable := range envVariables {
		fmt.Printf("\nValue of %s is %s\n", variable, os.Getenv(variable))
	}
}
