package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "wednesday", "thursday", "friday", "saturday"}
	for i := 0; i < len(days); i++ {
		if days[i] == "Tuesday" {
			continue
		} else {
			fmt.Println(days[i])
		}

	}

	for index, days := range days {
		fmt.Printf("index is %v and value is %v\n", index, days)
	}
}
