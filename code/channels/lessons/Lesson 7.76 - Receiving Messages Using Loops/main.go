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

	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)

	}

	// Waits for 5 (the length of the links slice) messages to be sent on the channel...
	// Then the main Go routine exits
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // A blocking call (the next line of code does not execute until this one returns)
	if err != nil {
		c <- (link + " should be down!") // Sends error into channel
		return
	}

	c <- (link + " should be up!") // Sends success into channel

}
