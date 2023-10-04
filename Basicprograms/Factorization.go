package main

import (
	"fmt"
)

func primeFactors(N int) {
	fmt.Printf("Prime factors of %d are: ", N)

	for i := 2; i*i <= N; i++ {
		for N%i == 0 {
			fmt.Printf("%d ", i)
			N /= i
		}
	}

	if N > 1 {
		fmt.Printf("%d", N)
	}

	fmt.Println()
}

func main() {
	var N int

	fmt.Print("Enter a number to find its prime factors: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	if N <= 1 {
		fmt.Println("Prime factorization is not applicable for numbers less than or equal to 1.")
		return
	}

	primeFactors(N)
}
