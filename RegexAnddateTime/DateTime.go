///print current date
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // Get the current date and time
//     currentTime := time.Now()

//     // Format and print the current date and time
//     currentDateTime := currentTime.Format("2006-01-02 15:04:05")
//     fmt.Println("Current Date and Time:", currentDateTime)
// }

///print next date
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // Get the current date
//     currentDate := time.Now()

//     // Calculate the next date by adding one day
//     nextDate := currentDate.AddDate(0, 0, 1)

//     // Format and print the next date
//     nextDateString := nextDate.Format("2006-01-02")
//     fmt.Println("Next Date:", nextDateString)
// }

///calculate age using birthdate and currentdate
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     // Input birthdate
//     birthdate := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)

//     // Calculate age
//     currentYear := time.Now().Year()
//     birthYear := birthdate.Year()
//     age := currentYear - birthYear

//     fmt.Printf("Age: %d\n", age)
// }

///calculate days btw 2 dates
// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     startDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
//     endDate := time.Date(2022, time.February, 15, 0, 0, 0, 0, time.UTC)

//     daysBetween := calculateDaysBetween(startDate, endDate)
//     fmt.Printf("Days between: %d\n", daysBetween)
// }

// func calculateDaysBetween(startDate, endDate time.Time) int {
//     duration := endDate.Sub(startDate)
//     days := int(duration.Hours() / 24)
//     return days
// }

///
