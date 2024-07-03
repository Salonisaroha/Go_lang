package main

import (
	"bytes"
	"fmt"

	"sort"
)

func main() {

	var fruitList = []string{"Apple", "berries", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fruitList = append(fruitList, "mango", "banana", "litchi")
	fmt.Println(fruitList)

	highScores := make([]int, 4) // here make is used for dyanmic allocation
	highScores[0] = 34
	highScores[1] = 56
	highScores[2] = 67
	highScores[3] = 89

	highScores = append(highScores, 43, 56, 78, 90, 123)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove value from slice by using index

	var courses = []string{"java", "ruby", "python", "goLang"}
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	arr1 := []int{1, 2, 3, 4, 5}
	slc := arr1[0:4]
	fmt.Println("value of arr1 are :- ", arr1)
	fmt.Println("value of slc are :-", slc)
	slc[0] = 100
	slc[1] = 101
	slc[2] = 102
	slc[3] = 104
	fmt.Println("value of arr1 are :-", arr1)
	fmt.Println("value of slice are :-", slc)

	// composite literals :- These are those literals in which elements are present along with the decalration of slices.
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s1)

	// sorting the element

	element := []int{23, 87, 45, 11, 10, 65, 29, 101}
	fmt.Println("element before sorting are :-", element)

	sort.Ints(element)
	fmt.Println("elements after soritng are :-", element)

	//triming of slices
	valuesOfSlice := []byte{'s', 'a', 'l', 'o', 'n', 'i'}
	for i := 0; i < len(valuesOfSlice); i++ {
		fmt.Println(valuesOfSlice[i])
	}

	fmt.Printf("original characters are :- %s", valuesOfSlice)

	//triming
	response := bytes.Trim(valuesOfSlice, "!s")
	fmt.Printf("\n slice with new value are :- %s", response)

	my_arr1 := [10]string{"My", "name", "is", "Saloni Saroha", "I", "doing", "B.tech"}

	fmt.Println("my_arr1 is :- ", my_arr1)

	slice1 := my_arr1[1:6]

	fmt.Println("My slice is :- ", slice1)

	fmt.Printf("Length of slice is :- %d ", len(slice1))
	fmt.Println("Capacity of slice is :- ", cap(slice1))

	arr := [4]int{1, 2, 3, 4}

	my_slice := arr[0:]
	my_slice1 := arr[1:2]
	my_slice2 := arr[1:]

	fmt.Println("value of my_slice is :- ", my_slice)
	fmt.Println("value of my_slice1 is :- ", my_slice1)
	fmt.Println("value of my_slice2 is :- ", my_slice2)

}
