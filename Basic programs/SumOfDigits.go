package main

import (
    "fmt"
)

func main() {
    var num, digit, sum int
    fmt.Print("Enter an integer: ")
    fmt.Scan(&num)

    originalNum := num
    for num != 0 {
        digit = num % 10
        sum += digit
        num /= 10
    }

    fmt.Printf("The sum of digits of %d is %d\n", originalNum, sum)
}
