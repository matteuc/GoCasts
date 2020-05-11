package main

import "fmt"

func main() {
	// map[string]string indicates an "object" that will have keys of type "string" and matching values of type "string"
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#12",
	}

	fmt.Println(colors)
}
