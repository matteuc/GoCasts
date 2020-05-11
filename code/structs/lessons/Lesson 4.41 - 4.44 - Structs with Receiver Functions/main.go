package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// A different way of embeddng a struct; the field is named after the type
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	presidentContact := contactInfo{
		email:   "potus@usa.com",
		zipCode: 90210,
	}

	president := person{
		firstName:   "Rob",
		lastName:    "Anderson",
		contactInfo: presidentContact,
	}

	// The & operator retrieves the memory address of the variable following it
	presidentPointer := &president

	presidentPointer.updateName("Donald")

	president.print()

}

// Does not correctly update the caller's firstName as "p" is a copy of the passed in value
// ------------------------------------------------------
// func (p person) updateName(nf string) {
// 	p.firstName = nf
// }

// *person refers to a type of person pointer
func (p *person) updateName(nf string) {
	// The * operator followed by a variable of type pointer retrieves the value at the specified address
	(*p).firstName = nf
}

// A receiver function that takes in a struct type
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
