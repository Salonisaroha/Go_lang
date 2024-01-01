package main

import (
	"fmt"
	"time"
)

func main() {
	green := make(chan string)
	yellow := make(chan string)
	red := make(chan string)

	go green_func(green)
	go yellow_func(green)
	go red_func(green)

	fmt.Println("waiting for messages...Blocked here")
	for {
		select {
		case msg := <-green:
			fmt.Println(msg)

		case msg := <-yellow:
			fmt.Println(msg)

		case msg := <-red:
			fmt.Println(msg)
		}
	}
}
func green_func(green chan string) {
	for {
		time.Sleep(1 * time.Second)
		green <- "Green message"
	}
}
func yellow_func(green chan string) {
	for {
		time.Sleep(2 * time.Second)
		green <- "yellow message"
	}
}
func red_func(green chan string) {
	for {
		time.Sleep(3 * time.Second)
		green <- "red message"
	}
}
