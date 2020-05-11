package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	president := person{"Rob", "Anderson"}

	president.firstName = "Donald"

	president.lastName = "Trump"

	fmt.Printf("%+v\n", president)

}
