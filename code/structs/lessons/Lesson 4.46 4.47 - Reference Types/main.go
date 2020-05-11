package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "Rob"}

	fmt.Println(mySlice)

	updateName(mySlice)

	fmt.Println(mySlice)

}

// This function will actually alter the "mySlice" variable as slices is a reference type

// Other reference types include: maps, channels, pointers, functions

// All other types are value types: int, float, string, bool, struct

func updateName(s []string) {
	// As usual, the variable "s" is copied

	// However because slices are actually referring to an array somewhere else in memory...

	// When the copy of the original slice "s" is edited, the original array is still changed

	s[0] = "Bye"
}
