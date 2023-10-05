// /Create and add in json file
// package main

// import (
// 	"encoding/json"
// 	"os"
// )

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	person := Person{Name: "John", Age: 30}

// 	file, err := os.Create("example.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	err = encoder.Encode(person)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //Read from json
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	file, err := os.Open("example.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	var person Person
// 	err = decoder.Decode(&person)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
// }

///
