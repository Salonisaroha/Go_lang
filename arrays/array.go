package main

import "fmt"

func main() {
	fruit := [5]string{"apple", "banana", "mango", "pineApple", "orange"}
	fmt.Println(fruit)

	vegetableList := [...]string{"potato", "capsicum", "brinjal", "tomato"}

	vegetableList[2] = "spinach"
	fmt.Println("length of vegetableList is ", len(vegetableList))
	fmt.Println("vegetable includes the following veges :- ")
	for i := 0; i < len(vegetableList); i++ {
		fmt.Println(vegetableList[i], " ")
	}

	var arr = [4]int{10, 20, 30, 40}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}
