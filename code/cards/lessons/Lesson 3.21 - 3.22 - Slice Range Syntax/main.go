package main

func main() {

	cards := newDeck()

	// Assign both return values into two new variables
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	
	cards = remainingDeck

	cards.print()
}

