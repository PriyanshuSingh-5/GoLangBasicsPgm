package main

import "fmt"

// Define a Shape interface with an Area() method.
type Shape interface {
	Area() float64
}

// Define a Rectangle struct that implements the Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Define a Circle struct that also implements the Shape interface.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}

func main() {
	// Create instances of Rectangle and Circle.
	rect := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 3}

	// Create a slice of Shape interface containing both shapes.
	shapes := []Shape{rect, circle}

	// Calculate and print the areas of the shapes.
	for _, shape := range shapes {
		fmt.Printf("Area of shape: %f\n", shape.Area())
	}
}
