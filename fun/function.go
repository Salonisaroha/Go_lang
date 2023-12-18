package main

import "fmt"

func greet(a, b int) int {
	return a + b
}
func main() {
	res := greet(23, 45)
	fmt.Println(res)
}
