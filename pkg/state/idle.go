package state

import (
	"ego/pkg/terrain"
	"log"
)

type idleState struct {
}

func (s *idleState) Enter() {

}

func (s *idleState) Update(a Actor, g terrain.Grid) State {
	log.Print("Idle state")
	return CreateState("explore")
}
