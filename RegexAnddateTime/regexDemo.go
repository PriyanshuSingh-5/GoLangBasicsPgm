// /validate Email
// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	var email string
// 	fmt.Print("Enter an email address: ")
// 	fmt.Scanln(&email)

// 	isValid := isValidEmail(email)
// 	if isValid {
// 		fmt.Println("Valid email address.")
// 	} else {
// 		fmt.Println("Invalid email address.")
// 	}
// }

// func isValidEmail(email string) bool {
// 	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
// 	regex := regexp.MustCompile(pattern)
// 	return regex.MatchString(email)
// }

///validate phoneNo.
// package main

// import (
//     "fmt"
//     "regexp"
// )

// func main() {
//     var phone string
//     fmt.Print("Enter a phone number (e.g., +1234567890): ")
//     fmt.Scanln(&phone)

//     isValid := isValidPhoneNumber(phone)
//     if isValid {
//         fmt.Println("Valid phone number.")
//     } else {
//         fmt.Println("Invalid phone number.")
//     }
// }

// func isValidPhoneNumber(phone string) bool {
//     pattern := `^\+\d{10,15}$`
//     regex := regexp.MustCompile(pattern)
//     return regex.MatchString(phone)
// }

///Validate username
// package main

// import (
//     "fmt"
//     "regexp"
// )

// func main() {
//     var username string
//     fmt.Print("Enter a username: ")
//     fmt.Scanln(&username)

//     isValid := isValidUsername(username)
//     if isValid {
//         fmt.Println("Valid username.")
//     } else {
//         fmt.Println("Invalid username.")
//     }
// }

// func isValidUsername(username string) bool {
//     pattern := `^[A-Z][a-z]{2,}$`
//     regex := regexp.MustCompile(pattern)
//     return regex.MatchString(username)
// }
