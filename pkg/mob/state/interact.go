package state

import (
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
	return nil
}

func (s *interactState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
