// / achieve generic using any type
// package main

// import (
// 	"fmt"
// )

// func Print[T any](s []T) {
// 	for _, v := range s {
// 		fmt.Print(v)
// 	}
// }

// func main() {
// 	Print([]string{"Hello, ", "playground\n"})
// 	Print([]int{1, 2, 3})
// }

// /using multiple types
package main

import "fmt"

func DemoGenericFun[age int64 | float64](myage age) {
	fmt.Println(myage)
}

func main() {
	var intAge int64 = 23
	var floatAge float64 = 45.5
	DemoGenericFun(intAge)
	DemoGenericFun(floatAge)
}
