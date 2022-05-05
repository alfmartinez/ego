package game

import (
	"log"
	"time"

	"github.com/alfmartinez/ego/examples/plumber/event"
)

type Game interface {
	Start()
}

func New() Game {
	return &game{}
}

type game struct{}

// Start implements Game
func (g *game) Start() {

	c, notify := event.Register()

	stop := make(chan bool)
	go g.Observe(c, stop)

	go func() {
		time.Sleep(time.Second)
		notify <- event.CreateEvent(event.StopEvent)
	}()
	log.Println("Game starting")
	notify <- event.CreateEvent(event.StartEvent)

loop:
	for {
		select {
		case <-stop:
			log.Println("Game shutting off")
			break loop
		}
	}
}

func (*game) Observe(c chan event.Event, stop chan bool) {
	for e := range c {
		switch e.Type() {
		case event.StopEvent:
			stop <- true
		}
	}
}
