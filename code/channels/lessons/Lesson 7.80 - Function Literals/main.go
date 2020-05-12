package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c {

		// Syntax for a function literal (or an anonymous function that does not have a name)
		go func() {
			time.Sleep(time.Second)
			checkLink(l, c) // The initial value of 'l' is captured by this function literal
		}()
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // A blocking call (the next line of code does not execute until this one returns)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link

}
