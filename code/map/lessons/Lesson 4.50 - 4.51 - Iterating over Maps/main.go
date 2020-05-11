package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	// When iterating over a map, the first assigned variable matches to the key and the second the value
	for color, hex := range c {
		fmt.Printf("The color %v has a hex value of %v\n", color, hex)
	}
}

// In addition some big differences between maps and struct variables are...

// Maps are iterable...Structs are not
// Maps are reference types...Structs are value types
// Maps can have undeclared key-value pairs...Structs must have all key-value pairs declared at compile time

// Maps are ideally used when all values are closely related
