package main

import (
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
}
