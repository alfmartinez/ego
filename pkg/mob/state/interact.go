package state

import (
	"ego/pkg/mob/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type interactState struct {
}

func (s *interactState) Label() string {
	return "interacting"
}

func (s *interactState) Enter() {

}

func (s *interactState) Update(a StateMachine, g terrain.Terrain) State {
	return CreateState("move", struct {
		Destination movement.Positionnable
		Next        string
	}{
		Destination: movement.CreateDummy(640, 480),
		Next:        "explore",
	})
}

func (s *interactState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
