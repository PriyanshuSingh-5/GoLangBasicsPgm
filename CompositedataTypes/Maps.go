///create a map and add data

// package main

// import (
//     "fmt"
// )

// func main() {
//     // Create an empty map with string keys and int values
//     studentGrades := make(map[string]int)

//     // Add data to the map
//     studentGrades["Alice"] = 95
//     studentGrades["Bob"] = 87
//     studentGrades["Charlie"] = 92

//     // Print the map
//     fmt.Println("Student Grades:", studentGrades)
// }

////deleting entries, and checking if a key exists
// package main

// import (
//     "fmt"
// )

// func main() {
//     // Create an empty map
//     studentGrades := make(map[string]int)

//     // Add entries to the map
//     studentGrades["Alice"] = 95
//     studentGrades["Bob"] = 87
//     studentGrades["Charlie"] = 92

//     // Delete an entry
//     delete(studentGrades, "Bob")

//     // Check if a key exists
//     _, aliceExists := studentGrades["Alice"]
//     _, bobExists := studentGrades["Bob"]

//     fmt.Println("Student Grades:", studentGrades)
//     fmt.Println("Alice exists:", aliceExists)
//     fmt.Println("Bob exists:", bobExists)
// }

////check frequency of a word
// package main

// import (
//     "fmt"
//     "strings"
// )

// func main() {
//     text := "This is a sample text. Sample text is a good way to learn programming."

//     words := strings.Fields(text)
//     wordFrequency := make(map[string]int)

//     for _, word := range words {
//         wordFrequency[word]++
//     }

//     fmt.Println("Word Frequency:", wordFrequency)
// }

///count characters of given string
// package main

// import (
//     "fmt"
// )

// func main() {
//     text := "hello world"
//     charFrequency := make(map[rune]int)

//     for _, char := range text {
//         charFrequency[char]++
//     }

//     fmt.Println("Character Frequency:", charFrequency)
// }
