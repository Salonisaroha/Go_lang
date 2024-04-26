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
	for i := 0; i < 5; i++ {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 500)
	}
}

func hello(val *sync.WaitGroup) {

	for i := 0; i < 4; i++ {
		defer val.Done()
		fmt.Println("Hello everyone")
	}
}
func goodBye(val *sync.WaitGroup) {
	defer val.Done()
	fmt.Println("Goodbye")
}

func fetchResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}
func main() {

	// Anonymous function
	go func() {
		fetchResource() // It does not return any value . So, here anonymous function is used.
	}()
	result := fetchResource()
	fmt.Println(result)
	go fetchResource()

	var val sync.WaitGroup
	val.Add(1)
	go hello(&val)
	go goodBye(&val)

	val.Wait()
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
