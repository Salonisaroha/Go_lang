package main
import "fmt"
type Students struct{
	name string
	rollNo int
	Email string
}
type Teacher struct{
	name string
	salary int
	experience int
}
func (t Teacher) response(){
fmt.Println("Teacher name is ", t.name)

}
func main(){

}