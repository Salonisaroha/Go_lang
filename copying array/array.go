package main

import (
	"fmt"
)

func printArray(array []int) {
	for _, value := range array {
		fmt.Println(value)
	}
}
func main() {
	//pass by value
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arr1 := arr
	fmt.Println("Before changing in original array")
	fmt.Println(arr)
	fmt.Println(arr1)
	arr[0] = 10
	fmt.Println("After changing original array we get only change in original array not in copied array")
	fmt.Println(arr)
	fmt.Println(arr1)

	// pass by refernce

	my_arr1 := [5]string{"C++", "Scala", "Python", "Golang", "Java"}
	my_arr2 := &my_arr1
	fmt.Println(my_arr1)
	fmt.Println(*my_arr2)

	my_arr1[0] = "Perl"

	fmt.Println(my_arr1)
	fmt.Println(*my_arr2)

	// If you want to copy array into another array then there are three methods for doing the same

	// Ist method is using for loop

	original_array := []int{10, 20, 30, 40, 50}
	copy_array := make([]int, len(original_array))
	for i, value := range original_array {
		copy_array[i] = value
	}
	fmt.Println("Original array is :- ", original_array)
	fmt.Println("Copied array is :- ", copy_array)

	// using function

	originalArray1 := []int{1, 2, 3, 4, 5}
	copyArray1 := make([]int, len(originalArray1))

	copy(copyArray1, originalArray1)

	fmt.Println("Original Array: ", originalArray1)
	fmt.Println("Copy Array: ", copyArray1)

	// By using append function

	originalArray2 := []int{1, 2, 3, 4, 5}
	copyArray2 := make([]int, 0, len(originalArray2))

	copyArray2 = append(copyArray2, originalArray2...)

	fmt.Println("Original Array: ", originalArray2)
	fmt.Println("Copy Array: ", copyArray2)

	array := []int{1, 2, 3, 4, 5}
	printArray(array)
}
