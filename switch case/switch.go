package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var day string
	fmt.Println("Enter day")
	fmt.Scanln(&day)
	switch day {
	case "Monday":
		fmt.Println("Today is monday")
	case "Tuesday":
		fmt.Println("Today is tuesday")
	case "Wednesday":
		fmt.Println("Today is wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")

	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	}

	// dice rollling.
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move second spot")
	case 3:
		fmt.Println("you can move to 3 spot")
	case 4:
		fmt.Println("you can move to 4th spot")
	case 5:
		fmt.Println("you can move to 5th spot")
	case 6:
		fmt.Println("you can move to 6 spot and roll dice again")
	default:
		fmt.Println("See dice again")
	}
}
