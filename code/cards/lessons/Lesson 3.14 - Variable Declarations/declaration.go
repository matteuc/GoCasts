package main

import "fmt"

// You cannot assign a value to a variable outside of a function

// fruitCount := 5		THIS WILL NOT COMPILE

var fruitCount int

func logCount() {
	fruitCount = 5
	
	fmt.Println(fruitCount)
}
