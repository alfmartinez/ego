package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("interact", func(data []interface{}) State {
		return &interactState{}
	})
}

type interactState struct {
}

func (s *interactState) Label() string {
	return "interacting"
}

func (s *interactState) Enter() {

}

func (s *interactState) Update(a interface{}, g terrain.Terrain) State {
	return nil
}

func (s *interactState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
