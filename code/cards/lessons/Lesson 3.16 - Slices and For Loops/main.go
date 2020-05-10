package main


import "fmt"

func main() {

	// Two types of lists
	// ===============================================================
	// 1. Arrays: Fixed length list
	// 2. Slices: Lists that can change length

	// Declaring a new slice
	cards := []string{newCard(), "Ace of Diamonds"}

	// Adding a new item to the slice
	cards = append(cards, "Six of Spades")

	// Iterating through a slice using a for-loop
	// The ':=' symbol is used to indicate that for every iteration, we want to reassign the values of 'i' and 'card' to new iteration's values
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of diamonds"
} 
