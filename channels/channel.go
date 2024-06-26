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
func sum(intArray []int, ch chan int) {
	result := 0
	var val int
	for _, val = range intArray {
		result += val
	}
	ch <- result
}
func main() {
	msgCh := make(chan string, 128)
	msgCh <- "A"
	msgCh <- "B"
	msgCh <- "C"
	msgCh <- "D"
	msgCh <- "E"
	msgCh <- "F"
	for err, val := range <-msgCh {
		fmt.Printf("Value of different channels is %v and the value is %v\n", val, err)
	}
	resultCh := make(chan string) //-> unbuffered channel
	// -> buffered channel make(chan string, 10), if size is given then it is buffered channel.

	// A channel is always blocked, if it is full.
	// We can solved it by buffering the channel.

	// Below output is print because we use anonymous function.
	go func() {
		result := <-resultCh
		fmt.Println(result)
	}()
	resultCh <- "Something is happen"
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

	//chanl1 := make(chan<- string)  This is only send only channel.
	//chanl1 := make(<-chan string)  This is receive only channel.

	go sending(chanl1)
	valueFromChanel := <-chanl1
	fmt.Println("valueFromChanel", valueFromChanel)
	go receiving(chanl2)
	chanl2 <- valueFromChanel

	//Creating a unidirectional channel

	chanl := make(chan string)
	go convert(chanl)
	fmt.Println(<-chanl)

	arr := []int{10, 20, 30, 40, 50}
	math := make(chan int)
	go sum(arr, math)
	fmt.Println(<-math)
}
