package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){	
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Convert variable of type deck to type slice of strings
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// Check if an error exists
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		
		// Option #2 - log the error and entirely quit the program (preferred)
		fmt.Println("Error: ", err)
		
		os.Exit(1)

	}
	
	ss := strings.Split(string(bs), ",")

	return deck(ss)

}

func (d deck) shuffle() {
	for i := range d {
		// Generate a random number of any int value from 0 to the max index of 'd'
		newPosition := rand.Intn(len(d) - 1)

		// Reassign two strings in 'd' to two new values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}