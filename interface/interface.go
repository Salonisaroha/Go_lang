package main

import "fmt"

// The key principle of an interface in Go is to provide method signature for similar types of objects.

type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Structure
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// Implementing method
// of the interface 1
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)

}

// Implementing method
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

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

	// Assigning values
	// to the structure
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the method
	// of the interface 1
	var i1 AuthorDetails = values
	i1.details()

	// Accessing the method
	// of the interface 2
	var i2 AuthorArticles = values
	i2.articles()
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
