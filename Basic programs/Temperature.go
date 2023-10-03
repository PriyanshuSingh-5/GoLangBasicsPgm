package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f°F is equal to %.2f°C\n", fahrenheit, celsius)
}
