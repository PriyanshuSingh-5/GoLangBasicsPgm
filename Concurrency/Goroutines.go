package main

func main() {
	go myFunction() // Start a new goroutine
	// Continue with other work
}

func myFunction() {
	// Code to run concurrently
}
