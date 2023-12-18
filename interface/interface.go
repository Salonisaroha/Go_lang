package main

import "fmt"

// The key principle of an interface in Go is to provide method signature for similar types of objects.

type Shapes interface {
	area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	length  float64
	breadth float64
}

func (c Circle) area() float64 {
	return 3.14 * (c.radius * c.radius)
}
func (r Rectangle) area() float64 {
	return (r.length * r.breadth)
}

// salary calculation using interface
type Employee interface {
	CalculateSalary() int
}
type Permanent struct {
	empId   string
	basePay int
	pf      int
}
type Contract struct {
	empId   string
	basePay int
}

func (p Permanent) CalculateSalary() int {
	return p.basePay + p.pf
}
func (c Contract) CalculateSalary() int {
	return c.basePay
}
func totlSalary(employees []Employee) int {
	totalExpense := 0
	for _, person := range employees {
		totalExpense = totalExpense + person.CalculateSalary()
	}
	return totalExpense
}
func main() {
	joy := Permanent{
		empId:   "jljhds",
		basePay: 45000,
		pf:      3000,
	}
	julia := Permanent{
		empId:   "4387",
		basePay: 34000,
		pf:      3499,
	}
	ramesh := Contract{
		empId:   "ramesh babu",
		basePay: 12000,
	}
	employees := []Employee{joy, julia, ramesh}

	total := totlSalary(employees)
	fmt.Println(total)
}
