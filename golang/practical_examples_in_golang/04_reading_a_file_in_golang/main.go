package main

import (
	"fmt"
	"io/ioutil"

	"bufio"
	"io"
	"os"
)

/*
 * We need to create a new buffered reader and pass the io.Reader interface of
 * the opened file. Then a simple call to bufio.ReadLine() will read a line of
 * the file.
 */
func LineByLineReading(fileToRead string) {
	var bufReader *bufio.Reader
	fd, err := os.Open(fileToRead)

	if err == nil {
		fmt.Fprintf(os.Stderr, "\nOpened the file successfully")
		defer fd.Close()
	} else {
		fmt.Fprintf(os.Stderr, "\nError opening file, %s, error is %s", fileToRead, err)
		os.Exit(-1)
	}

	bufReader = bufio.NewReader(fd)

	for {
		line, _, err := bufReader.ReadLine()

		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "\nFile read is complete")
			return
		}

		fmt.Fprintf(os.Stdout, "\n%s", line)
	}
}

/*
 * Get a new bufio with a io.Reader. Then keep reading till you get nil.
 */
func BufferedReading(fileToRead string) {
	var bufReader *bufio.Reader
	fd, err := os.Open(fileToRead)

	if err == nil {
		fmt.Fprintf(os.Stderr, "\nOpened the file successfully")
		defer fd.Close()
	} else {
		fmt.Fprintf(os.Stderr, "\nError opening file, %s, error is %s", fileToRead, err)
		os.Exit(-1)
	}

	bufReader = bufio.NewReader(fd)

	buffer := make([]byte, 256)

	for {
		bytesRead, err := bufReader.Read(buffer)

		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "\nFile read is complete")
			return
		}

		if bytesRead != 0 {
			fmt.Fprintf(os.Stdout, "%s", buffer[:bytesRead])
		}
	}

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

	/*
	 * ReadingFile1(fileToRead)
	 * ReadingFile2(fileToRead)
	 */

	BufferedReading(fileToRead)
	LineByLineReading(fileToRead)
}
