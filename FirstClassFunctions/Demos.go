///assign func to variable
// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func main() {
//     operation := add
//     result := operation(3, 5)
//     fmt.Println(result) // Prints "8"
// }

///pass func as arguement
// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func compute(operation func(int, int) int, x, y int) int {
//     return operation(x, y)
// }

// func main() {
//     result := compute(add, 3, 5)
//     fmt.Println(result) // Prints "8"
// }

///return func from func
// package main

// import "fmt"

// func multiplier(factor int) func(int) int {
//     return func(x int) int {
//         return x * factor
//     }
// }

// func main() {
//     double := multiplier(2)
//     result := double(5)
//     fmt.Println(result) // Prints "10"
// }

///storing func in data structure
// package main

// import "fmt"

// func add(a, b int) int {
//     return a + b
// }

// func subtract(a, b int) int {
//     return a - b
// }

// func main() {
//     operationList := []func(int, int) int{add, subtract}

//     for _, operation := range operationList {
//         result := operation(10, 5)
//         fmt.Println(result)
//     }
// }
