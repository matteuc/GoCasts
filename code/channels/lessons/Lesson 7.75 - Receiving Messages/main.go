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
	// Waiting for messages is considered a blocking call (blocks this routine until any message is received)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// As there are only five links being checked, the line below causes the program to wait indefinitely (as no more data is being sent on this channel)
	fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // A blocking call (the next line of code does not execute until this one returns)
	if err != nil {
		c <- (link + " should be down!") // Sends error into channel
		return
	}

	c <- (link + " should be up!") // Sends success into channel

}
