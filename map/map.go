package main

import "fmt"

type count struct {
	found int
}

func (f *count) isFound() {
	f.found++
}

func main() {

	fn := count{
		found: 0,
	}
	fmt.Println(fn.found)

	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	fn.isFound()
	fmt.Println(fn.found)
	a := 10
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b++
	fmt.Println(*b)
	fmt.Println(a)

}
