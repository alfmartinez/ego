package game

import (
	"log"
	"time"

	baseScene "github.com/alfmartinez/ego/engine/scene"

	"github.com/alfmartinez/ego/examples/plumber/setup"
)

type Game interface {
	Start()
}

func New() Game {
	return &game{scene: setup.DefaultScene()}
}

type game struct {
	scene baseScene.Scene
}

// Start implements Game
func (g *game) Start() {
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
