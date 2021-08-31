package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo contact
}

type contact struct {
	email   string
	zipcode int
}

func main() {
	p1 := person{firstName: "Tangi", lastName: "Tomson"}
	p1.contactInfo.email = "aaaaaaa"
	p1.contactInfo.zipcode = 23435
	p1.updateName("Tangeri")
	p1.print()

	var p2 person
	p2.firstName = "Meowy"
	p2.lastName = "Meowington"

	p2.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pp *person) updateName(newFirstName string) {
	(*pp).firstName = newFirstName
}
