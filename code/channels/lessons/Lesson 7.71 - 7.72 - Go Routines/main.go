package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		// When a function is invoked with the keyword 'go', a new thread (or Go routine) is created to run the specified function

		// When a new Go routine is created, the Go scheduler runs ONE routine until it finishes or stops at a blocking call

		// This is referred to as CONCURRENCY (loading multiple routines, but only working on ONE Go routine)

		// With the addition of additional CPU cores, Go Routines are running in PARALLEL (multiple routines running in parallel on multiple cores)
		go checkLink(link)

		// Therefore the above line causes the program to quit as the completion of the Main routine (Go Routine started when the 'main()' is ran) is prioritized over child (non-main) Go routines are stopped in favor for the completion of the Main routine (Go Routine started when the 'main()' is ran)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // A blocking call (the next line of code does not execute until this one returns)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
