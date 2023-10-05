// /create a file and add data
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	data := []byte("Hello, World!\n")
// 	err := ioutil.WriteFile("example.txt", data, 0644)
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	fmt.Println("File created successfully!")
// }

// /Read data from file
// // package main

// // import (
// // 	"fmt"
// // 	"io/ioutil"
// // )

// // func main() {
// // 	data, err := ioutil.ReadFile("example.txt")
// // 	if err != nil {
// // 		fmt.Println("Error reading file:", err)
// // 		return
// // 	}
// // 	fmt.Println("File content:")
// // 	fmt.Println(string(data))
// // }

// /Update data
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	newData := []byte("Updated content!\n")
// 	err := ioutil.WriteFile("example.txt", newData, 0644)
// 	if err != nil {
// 		fmt.Println("Error updating file:", err)
// 		return
// 	}
// 	fmt.Println("File updated successfully!")
// }

// /Delete file
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	err := os.Remove("example.txt")
// 	if err != nil {
// 		fmt.Println("Error deleting file:", err)
// 		return
// 	}
// 	fmt.Println("File deleted successfully!")
// }
