package main

import (
	"fmt"
	"sync"
	"time"
)

func helloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello world")
}
func goodBye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("GoodBye!")
}

func display(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func emp(val int) {
	fmt.Println(val)
}
func main() {
	// Goroutine is a light weighted thread.All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated. Goroutine always works in the background.
	go display("Welcome")
	display("Awesome coding language")

	//Anonymous Go routine
	fmt.Println("Welcome to main function")
	go func() {
		fmt.Println("Welcome to anonymous go routine")
	}()
	time.Sleep(1 * time.Second)

	go emp(12)
	time.Sleep(2 * time.Second)
	go emp(13)
	time.Sleep(2 * time.Second)

	// In the above examples main function is block for sometime for executing the goroutine suceesfully but if we have huge lines of code and millions of concurrent requests become stop.To avoid this problem we use wait groups

	var wg sync.WaitGroup
	wg.Add(2)
	helloWorld(&wg)
	goodBye(&wg)
	wg.Wait()
}
