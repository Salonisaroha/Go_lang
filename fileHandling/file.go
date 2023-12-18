package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file - learnCodeOnline"
	file, err := os.Create("./myfile.txt") // create file 

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is \n ", string(databyte))
}
