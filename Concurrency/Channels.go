package main

import "fmt"

func main() {
	ch := make(chan int) // Create a new channel
	go sendValue(ch)
	value := <-ch // Receive a value from the channel
	fmt.Println(value)
}

func sendValue(ch chan int) {
	ch <- 42 // Send a value to the channel
}
