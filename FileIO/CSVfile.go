package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name string
	Age  string
}

func main() {
	// Create a CSV file
	createCSV()

	// Read data from the CSV file
	readCSV()

	// Update data in the CSV file
	updateCSV()

	// Read the updated data
	readCSV()

	// Delete the CSV file
	//deleteCSV()
}

func createCSV() {
	// Create or open the CSV file for writing
	file, err := os.Create("example.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Define data
	records := []Record{
		{"Alice", "25"},
		{"Bob", "30"},
		{"Charlie", "22"},
	}

	// Write data to the CSV file
	for _, record := range records {
		err := writer.Write([]string{record.Name, record.Age})
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}

	fmt.Println("CSV file created successfully!")
}

func readCSV() {
	// Open the CSV file for reading
	file, err := os.Open("example.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read data from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Print the data
	fmt.Println("CSV file content:")
	for _, record := range records {
		fmt.Println("Name:", record[0], "Age:", record[1])
	}
}

func updateCSV() {
	// Open the CSV file for reading and writing
	file, err := os.OpenFile("example.csv", os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Update the data
	records[1][1] = "31" // Update Bob's age to 31

	// Create a CSV writer
	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}

	fmt.Println("CSV file updated successfully!")
}

func deleteCSV() {
	err := os.Remove("example.csv")
	if err != nil {
		fmt.Println("Error deleting CSV file:", err)
		return
	}

	fmt.Println("CSV file deleted successfully!")
}
