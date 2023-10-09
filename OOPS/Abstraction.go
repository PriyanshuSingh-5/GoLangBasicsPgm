package main

///using interface
// package main

// import (
//     "fmt"
// )

// // Animal is an interface representing common behaviors of animals.
// type Animal interface {
//     Speak() string
//     Move() string
// }

// // Dog represents a dog with a name.
// type Dog struct {
//     Name string
// }

// // Implement the Speak method for Dog.
// func (d Dog) Speak() string {
//     return "Woof!"
// }

// // Implement the Move method for Dog.
// func (d Dog) Move() string {
//     return "Running on four legs"
// }

// // Bird represents a bird with a name.
// type Bird struct {
//     Name string
// }

// // Implement the Speak method for Bird.
// func (b Bird) Speak() string {
//     return "Chirp!"
// }

// // Implement the Move method for Bird.
// func (b Bird) Move() string {
//     return "Flying in the sky"
// }

// func main() {
//     // Create instances of Dog and Bird.
//     dog := Dog{Name: "Buddy"}
//     bird := Bird{Name: "Robin"}

//     // Create a slice of Animal interfaces containing both animals.
//     animals := []Animal{dog, bird}

//     // Interact with the animals.
//     for _, animal := range animals {
//         fmt.Printf("%s says: %s\n", animal.(Animal).Speak(), animal.(Animal).Move())
//     }
// }

////Using struct
// package main

import "fmt"

// Defining a struct to represent a car
type Car struct {
	Make  string
	Model string
	Year  int
}

// Defining a method on the Car struct
func (c Car) Start() {
	fmt.Printf("Starting the %s %s\n", c.Make, c.Model)
}

func main() {
	// Creating an instance of the Car struct
	myCar := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2022,
	}

	// Using the Start method to abstract away the details of starting the car
	myCar.Start()
}
