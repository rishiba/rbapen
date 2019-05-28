package main

import (
	"fmt"
	"os"
)

// The program demonstrates logging to stderr and logging to a log file.
func main() {

	LOGFILE, ok := os.LookUpEnv("LOGFILE")
	if ok != true {
		fmt.Printf("\nLOGFILE is not set %s", LOGFILE)
		os.Exit(-1)
	}
}
