///Student strcut
package main

import (
    "fmt"
)

func main() {
    text := "Hello, World!"
    vowelCount := 0

    for _, char := range text {
        switch char {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            vowelCount++
        }
    }

    fmt.Printf("Number of vowels: %d\n", vowelCount)
}



///book Struct
// package main

// import "fmt"

// type Book struct {
//     Title  string
//     Author string
//     Price  float64
// }

// func main() {
//     // Create a slice of Book structs
//     books := []Book{
//         {Title: "Book 1", Author: "Author 1", Price: 29.99},
//         {Title: "Book 2", Author: "Author 2", Price: 19.99},
//         {Title: "Book 3", Author: "Author 3", Price: 34.99},
//     }

//     // Calculate and print total cost of books
//     totalCost := calculateTotalCost(books)
//     fmt.Printf("Total cost of books: $%.2f\n", totalCost)

//     // Display book details
//     displayBookDetails(books)
// }

// func calculateTotalCost(books []Book) float64 {
//     total := 0.0
//     for _, book := range books {
//         total += book.Price
//     }
//     return total
// }

// func displayBookDetails(books []Book) {
//     fmt.Println("Book Details:")
//     for _, book := range books {
//         fmt.Printf("Title: %s, Author: %s, Price: $%.2f\n", book.Title, book.Author, book.Price)
//     }
// }

