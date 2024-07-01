package main

import "fmt"

type Students struct {
	name   string
	rollNo int
	Email  string
}
type Teacher struct {
	name       string
	salary     int
	experience int
}

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}
type data int

type employee struct {
	name   string
	rollNo int
	group  string
}

func (e *employee) bio(chngGroup string) {
	(*e).group = chngGroup
}

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func (a author) show() {
	fmt.Println("name of author is :- ", a.name)
	fmt.Println("branch of author is :-", a.branch)
	fmt.Println("number of particles are :- ", a.particles)
	fmt.Println("salary of author is :- ", a.salary)
}
func (t Teacher) response() {
	fmt.Println("Teacher name is ", t.name)

}
func main() {
	res := author{
		name:      "Saloni Saroha",
		branch:    "CSE",
		particles: 203,
		salary:    100000,
	}
	res.show()
	bio := Teacher{
		name:       "Suhani",
		salary:     12000,
		experience: 2,
	}
	bio.response()

	// non struct type

	value1 := data(30)
	value2 := data(20)
	mul := value1.multiply(value2)
	fmt.Println(mul)

	// Methods with pointer receiver
	info := employee{
		name:   "Huma",
		rollNo: 11222,
		group:  "D",
	}
	fmt.Println(info.name)
	fmt.Println(info.rollNo)
	fmt.Println(info.group)

	newInfo := &info

	newInfo.bio("A")
	fmt.Println(info.group)

}
