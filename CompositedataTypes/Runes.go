///count number of vowels
// package main

// import (
//     "fmt"
// )

// func main() {
//     text := "Hello, World!"
//     vowelCount := 0

//     for _, char := range text {
//         switch char {
//         case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
//             vowelCount++
//         }
//     }

//     fmt.Printf("Number of vowels: %d\n", vowelCount)
// }

///reverse a string
// package main

// import (
//     "fmt"
// )

// func main() {
//     text := "Hello, World!"
//     reversed := reverseString(text)

//     fmt.Println("Original: ", text)
//     fmt.Println("Reversed: ", reversed)
// }

// func reverseString(input string) string {
//     // Convert the input string to a rune slice
//     inputRunes := []rune(input)
//     length := len(inputRunes)

//     // Reverse the rune slice
//     for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
//         inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
//     }

//     // Convert the reversed rune slice back to a string
//     reversed := string(inputRunes)

//     return reversed
// }
