package state

import (
	"ego/pkg/terrain"
	"log"
)

type exploreState struct {
}

func (s *exploreState) Enter() {
	log.Print("Entering Explore State")
}

func (s exploreState) Update(a *StateMachine, g terrain.Grid) State {

	return nil
}
