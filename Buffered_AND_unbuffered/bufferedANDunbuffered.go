package main

import (
	"fmt"
	"time"
)

// A channel with a single "arrow" pointing towards it (e.g. <-chan int) is a receive-only channel, and can only be used to receive values. A channel with a single "arrow" pointing away from it (e.g. chan<- int) is a send-only channel
func main() {
	bufferedChannel := make(chan string, 3)

	bufferedChannel <- "First Message"
	bufferedChannel <- "Second Message"
	bufferedChannel <- "Third Message"

	time.Sleep(time.Second * 3)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	externalService1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		externalService1 <- "Data received from service 1"
	}()
	select {
	case result := <-externalService1:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout.External service 1 is not respond.")
	}
}
