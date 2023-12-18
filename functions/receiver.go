package main

import (
	"fmt"
)

func main() {
	// fmt.Println("welcome to function")
	// value := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your country name :- ")
	// input, _ := value.ReadString('\n')
	// if input == "India" {
	// 	fmt.Println("Namste India")
	// } else {
	// 	fmt.Println("Good Morning")
	// }

	var age int
	fmt.Println("Enter your age :-")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("you cannot give vote ")
	} 
	

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
	

	var name string
	fmt.Println("Enter your first name :- ")
	fmt.Scanln(&name)
	var lastName string
	fmt.Println("Enter your lastName :- ")
	fmt.Scanln(&lastName)

	fmt.Println("Your full Nmae is " + name + " " + lastName)
}
// func greeting(greet string, name string)string{
// var country string
// fmt.Println("Entyer your country name :- ")
// fmt.Scanln(&country)
// if country=="india"{
// 	fmt.Println("namste"+ " "+ name)
// } else{
// 	fmt.Println("Good Morning "+ " "+ name)
// }

// }
