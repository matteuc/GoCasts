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

	// Makes a new channel that will pass data of type 'string'
	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)

	}

	// Once a message is in the channel, print out that message
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // A blocking call (the next line of code does not execute until this one returns)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- (link + " should be down!") // Sends error into channel
		return
	}

	fmt.Println(link, "is up!")
	c <- (link + " should be up!") // Sends success into channel

}
