package main

import (
	"fmt"
	"io/ioutil"

	"bufio"
	"os"
)

func BufferedReading(fileToRead string) {
	var bufReader *bufio.Reader
	fd, err := os.Open(fileToRead)

	bufReader = bufio.NewReader(fd)

	buffer := make([]byte, 256)

	bytesRead, err := bufReader.Read(buffer)

}
func ReadingFile2(fileToRead string) {

	var buffer []byte
	var err error

	buffer, err = ioutil.ReadFile(fileToRead)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError reading the file %s", fileToRead)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "\nFile data is %s", buffer)

}

func ReadingFile1(fileToRead string) {
	var fd *os.File
	var err error
	fd, err = os.Open(fileToRead)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError opening the file %s. Error is %s", fileToRead, err)
	} else {
		fmt.Fprintf(os.Stderr, "\nSuccessful in opening file %s", fileToRead)
		defer fd.Close()
	}

	buffer := make([]byte, 256)
	/*
	 *
	 *     for {
	 *         bytesRead, err := fd.Read(buffer)
	 *
	 *         if bytesRead == 0 {
	 *             break
	 *         } else if err != nil {
	 *             fmt.Fprintf(os.Stderr, "\nError while reading file, error is %s", err)
	 *         } else {
	 *             fmt.Fprintf(os.Stderr, "\nBytes read is %d", bytesRead)
	 *             fmt.Fprintf(os.Stdout, "\nFile Contents is %s", buffer)
	 *         }
	 *     }
	 *
	 *     fmt.Fprintf(os.Stderr, "\nHere1")
	 *     fd.Seek(0, 0)
	 *     var bytesRead int
	 *
	 *     for {
	 *         bytesRead, err = fd.Read(buffer)
	 *
	 *         if bytesRead == 0 {
	 *             break
	 *         } else if err != nil {
	 *             fmt.Fprintf(os.Stderr, "\nError while reading file, error is %s", err)
	 *         } else {
	 *             fmt.Fprintf(os.Stderr, "\nBytes read is %d", bytesRead)
	 *             fmt.Fprintf(os.Stdout, "\nFile Contents is %s", buffer)
	 *         }
	 *     }
	 */

	fmt.Fprintf(os.Stderr, "\nHere 2")
	fd.Seek(0, 0)

	for {
		bytesReadx, err := fd.Read(buffer)
		if bytesReadx == 0 {
			fmt.Fprintf(os.Stdout, "--Reached the end of file")
			break
		} else if err != nil {
			fmt.Fprintf(os.Stdout, "\nError while reading file, error is %s", err)
		} else {
			fmt.Fprintf(os.Stdout, "\nBytes read is %d", bytesReadx)
			fmt.Fprintf(os.Stdout, "\nFile Contents is\n%s", buffer[:bytesReadx])
		}
	}
}

func main() {
	var fileToRead string

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "\nPlease pass the filename to read")
		os.Exit(-1)
	} else {
		fileToRead = os.Args[1]
	}

	ReadingFile1(fileToRead)
	ReadingFile2(fileToRead)

}
