//create slice
package main

// import "fmt"

// func main() {
//     // Create a slice of integers
//     numbers := []int{1, 2, 3, 4, 5}

//     // Print the slice
//     fmt.Println("Slice:", numbers)
// }


///length and capacity
// package main

// import "fmt"

// func main() {
//     // Create an empty slice of integers
//     numbers := make([]int, 0, 5)

//     fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

//     // Append elements to the slice
//     numbers = append(numbers, 1, 2, 3)

//     fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))
// }



///copy data from one to other
// package main

// import "fmt"

// func main() {
//     // Create two slices of integers
//     slice1 := []int{1, 2, 3}
//     slice2 := make([]int, len(slice1))

//     // Copy elements from slice1 to slice2
//     copy(slice2, slice1)

//     fmt.Println("Original Slice:", slice1)
//     fmt.Println("Copied Slice:", slice2)
// }



///append to slice
package main

import "fmt"

func main() {
    // Create an empty slice of integers
    numbers := make([]int, 0)

    // Append elements to the slice dynamically
    numbers = append(numbers, 1, 2, 3, 4, 5)

    fmt.Println("Updated Slice:", numbers)
}







