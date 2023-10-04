package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero!")
			return
		}
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
