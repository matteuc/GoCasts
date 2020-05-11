package main

import "fmt"

// Declaring a new type "person" that is of type struct with the following properties
type person struct {
	firstName string
	lastName  string
}

func main() {
	// One way to intialize a struct; order of values in initializer correspond to order of fields in struct definition
	president := person{"Rob", "Anderson"}

	// Another way to initialize a struct by directly referencing the field the value is intended to be
	vicePresident := person{lastName: "Pence", firstName: "Mike"}

	// Declaring an empty struct assigns "zero values" to all its properties
	// string: ""
	// int: 0
	// float: 0
	// bool: false

	secretaryOfState := person{}

	// Printing a struct along with all its fields
	fmt.Printf("%+v\n", president)

	// Should print out "{firstName: lastName: }"
	fmt.Printf("%+v\n", secretaryOfState)

	// Printing specific fields from a struct
	fmt.Printf("%v %v is the current president\n", vicePresident.firstName, vicePresident.lastName)

}
