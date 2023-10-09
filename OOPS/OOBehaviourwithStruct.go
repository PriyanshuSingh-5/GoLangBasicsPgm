package main

import (
	"fmt"
)

// Define a struct named 'Person' to represent a person object.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Define a method named 'FullName' for the 'Person' struct.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Define another method named 'Greetings' for the 'Person' struct.
func (p Person) Greetings() string {
	return fmt.Sprintf("Hello, my name is %s, and I am %d years old.", p.FullName(), p.Age)
}

func main() {
	// Create a new instance of the 'Person' struct.
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Call the methods on the 'person' object.
	fmt.Println(person.FullName())
	fmt.Println(person.Greetings())
}
