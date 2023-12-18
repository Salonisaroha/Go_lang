package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Prompt user for age input
	fmt.Print("Enter your age: ")

	// Read user input
	var input string
	fmt.Scanln(&input)

	// Convert input to integer
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid age.")
		os.Exit(1)
	}

	// Check eligibility to vote
	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote yet.")
	}
}
