package main

import (
    "fmt"
    "strings"
)

func isPalindrome(input string) bool {
    input = strings.ToLower(input)
    for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
        if input[i] != input[j] {
            return false
        }
    }
    return true
}

func main() {
    var word string
    fmt.Print("Enter a word or phrase: ")
    fmt.Scan(&word)

    if isPalindrome(word) {
        fmt.Println("It's a palindrome!")
    } else {
        fmt.Println("It's not a palindrome.")
    }
}
