package game

import "log"

type Game interface {
	Start()
}

func New() Game {
	return &game{}
}

type game struct{}

// Start implements Game
func (g *game) Start() {
	stop := make(chan bool)

	go func() {
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
