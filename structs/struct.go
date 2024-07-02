package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type student struct {
	name   string
	age    int
	group  string
	course string
}
type department struct {
	dept_name string
	courses   string
	location  string
	dId       int
}

type employee struct {
	name    string
	eid     int
	details department
}

func main() {
	details := User{
		"Saloni", "salonisaroha5@gmail.com", true, 19,
	}
	fmt.Println(details)
	fmt.Printf("Details are : %+v\n", details)

	bio := &student{
		"Saloni Saroha",
		20,
		"D",
		"CSE",
	}

	fmt.Println((*bio).name)
	fmt.Println((*bio).age)
	fmt.Println((*bio).group)
	fmt.Println((*bio).course)
	// without dereference operator
	fmt.Println((bio).name)
	fmt.Println((bio).age)
	fmt.Println((bio).group)
	fmt.Println((bio).course)

	// Nested structures

	res := employee{
		name: "John Doe",
		eid:  1678934,
		details: department{
			"ECE",
			"Mechanical",
			"Block 1",
			123456,
		},
	}
	fmt.Println(res.name)
	fmt.Println(res.eid)
	fmt.Println(res.details.courses)

}
