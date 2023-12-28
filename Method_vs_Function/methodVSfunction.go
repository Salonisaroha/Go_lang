package main

import "fmt"

type data int
type user struct {
	name  string
	email string
}

//func(receiver_name type) method_name(parameter_list type)(return type){code}
func (v1 data) div(v2 data) data {
	return v1 / v2
}

func main() {
	value1 := data(23)
	value2 := data(20)
	res := value1.div(value2)
	fmt.Println(res)

	//function

	res1 := user{
		name:  "saloni saroha",
		email: "abc@gmail.com",
	}
	fmt.Println("user's name :", res1.name)
	fmt.Println("user's email ;", res1.email)

	//P := &res1
	//P.correctEmail("xyz@gmail.com")

	//fmt.Println(res1.name)
	//fmt.Println(res1.email)

	s := simple()
	fmt.Println(s(70, 7))
}
func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}
