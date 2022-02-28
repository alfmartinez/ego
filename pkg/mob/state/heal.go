package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type healState struct {
}

func (s *healState) Label() string {
	return "healing"
}

func (s *healState) Enter() {

}

func (s *healState) Update(a StateMachine, g terrain.Terrain) State {
	return nil
}

func (s *healState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
