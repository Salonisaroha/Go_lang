package main

import (
	"fmt"
	"time"
)

func myfun(chan1 chan string) {
	for v := 0; v < 4; v++ {
		chan1 <- "This is my channel"
	}
	close(chan1)
}

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

// A channel with a single "arrow" pointing towards it (e.g. <-chan int) is a receive-only channel, and can only be used to receive values. A channel with a single "arrow" pointing away from it (e.g. chan<- int) is a send-only channel
func main() {
	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}

	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
	//ch := make(chan string )
	msgCh := make(chan string, 128)
	msgCh <- "A"
	msgCh <- "B"
	msgCh <- "C"
	msgCh <- "D"
	close(msgCh)

	// This piece of code is our consumer
	// Ist way
	fmt.Println(<-msgCh)
	for msg := range msgCh {
		fmt.Println(msg)
	}

	// Another way that is second way
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
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout.External service 1 is not respond.")
	}

	var my_chan chan int

	fmt.Println("The value of my_chan is", my_chan)
	fmt.Printf("Type of my_chan is :- %T", my_chan)
	//close(my_chan)

	mychan := make(chan string)
	//mychan <- "Hello from coding world"
	fmt.Println("value of mychan is :- ", mychan)
	fmt.Printf("Type of my chan is %T  ", mychan)

	// Using forloop

	mychnl := make(chan string)
	// Anonymous goroutine
	go func() {
		mychnl <- "go"
		mychnl <- "C++"
		mychnl <- "java"
		mychnl <- "python"
		close(mychnl)
	}()

	for res := range mychnl {
		fmt.Println(res)
	}

	// Only for receiving
	mychanl1 := make(<-chan string)

	// Only for sending
	mychanl2 := make(chan<- string)

	// Display the types of channels
	fmt.Printf("%T", mychanl1)
	fmt.Printf("\n%T", mychanl2)

}
