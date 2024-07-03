package main

import "fmt"

func ptf(a *int) {
	*a = 200
}

type student struct {
	name   string
	rollNo int
}

// type Player struct {
// 	HP int
// }
// type Bigdata struct {
// 	// 1GB of memory
// 	// .
// }
// type Database struct {
// 	user string
// }
// type Server struct {
// 	db *Database
// }

// func (s *Server) GetUserFromDB() string {
// 	if s.db == nil {
// 		panic("database == nil hence, is not initialized")
// 	}
// 	return s.db.user
// }

// // func doSomething(data *Bigdata) {

// // }

// // if player is not a pointer we are adjusting the copy of the player
// // not the actual player
// func TakeDamage(player *Player, amount int) {

// 	player.HP -= amount
// 	fmt.Println("player is taking amount -> ", player.HP)
// }

// // function receiver
// func (p *Player) TakeDamage(amount int) {
// 	p.HP -= amount
// 	fmt.Println("Player is taking damage. New HP -> ", p.HP)
// }
func main() {
	// number := 23
	// var ptr *int = &number
	// fmt.Println(ptr)
	// fmt.Println(*ptr)
	// *ptr = *ptr + 2
	// fmt.Println("new value is ", number)

	// player := &Player{
	// 	HP: 100,
	// }
	// TakeDamage(player, 10)
	// TakeDamage(player, 80)
	// player.TakeDamage(200)
	// fmt.Println(player)

	// data := &Bigdata{}
	// fmt.Println(data)

	// s := &Server{}
	// s.GetUserFromDB()

	var x int = 34
	var p *int = &x // It only stores int type of address, so we generally avoid it.
	fmt.Println("value of p is :- ", p)
	fmt.Println("value of p is :- ", &x)

	q := &x
	fmt.Println("Address of p variable is :- ", q)
	fmt.Println("Value of p variable is :- ", *q)

	*q = 20

	fmt.Println(x)
	fmt.Println(*q)

	var y = 100
	fmt.Printf("The value of y before function call is: %d\n", y)

	var pa *int = &y

	ptf(pa)

	fmt.Printf("The value of y after function call is: %d\n", y)
	ptf(&y)

	fmt.Printf("The value of y after function call is: %d\n", y)

	// Making instance of student variable
	stu := student{"Saloni", 1020547676}

	pts := &stu
	fmt.Println(pts)
	fmt.Println(pts.name)
	fmt.Println(pts.rollNo)
	fmt.Println(*pts)

	//changing values of struct
	pts.name = "Suhani"

	fmt.Println(pts)

	//Double Pointers

	var V int = 100
	var pt1 *int = &V
	var pt2 **int = &pt1
	fmt.Println("The Value of Variable V is = ", V)
	fmt.Println("Address of variable V is = ", &V)

	fmt.Println("The Value of pt1 is = ", pt1)
	fmt.Println("Address of pt1 is = ", &pt1)

	fmt.Println("The value of pt2 is = ", pt2)
	fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)
	fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)
}
