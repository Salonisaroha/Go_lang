package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannel := make(chan string, 3)

	bufferedChannel <- "First Message"
	bufferedChannel <- "Second Message"
	bufferedChannel <- "Third Message"

	time.Sleep(time.Second * 3)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}
