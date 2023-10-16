package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 42
	}()

	go func() {
		ch2 <- "Hello, Golang"
	}()

	select {
	case value := <-ch1:
		fmt.Println("Received from ch1:", value)
	case message := <-ch2:
		fmt.Println("Received from ch2:", message)
	}
}
