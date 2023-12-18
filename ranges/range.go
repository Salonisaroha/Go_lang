package main

import "fmt"

func main() {
	slice := []string{"Hello", "I", "am", "saloni", "saroha"}
	for i, element := range slice {
		fmt.Println("I:  ", i, "Element:", element)
		for _, ch := range element {
			fmt.Printf("%q\n", ch)
		}
	}
	cars := map[string]int{
		"bmw":      2034567,
		"alto":     456788,
		"mercedes": 3487507542,
	}
	for key, value := range cars {
		fmt.Println(key, value)
	}
	myMap := map[string]string{
		"name": "Saloni",
		"age":  "19",
	}

	for key, val := range myMap {
		fmt.Println("Key : ", key, "Value", val)
	}
	fmt.Println(myMap["age"])
	myMap["age"] = "24"
	fmt.Println(myMap["age"])
	delete(myMap, "age")
	fmt.Println(myMap)

	bikes := map[string]int{
		"honda":    2300,
		"splendor": 4500,
		"ct100":    2500,
	}
	for key, value := range bikes {
		fmt.Println("key : ", key, "value : ", value)
	}
	fmt.Println()
	students := make(map[int]string)
	students[11222631] = "Saloni"
	students[11222531] = "susi"
	for key, value := range students {
		fmt.Println(key, value)
	}
	delete(students, 11222631)

	_, found := students[11222631]
	if !found {
		fmt.Println("not found")
	} else {
		fmt.Println("student found")
	}

}
