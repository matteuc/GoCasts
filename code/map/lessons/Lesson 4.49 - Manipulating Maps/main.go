package main

import "fmt"

// var colors map[string]string

func main() {

	// Equivalent to the line above
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	// Works with variables of type struct but NOT with those of type map
	// colors.white = "#ffffff"

	// Deletes the key-value pair of with a key of "white" in the map "colors"
	delete(colors, "white")

	fmt.Println(colors)
}
