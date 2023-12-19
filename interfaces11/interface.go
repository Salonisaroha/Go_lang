package main

import "fmt"

type MotorVehicle interface {
	mileage() float64
}
type Bmw struct {
	distance     float32
	fuel         float32
	averageSpeed string
}
type Audi struct {
	distance float64
	fuel     float64
}

func (b Bmw) mileage() float64 {
	return float64(b.distance) / float64(b.fuel)
}
func (a Audi) mileage() float64 {
	return a.distance / a.fuel
}
func totalMileage(m []MotorVehicle) {
	tm := 0.0
	for _, v := range m {
		tm = tm + v.mileage()
	}
	fmt.Printf("Total Mileage per month %f km/1", tm)
}

type Employee interface {
	GetName() string
}
type Engineer struct {
	Name string
}
type Finance struct {
	Name string
}

func (e Engineer) GetName() string {
	return "employee name:" + e.Name
}
func (f Finance) GetName() string {
	return "employee name:" + f.Name
}
func EmployeeDetails(e Employee) {
	fmt.Println(e.GetName())
}

func main() {
	b1 := Bmw{
		distance:     167.9,
		fuel:         36,
		averageSpeed: "58",
	}
	a1 := Audi{
		distance: 152.9,
		fuel:     30,
	}

	person := []MotorVehicle{b1, a1}
	totalMileage(person)

	// declare an interface employee.

	engineer := Engineer{
		Name: "Saloni",
	}
	EmployeeDetails(engineer)
	finance := Finance{
		Name: "Saroha",
	}
	EmployeeDetails(finance)
}
