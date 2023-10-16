package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 goroutines to the wait group

	go func() {
		defer wg.Done() // Mark this goroutine as done
		// Perform some work
	}()

	go func() {
		defer wg.Done()
		// Perform some work
	}()

	wg.Wait() // Wait for both goroutines to finish
}
