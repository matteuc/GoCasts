package main

import "fmt"

func main() {
	// Runs without error as the function 'newCard()' is within the same package
	card := newCard()

	fmt.Println(card)
}
