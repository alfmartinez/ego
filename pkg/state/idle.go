package state

import (
	"ego/pkg/terrain"
	"log"
)

type idleState struct {
}

func (s *idleState) Enter() {
	log.Print("Entering Idle State")
}

func (s *idleState) Update(a Actor, g terrain.Grid) State {
	log.Print("Idle state")
	return CreateState("explore")
}
