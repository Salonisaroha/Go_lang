package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	var slice = make([]int, 4, 10)
	fmt.Printf("My slice is :- %v\n Length of slice is %d\n Capacity of slice is %d\n", slice, len(slice), cap(slice))
	my_slice := []string{"Go", "Language", "is", "most", "popular"}

	for index, val := range my_slice {
		fmt.Printf("Index is %d and value is %s\n", index, val)
	}
	for _, ele := range my_slice {
		fmt.Printf("Element = %s\n", ele)
	}

	// Slice is a refernced array, if some changes made int the slice then it will also reflect in the main array because slice is refernced type of it.

	// Multidimensional slice

	slice1 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(slice1)

	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice1); j++ {
			fmt.Println(slice1[i][j])
		}
	}

	// Sorting by using sort library

	sl1 := []string{"Coding", "is", "awesome"}
	sl2 := []int{30, 80, 100, 12, 14, 18}

	fmt.Println("Before sorting elements are :-")
	fmt.Println(sl1)
	fmt.Println(sl2)
	sort.Strings(sl1)
	sort.Ints(sl2)
	fmt.Println("After sorting elements are :- ")

	fmt.Println(sl1)
	fmt.Println(sl2)

	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	res3 := bytes.Trim([]byte("^^Geek&&"), "$")

	// Display the results
	fmt.Printf("Final Slice:\n")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)

	slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's',
		'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}

	slice_2 := []byte{'A', 'p', 'p', 'l', 'e'}

	slice_3 := []byte{'%', 'g', '%', 'e', '%', 'e',
		'%', 'k', '%', 's', '%'}

	res_1 := bytes.Split(slice_1, []byte("eek"))
	res_2 := bytes.Split(slice_2, []byte(""))
	res_3 := bytes.Split(slice_3, []byte("%"))

	// Display the results
	fmt.Printf("\n\nAfter splitting:")
	fmt.Printf("\nSlice 1: %s", res_1)
	fmt.Printf("\nSlice 2: %s", res_2)
	fmt.Printf("\nSlice 3: %s", res_3)
}
