package main

import (
	"fmt"
	"strconv"
)

// Person : Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// increaseAge method (pointer receiver)
func (p *Person) increaseAge() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.increaseAge()
	person1.getMarried("William")
	fmt.Println(person1.greet())

	// Alternative
	person2 := Person{"Ken", "Smith", "Boston", "m", 30}
	// fmt.Println(person2)
	person2.getMarried("William")
	fmt.Println(person2.greet())
}
