package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// Instance variable omitted as not used; only receiver type needs to be defined
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)

	// Throws error as printGreeting takes in a variable of type 'englishBot'
	printGreeting(sb)

}
