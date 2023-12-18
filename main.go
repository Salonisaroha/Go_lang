package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	rollNo     int
	name       string
	isVerified bool
}

func Calculator(a, b int) string {
	add := a + b
	sub := a - b
	multi := a * b
	divide := a / b

	output := "addition of two numbers is " + fmt.Sprint(add) +
		", subtraction of two numbers is " + fmt.Sprint(sub) +
		", multiplication of two numbers is " + fmt.Sprint(multi) +
		", division of two numbers is " + fmt.Sprint(divide)

	return output
}

func main() {
	var sushant = Student{11222631,
		"saloni saroha", true}
	fmt.Println(sushant.name)

	resultOfCalculator := Calculator(90, 40)
	fmt.Println(resultOfCalculator)

	var array1 = []int{12, 34, 56, 76, 45, 23}
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
	arr := []string{"saloni", "susi", "tanu", "sakshi", "riya"}
	for j := 0; j < len(arr); j++ {
		if arr[j] == "tanu" {
			continue
		}

		fmt.Println(arr[j])

	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

}
