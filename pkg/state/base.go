package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

// Base state
type baseState struct{}

func (s *baseState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}

func (s *baseState) Update(a interface{}, g terrain.Terrain) State {
	return nil
}

func (s *baseState) Enter() {}
