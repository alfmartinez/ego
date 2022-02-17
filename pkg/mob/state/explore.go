package state

import (
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type exploreState struct {
	position utils.Position
}

func (s *exploreState) Enter() {
	log.Print("Entering Explore State")
}

func (s exploreState) Update(a *StateMachine, g terrain.Grid) State {
	m := a.Memory()
	done := m.ExplorePlace(s.position)

	if done {
		return CreateState("move", struct {
			Position utils.Position
			Next     string
		}{
			Position: s.position.Relative(1, 0),
			Next:     "explore",
		})
	}

	return nil
}
