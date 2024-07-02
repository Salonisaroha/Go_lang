package main

import (
	"fmt"
	//"reflect"
)

func main() {
	// types of array :-
	// one dimension
	// multi dimension

	// type 1
	//var arr1 [2]int

	// type 2
	//arr2 := [2]string{}

	// arr1[0] = 1
	// arr1[1] = 2
	// arr2[0] = "1"
	// arr2[1] = "2"
	// fmt.Println(arr1, arr2)

	// for i := 0; i < 2; i++ {
	// 	fmt.Println(arr1[i], reflect.TypeOf(arr1[i]))
	// }

	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println(arr2[i], reflect.TypeOf(arr2[i]))
	// }

	// fruit := [5]string{"apple", "banana", "mango", "pineApple", "orange"}
	// fmt.Println(fruit)

	// vegetableList := [...]string{"potato", "capsicum", "brinjal", "tomato"}

	// vegetableList[2] = "spinach"
	// fmt.Println("length of vegetableList is ", len(vegetableList))
	// fmt.Println("vegetable includes the following veges :- ")
	// for i := 0; i < len(vegetableList); i++ {
	// 	fmt.Println(vegetableList[i], " ")
	// }

	// var arr = [4]int{10, 20, 30, 40}

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i], " ")
	// }

	data := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(data)

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	bio := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(bio[i][j], " ")
		}

	}
}
