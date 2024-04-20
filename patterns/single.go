package main

import "fmt"


type Singleton struct {
	data string
}


var instance *Singleton


func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{"singleton data"}
	}
	return instance
}

func main() {
	singleton1 := GetInstance()
	fmt.Println("Singleton 1 data:", singleton1.data)

	singleton1.data = "modified data"

	
	singleton2 := GetInstance()
	fmt.Println("Singleton 2 data:", singleton2.data)
}
