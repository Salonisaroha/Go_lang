package main

import (
	"fmt"
)

//anonymous field

type student struct {
	int
	string
	float64
}

func main() {
	//Anonymous structures are those structures which does not have any name and it is one time usable structures.

	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Saloni Saroha",
		branch:    "CSE",
		language:  "Coding languages",
		Particles: 102,
	}
	fmt.Println(Element)

	// Anonymous fields are those which does not contain any field name

	result := student{123, "Suhani Gupta", 123.980}
	fmt.Println(result)
	fmt.Println(result.int)
	fmt.Println(result.string)
	fmt.Println(result.float64)
}
