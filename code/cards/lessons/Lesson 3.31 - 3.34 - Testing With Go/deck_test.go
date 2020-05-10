package main

import (
	"os"
	"testing"
)

// 't' refers to the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		// The statement '%v' declares a formatting directive and is substituted with the arguments following the format string

		t.Errorf("Expected deck length of 48, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"

	os.Remove(testFile)

	d := newDeck()

	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(d))
	}

	os.Remove(testFile)
}
