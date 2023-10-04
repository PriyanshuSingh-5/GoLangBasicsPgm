package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Print("Enter three numbers separated by spaces: ")
	fmt.Scan(&num1, &num2, &num3)

	max := num1

	if num2 > max {
		max = num2
	}
	if num3 > max {
		max = num3
	}

	fmt.Printf("The largest number is %d\n", max)
}
