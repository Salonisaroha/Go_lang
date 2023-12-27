package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Anonymous functions are those functions which are used only once.
func twoValues(a, b int) (int, int) {
	sum := a + b
	multiply := a * b
	return sum, multiply
}
func addition(num1, num2 int) int {
	res := num1 + num2
	fmt.Println(res)
	return res
}

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

	items := [5]string{"sugar", "tea", "coffee", "cinemon", "elaichi"}
	for key, value := range items {
		fmt.Println(key, value)
	}

	colors := map[string]string{"Red": "#ff00", "blue": "#cf89076", "green": "#wert56789"}
	for color, hex := range colors {
		//fmt.Println("The name of colors is ", color, "and the hexadecimal value of color is ", hex)
		fmt.Println(color, hex)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)

	}
	defer addition(23, 45)
	add, product := twoValues(4, 5)
	fmt.Println(add, product)

	file, err := os.Open("Example.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
