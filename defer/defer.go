package main

import "fmt"

func printData() {
	defer fmt.Println("Hello World!")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// defer uses last in first out approach.
func main() {
	defer fmt.Println("world")
	defer fmt.Println("today I am learning goLang")
	fmt.Println("Hello !")
	fmt.Println("Printing...")
	printData()
	fmt.Println("Printing Done!")
	myDefer()

}
