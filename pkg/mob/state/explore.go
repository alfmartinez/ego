package state

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
	"log"
)

type exploreState struct {
}

func (s exploreState) Label() string {
	return "explore"
}

func (s *exploreState) Enter() {
	log.Print("Entering Explore State")
}

func (s exploreState) Update(a *StateMachine, g terrain.Terrain) State {
	position := a.Position()
	done := a.ExplorePlace(position)

	if done {
		return CreateState("move", struct {
			Position utils.Position
			Next     string
		}{
			Position: position.Relative(1, 0),
			Next:     "explore",
		})
	}

	return nil
}

func (s exploreState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
