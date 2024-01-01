package main

import (
	"fmt"
	"math/rand"

	"sync"
	"time"
)

//Every process which run concurrently in golang called goroutines.
//light weighted thread.
//creation cost of goroutines is very small as compared to thread.
//Every program had at least single goroutine called main function.

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}
func printSomething(value string) {
	for {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	dataChan := make(chan int)

	go func() { // go-routine that's running on the background thread
		wg := sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() { // go-routine that's running on the background thread
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}

		wg.Wait()
		close(dataChan)
	}()
	for x := range dataChan {
		fmt.Printf("%d\n", x)

	}

	go printSomething("one") // here go routine is used by which we get one and two after one another.If we remove go then only one is print two is not print
	printSomething("two")

}
