package main

import "fmt"

func produce(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// using a channel to communicate between a producer and a consumer.
// When we are working with multiple goroutines we use channels.

func consume(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func sending(s chan string) {
	s <- "go lang is intresting."
}
func receiving(s chan string) {
	fmt.Println(<-s)
}
func convert(s chan<- string) {
	s <- "some values"
}
func main() {
	ch := make(chan int)
	go produce(ch)
	consume(ch)

	//Another channel

	data := make(chan string)
	go func() {
		data <- "hey champ"
	}()
	c := <-data
	fmt.Println(c)
	// Bidirectional channel and unidirectional channel

	// Creating a bidirectional channel

	chanl1 := make(chan string)
	chanl2 := make(chan string)

	//chanl1:= make(chan<- string) This is only send only channel.
	//chanl1 := make(<-chan string) This is receive only channel.

	go sending(chanl1)
	valueFromChanel := <-chanl1
	fmt.Println("valueFromChanel", valueFromChanel)
	go receiving(chanl2)
	chanl2 <- valueFromChanel

	// Creating a unidirectional channel

	chanl := make(chan string)
	go convert(chanl)
	fmt.Println(<-chanl)
}
