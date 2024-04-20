package main

import (
	"fmt"
)

type Cloneable interface {
	Clone() Cloneable
}

type ConcretePrototype struct {
	data int
}


func NewConcretePrototype(data int) *ConcretePrototype {
	return &ConcretePrototype{data}
}


func (cp *ConcretePrototype) Clone() Cloneable {
	return &ConcretePrototype{cp.data}
}

func main() {

	original := NewConcretePrototype(10)
	fmt.Println("Original data:", original.data)

	clone := original.Clone().(*ConcretePrototype)
	fmt.Println("Clone data:", clone.data)

	
	original.data = 20
	fmt.Println("Original data after modification:", original.data)
	fmt.Println("Clone data:", clone.data)
}
