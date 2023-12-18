package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println(languages)

	// loops

	for key, value := range languages {
		fmt.Print(key, " ", value)
	}

	cars := map[string]int{
		"BMW":  2008,
		"Audi": 2021,
		"jeep": 2023,
	}

	for key, value := range cars {
		fmt.Println(key, " ", value)
	}
	fmt.Println()
	students := make(map[int]string)
	students[11222631] = "Saloni"
	students[11222531] = "Suhani"

	delete(students, 11222631)

	for key, value := range students {
		fmt.Println(key, value)
	}
	_, found := students[11222531]
	if !found {
		fmt.Println("not found")
	} else {
		fmt.Println("Student found")
	}
}
