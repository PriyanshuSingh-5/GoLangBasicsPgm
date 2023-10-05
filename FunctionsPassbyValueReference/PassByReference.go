// /pass pointers
// package main

// import "fmt"

// func modifyValueByReference(x *int) {
// 	*x = 42
// }

// func main() {
// 	num := 10
// 	modifyValueByReference(&num)
// 	fmt.Println("Modified value:", num) // Prints "Modified value: 42"
// }

///pass pointer to slice
// package main

// import "fmt"

// func modifySliceByReference(slice *[]int) {
//     (*slice)[0] = 42
// }

// func main() {
//     numbers := []int{1, 2, 3}
//     modifySliceByReference(&numbers)
//     fmt.Println("Modified slice:", numbers) // Prints "Modified slice: [42 2 3]"
// }
