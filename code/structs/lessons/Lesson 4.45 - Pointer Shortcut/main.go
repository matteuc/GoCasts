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

	president.updateName("Donald")

	president.print()

}

// Go provides a shortcut so that if the receiver is of type "person" pointer BUT the invoking variable is of type "person", Go automatically coverts the receiver to type "person" pointer
func (p *person) updateName(nf string) {
	(*p).firstName = nf
}

// A receiver function that takes in a struct type
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
