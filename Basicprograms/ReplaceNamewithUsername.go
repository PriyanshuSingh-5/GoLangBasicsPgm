package main

import (
	"fmt"
	"strings"
)

func main() {
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	// Check if userName has at least 3 characters
	if len(userName) < 3 {
		fmt.Println("User name must have a minimum of 3 characters.")
		return
	}

	// String template
	template := "Hello <<UserName>>, How are you?"

	// Replace <<UserName>> with the user's name
	result := strings.Replace(template, "<<UserName>>", userName, 1)

	fmt.Println(result)
}
