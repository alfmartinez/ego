package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("explore", func(data []interface{}) State {
		return &exploreState{}
	})
}

type exploreState struct {
}

func (s *exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Enter() {

}

func (s *exploreState) Update(a StateMachine, g terrain.Terrain) State {
	if !a.HasInterests() {
		interests := g.SearchAround(a, 3, func(t terrain.Tile) bool {
			return !a.HasExplored(t)
		})
		a.AddInterests(interests)
	}
	return nil
}

func (s *exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
