package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // it takes 0 bytes
	msgCh  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgCh:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

func (s *Server) loop() {
	for {
		select {
		case <-s.quitch:
			// do some stuff when we need to quit
			fmt.Println("Quiting server")
		case msg := <-s.msgCh:
			// do some stuff when we have a message
			s.handleMessage(msg)
		default:
			fmt.Println("nothing happened")
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("we received a message :", msg)
}
func (s *Server) sendMessage(msg string) {
	s.msgCh <- msg
}
func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func main() {
	server := newServer()
	go func() {
		time.Sleep(time.Millisecond * 5)
		server.quit()

	}()
	go server.start()
	for i := 0; i < 100; i++ {
		server.sendMessage(fmt.Sprintf("handle this number %d", i))
	}

	time.Sleep(time.Millisecond * 2)
}
