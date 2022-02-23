package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
)

type exploreState struct {
}

func (s exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Enter() {

}

func (s exploreState) Update(a StateMachine, g terrain.Terrain) State {
	position := a.Position()
	done := a.ExplorePlace(position)

	if done {
		nextTile, found := a.SearchNextPositionToExplore(position)
		if found {
			return CreateState("move", struct {
				Position utils.Position
				Next     string
			}{
				Position: nextTile,
				Next:     "explore",
			})
		} else {
			return CreateState("idle")
		}

	}

	return nil
}

func (s exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
