package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://stackoverflow.com/questions/39859222/golang-how-to-overcome-scan-buffer-limit-from-bufio

func findWord(fileName string, searchString string) {

	fd, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError opening file, error is %s. Exiting", err)
		os.Exit(-1)
	}

	var myScanner *bufio.Scanner
	myScanner = bufio.NewScanner(fd)
	capacity := 256
	buffer := make([]byte, capacity)

	myScanner.Buffer(buffer, capacity)

	myScanner.Split(bufio.ScanWords)

}

func scanWordsFromString() {
	var myScanner *bufio.Scanner
	myStr := "a quick brown fox jumps over a lazy dog"

	var stringReader io.Reader

	stringReader = strings.NewReader(myStr)
	myScanner = bufio.NewScanner(stringReader)

	myScanner.Split(bufio.ScanWords)
	var count int
	count = 0

	for myScanner.Scan() {
		count++
	}

	fmt.Printf("\nNumber of words is %d", count)
}

func smallBufferSizeOfScanner() {

	// https://stackoverflow.com/questions/39859222/golang-how-to-overcome-scan-buffer-limit-from-bufio

}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\nGive the command as ./command FileName WordToSearch")
		os.Exit(-1)
	}

	scanWordsFromString()

	fileName := os.Args[1]
	searchString := os.Args[2]

	findWord(fileName, searchString)
	fmt.Println("vim-go")
}
