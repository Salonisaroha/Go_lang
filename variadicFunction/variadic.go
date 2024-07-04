package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	//varidiac function denoted by ...
	// These are those functions in which any no. of arguements can be passed.
	sum(2, 3)
	sum(3, 4, 5, 6, 7)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
