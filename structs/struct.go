package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	details := User{
		"Saloni", "salonisaroha5@gmail.com", true, 19,
	}
	fmt.Println(details)
	fmt.Printf("Details are : %+v\n", details)
}
