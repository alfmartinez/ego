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
	return "exploring"
}

func (s *exploreState) Enter() {
	log.Print("Entering Explore State")
}

func (s exploreState) Update(a *StateMachine, g terrain.Terrain) State {
	position := a.Position()
	done := a.ExplorePlace(position)

	if done {
		nextTile, found := a.SearchNextPositionToExplore()
		if found {
			return CreateState("move", struct {
				Position utils.Position
				Next     string
			}{
				Position: nextTile,
				Next:     "explore",
			})
		}

	}

	return nil
}

func (s exploreState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
