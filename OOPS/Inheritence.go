///struct
// package main

// import "fmt"

// type Person struct {
//     FirstName string
//     LastName  string
//     Age       int
// }

// type Employee struct {
//     Person
//     EmployeeID int
// }

// func main() {
//     emp := Employee{
//         Person: Person{
//             FirstName: "John",
//             LastName:  "Doe",
//             Age:       30,
//         },
//         EmployeeID: 12345,
//     }

//     fmt.Println(emp.FirstName)
//     fmt.Println(emp.EmployeeID)
// }

// /interface
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159265359 * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 10},
		Circle{Radius: 3},
	}

	for _, shape := range shapes {
		fmt.Printf("Area of shape: %f\n", shape.Area())
	}
}
