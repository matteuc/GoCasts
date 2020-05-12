package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

// Read contents of a file and log it to the terminal
func main() {
	if len(os.Args) == 1 {
		fmt.Println("You must provide a file to read from.")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("There are too many arguments\nUsage (.exe): print_file [filename]\nUsage: go run print_file.go [filename]")
		os.Exit(1)
	}

	// The first argument is the command, any following values are additional arguments
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
