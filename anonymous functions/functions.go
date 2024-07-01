package main

import (
	"fmt"
)

func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))

}
func main() {
	// Creating an anonymous function.
	func() {
		fmt.Println("Hello from coding world!")
	}()

	// assigning anonymous functions to a variable

	value := func() {
		fmt.Println("Welcome to coding world! ")
	}
	value()

	// passing parameters to anonymous functions

	func(ele string) {
		fmt.Println(ele)
	}("Go language is at BOOM")

	// use of anonymous function as an arguement in another function

	value1 := func(p, q string) string {
		return p + q + "Geeks"
	}
	GFG(value1)
}
