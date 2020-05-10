package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	// The underscore indicates that this variable is not being used
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}

	return cards
}

// The parentheses before the 'print()' keyword indicates that the receiver of this function is of type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}