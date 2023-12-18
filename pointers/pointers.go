package main

import "fmt"

func main() {
	number := 23
	var ptr *int = &number
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 2
	fmt.Println("new value is ", number)
}
