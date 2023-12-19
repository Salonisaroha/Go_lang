package main

import "fmt"

func produce(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// using a channel to communicate between a producer and a consumer.

func consume(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
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
}
