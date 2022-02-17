package state

import (
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type moveState struct {
	position utils.Position
	next     string
}

func (s moveState) Enter() {
	log.Print("Entering Idle State")
}

func (s moveState) Update(a *StateMachine, g terrain.Grid) State {
	return nil
}
