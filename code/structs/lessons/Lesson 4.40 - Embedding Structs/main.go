package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// the field "contact" is an embedded struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	presidentContact := contactInfo{
		email:   "potus@usa.com",
		zipCode: 90210,
	}

	president := person{
		firstName: "Rob",
		lastName:  "Anderson",
		contact:   presidentContact,
	}

	fmt.Printf("%+v\n", president)

}
