package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	// special variable that is use to exchange information among goroutines.

	ch <- "this is channel 1"

	ch <- "this is channel 1"

	ch <- "this is channel 1"
	fmt.Println(ch)
	fmt.Println(ch)
	fmt.Println(ch)
}
