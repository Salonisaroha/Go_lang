package main

import (
	"fmt"
	"sync"
	"time"
)

var accountBalance int
var mtx sync.Mutex
var wait sync.WaitGroup

func transaction(amount int, transactionType string, name string) {
	if transactionType == "CREDIT" {
		accountBalance += amount
		fmt.Printf("$ %d  has been credited by %s\n", amount, name)
	}
	if transactionType == "DEBIT" {
		if (accountBalance - amount) < 0 {
			fmt.Printf("This transaction is not valid. The amount is greater than balance")
		}
	} else {
		mtx.Lock()
		accountBalance = accountBalance - amount
		fmt.Printf("$ %d has been debited by %s\n", amount, name)
		mtx.Unlock()
	}

}
func accountHolder1(name string) {
	defer wait.Done()
	for i := 0; i < 5; i++ {
		transaction(10, "CREDIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}
func accountHolder2(name string) {
	defer wait.Done()
	for i := 0; i < 5; i++ {
		transaction(7, "DEBIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}
func accountHolder3(name string) {
	defer wait.Done()
	for i := 0; i < 5; i++ {
		transaction(11, "CREDIT", name)
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	// Mutex is used for locking the process and releasing the process.
	// For example 3 persons have joint account and balance in their account is 100. One person wants to withdraw 100, another wants to withdraw 70 and third one wants to withdraw 80.But this creates panic. This is possible by using mutex.
	wait.Add(3)
	go accountHolder1("saloni")
	go accountHolder2("riya")
	go accountHolder3("gayu")
	wait.Wait()
	fmt.Printf("rest account balance is %d", accountBalance)
}
