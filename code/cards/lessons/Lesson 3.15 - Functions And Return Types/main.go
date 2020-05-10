package main


import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// The below function will return an error as the compiler is unaware what type of variable the function returns
// -------------------------------------------------------------------------------------------------------------
// func newCard() {
// 	return "Five of diamonds"
// } 

func newCard() string {
	return "Five of diamonds"
} 
