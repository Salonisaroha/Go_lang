package main

import (
	"fmt"
	"time"
)

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

// A channel with a single "arrow" pointing towards it (e.g. <-chan int) is a receive-only channel, and can only be used to receive values. A channel with a single "arrow" pointing away from it (e.g. chan<- int) is a send-only channel
func main() {
	//ch := make(chan string )
	msgCh := make(chan string, 128)
	msgCh <- "A"
	msgCh <- "B"
	msgCh <- "C"
	msgCh <- "D"
	close(msgCh)

	// This piece of code is our consumer
	// fmt.Println(<-msgCh)
	// for msg := range msgCh {
	// 	fmt.Println(msg)
	// }

	// Another way
	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}
		fmt.Println("The message -> ", msg)
	}

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
