package main

import "fmt"

func Company() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

func main() {

	//A closure is a function value that references variables from outside its body.
	func(s string) {
		fmt.Println(s)
	}("hello") //anonymous function

	val := Company()
	fmt.Println(val()) // output is 1
	fmt.Println(val()) // output is 2
	fmt.Println(val()) // output is 3
	fmt.Println(val()) // output is 4
	fmt.Println(val()) // output is 5

	//Reinitialize the value

	v := Company()
	fmt.Println(v())
	fmt.Println(v())
	fmt.Println(v())

}
