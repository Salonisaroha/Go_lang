package main

import "fmt"

type CustomMap[k comparable, V any] struct {
	data map[k]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}
func NewCustomMap[k comparable, V any]() *CustomMap[k, V] {
	return &CustomMap[k, V]{
		data: make(map[k]V),
	}
}
func foo[T any](val T) {
	fmt.Println(val)
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 1000.890)

	foo("egg")

}
