package main

import (
	"fmt"
)
// Init function always execute before main function
func init() {
	fmt.Println("Welcome to init function")
}
func main() {
	fmt.Println("Welcome to main function")

}
func init() {
	fmt.Println("Welcome to go lang")
}
