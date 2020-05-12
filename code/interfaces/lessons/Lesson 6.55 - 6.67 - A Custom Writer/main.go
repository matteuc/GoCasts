package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// The type logWriter implements the Writer interface so it can be passed into 'io.Copy()'

	// Response has a value named Body that implements the Reader interface so it can be passed into 'io.Copy()'
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just Wrote this many bytes: ", len(bs))

	return len(bs), nil
}
