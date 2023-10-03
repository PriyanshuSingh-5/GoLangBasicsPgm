package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Enter the number of terms: ")
    fmt.Scan(&n)

    a, b := 0, 1
    fmt.Printf("Fibonacci Series (first %d terms): ", n)
    for i := 0; i < n; i++ {
        fmt.Print(a, " ")
        a, b = b, a+b
    }
    fmt.Println()
}
