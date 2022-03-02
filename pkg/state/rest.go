package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("rest", func(data []interface{}) State {
		return &restState{}
	})
}

type restState struct {
}

func (s *restState) Label() string {
	return "resting"
}

func (s *restState) Enter() {

}

func (s *restState) Update(a interface{}, g terrain.Terrain) State {
	return nil
}

func (s *restState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
