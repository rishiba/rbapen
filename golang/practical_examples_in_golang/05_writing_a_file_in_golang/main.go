package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeToFile(fileName string) {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	if err == nil {
		fmt.Fprintf(os.Stderr, "\nSuccess in opening file")
		defer fd.Close()
	} else {
		fmt.Fprintf(os.Stderr, "\nError in opening file")
	}

	buffer := make([]byte, 256)

	for i := 0; i < 256; i++ {
		buffer[i] = 'a'
	}

	bytesWritten, err := fd.Write(buffer)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError in writing file")
	} else {
		fmt.Fprintf(os.Stderr, "\nSuccessful in writing file, bytesWritten is %d", bytesWritten)
	}
}

func writeBuffered(fileName string) {
	fd, err := os.OpenFile(fileName, os.O_RDWR, 0755)
	if err == nil {
		fmt.Fprintf(os.Stderr, "\nSuccess in opening file")
		defer fd.Close()
	} else {
		fmt.Fprintf(os.Stderr, "\nError in opening file, %s", err)
		os.Exit(-1)

	}

	bufWriter := bufio.NewWriter(fd)

	bufToWrite := []byte("This is a sample sentance")

	bytesWritten, err := bufWriter.Write(bufToWrite)

	if err == nil {
		fmt.Fprintf(os.Stderr, "\nSuccessfully wrote %d bytes in the file", bytesWritten)
		bufWriter.Flush() // You must flush so that the buffer will be actually written
	} else {
		fmt.Fprintf(os.Stderr, "\nError in writing data to file")
	}
}

func main() {
	writeToFile("temp")
	writeBuffered("temp")
}
