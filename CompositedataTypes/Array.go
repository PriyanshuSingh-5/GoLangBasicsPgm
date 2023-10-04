//Write a program that calculates the sum of all elements in an integer array.
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0

    for _, num := range numbers {
        sum += num
    }

    fmt.Printf("Sum of elements: %d\n", sum)
}

// Largest Element
// package main

// import "fmt"

// func main() {
//     numbers := []int{10, 5, 7, 21, 14}
//     max := numbers[0]

//     for _, num := range numbers {
//         if num > max {
//             max = num
//         }
//     }

//     fmt.Printf("Largest element: %d\n", max)
// }

///Reverse an array
// package main

// import "fmt"

// func main() {
//     numbers := []int{1, 2, 3, 4, 5}
//     length := len(numbers)

//     for i := 0; i < length/2; i++ {
//         numbers[i], numbers[length-i-1] = numbers[length-i-1], numbers[i]
//     }

//     fmt.Println("Reversed array:", numbers)
// }


///length of array
package main

import "fmt"

func main() {
    // Create an empty slice of integers
    numbers := make([]int, 0)

    // Append elements to the slice dynamically
    numbers = append(numbers, 1, 2, 3, 4, 5)

    fmt.Println("Updated Slice:", numbers)
}






