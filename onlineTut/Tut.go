package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// fmt.Println("hello go lang")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("enter your name")
	// value, _ := reader.ReadString('\n')
	// fmt.Println("name of client is ", value)
	fmt.Println("Welcome to our pizza app")
	fmt.Println("please rate our app between 1 and 5 ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your rating ,", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating : ", numRating+1)
	}

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
   // memory management
   
}
