package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numFlips int
	fmt.Print("Enter the number of times to flip the coin: ")
	fmt.Scan(&numFlips)

	if numFlips <= 0 {
		fmt.Println("Please enter a positive integer for the number of flips.")
		return
	}

	headsCount := 0
	tailsCount := 0

	for i := 0; i < numFlips; i++ {
		// Generate a random number between 0 and 1
		randomValue := rand.Float64()

		if randomValue < 0.5 {
			tailsCount++
		} else {
			headsCount++
		}
	}

	totalFlips := float64(numFlips)
	headsPercentage := (float64(headsCount) / totalFlips) * 100
	tailsPercentage := (float64(tailsCount) / totalFlips) * 100

	fmt.Printf("Heads: %.2f%%\n", headsPercentage)
	fmt.Printf("Tails: %.2f%%\n", tailsPercentage)
}
