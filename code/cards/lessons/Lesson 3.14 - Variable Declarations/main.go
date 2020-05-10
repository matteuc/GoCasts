package main


import "fmt"

func main() {
	// var card string = "Ace of Spades"

	// Equivalent to the line above (':=' infers data type); used only when declaring new variables
	card := "Ace of Spades" 

	card = "Five of Diamonds"

	fmt.Println(card)
}