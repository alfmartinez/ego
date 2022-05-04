package game

import (
	"log"
	"time"
)

type Game interface {
	Start()
}

func New() Game {
	return &game{}
}

type game struct{}

// Start implements Game
func (*game) Start() {
	stop := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		stop <- true
	}()
	log.Println("Game starting")

loop:
	for {
		select {
		case <-stop:
			log.Println("Game shutting off")
			break loop
		}
	}
}
