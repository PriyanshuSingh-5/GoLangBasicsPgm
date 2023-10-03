package main

import (
	"fmt"
)

func main() {
	var length, width float64
	fmt.Print("Enter the length: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width: ")
	fmt.Scan(&width)

	area := length * width
	fmt.Printf("The area of the rectangle is %.2f square units.\n", area)
}
