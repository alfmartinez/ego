package state

import (
	"ego/pkg/terrain"
	"log"
)

type moveState struct {
}

func (s *moveState) Enter() {
	log.Print("Entering Idle State")
}

func (s *moveState) Update(a Actor, g terrain.Grid) State {
	return nil
}
