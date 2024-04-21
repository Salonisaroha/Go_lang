package main

import "fmt"

type Player struct {
	HP int
}
type Bigdata struct {
	// 1GB of memory
	// .
}
type Database struct {
	user string
}
type Server struct {
	db *Database
}

func (s *Server) GetUserFromDB() string {
	if s.db == nil {
		panic("database == nil hence, is not initialized")
	}
	return s.db.user
}

// func doSomething(data *Bigdata) {

// }

// if player is not a pointer we are adjusting the copy of the player
// not the actual player
func TakeDamage(player *Player, amount int) {

	player.HP -= amount
	fmt.Println("player is taking amount -> ", player.HP)
}

// function receiver
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("Player is taking damage. New HP -> ", p.HP)
}
func main() {
	number := 23
	var ptr *int = &number
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 2
	fmt.Println("new value is ", number)

	player := &Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	TakeDamage(player, 80)
	player.TakeDamage(200)
	fmt.Println(player)

	data := &Bigdata{}
	fmt.Println(data)

	s := &Server{}
	s.GetUserFromDB()

}
