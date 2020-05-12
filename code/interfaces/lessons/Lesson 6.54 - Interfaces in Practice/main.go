package main

import "fmt"

// Any type that has a function called 'getGreeting' that returns a string, that type is also of type 'bot'
type bot interface {
	getGreeting() string
}

// A variable of type 'bot' cannot be created as it is an interface type

// Variables can only be created from concrete types (non-interface types like englishBot, string, int, map)

type englishBot struct{} // implicitly defined as type 'bot'
type spanishBot struct{} // implicitly defined as type 'bot'

// Instance variable omitted as not used; only receiver type needs to be defined
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// Because an englishBot has a function called 'getGreeting' that returns a string, it can be used with the function 'printGreeting'
	printGreeting(eb)
	printGreeting(sb)

}
