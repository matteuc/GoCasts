package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// The parentheses before the 'print()' keyword indicates that the receiver of this function is of type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}