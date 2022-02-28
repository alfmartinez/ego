package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type restState struct {
}

func (s *restState) Label() string {
	return "resting"
}

func (s *restState) Enter() {

}

func (s *restState) Update(a StateMachine, g terrain.Terrain) State {
	return nil
}

func (s *restState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
