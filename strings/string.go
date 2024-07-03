package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Creating string from a slice

	mySlice := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	myString := string(mySlice)
	fmt.Println(myString)

	myslice2 := []rune{0x0047, 0x0065, 0x0065,
		0x006b, 0x0073}

	mystring2 := string(myslice2)
	fmt.Println("String 2: ", mystring2)
	// Ist method of finding length of string
	string1 := "Coding is fun"
	res := len(string1)
	fmt.Println(res)

	// 2nd method for finding length

	res1 := utf8.RuneCountInString(string1)

	fmt.Println(res1)

	// Trimming strings

	details := "@@@@Bangalore is known for its IT Companies!!!!"

	result := strings.Trim(details, "@!")
	fmt.Println(result)

	right := strings.TrimRight(details, "!")
	fmt.Println(right)

	left := strings.TrimLeft(details, "@")

	fmt.Println(left)

	student := "        Students must do hard work to get their dream job"
	fmt.Println(student)
	seq := strings.TrimSpace(student) // Trimspace function is used to remove trailing and leading spaces of strings.
	fmt.Println(seq)

	line1 := "It is to remove leading space"
	line1_res := strings.TrimSuffix(line1, "space")

	fmt.Println(line1_res)

	line2 := "It is to remove trailing spaces"

	line2_res := strings.TrimPrefix(line2, "It")

	fmt.Println(line2_res)

	str1 := "Welcome, to the, online portal, of Coding"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	res4 := strings.Split(str1, ",")
	res2 := strings.Split(str2, "")
	res3 := strings.Split(str3, "!")

	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
}
