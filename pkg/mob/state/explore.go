package state

import (
	"ego/pkg/mob/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type exploreState struct {
}

func (s *exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Enter() {

}

func (s *exploreState) Update(a StateMachine, g terrain.Terrain) State {
	return CreateState("move", struct {
		Destination movement.Positionnable
		Next        string
	}{
		Destination: movement.CreateDummy(640, 480),
		Next:        "explore",
	})
}

func (s *exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
