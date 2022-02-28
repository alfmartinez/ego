package state

import (
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
	return nil
}

func (s *exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
