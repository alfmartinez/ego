package state

import (
	"ego/pkg/memory"
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("explore", func(data []interface{}) State {
		return &exploreState{}
	})
}

type Explorer interface {
	memory.Memory
	movement.Positionnable
}

type exploreState struct {
	exploring movement.Positionnable
}

func (s *exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Enter() {}

func (s *exploreState) Update(sm interface{}, g terrain.Terrain) State {
	a, _ := sm.(Explorer)
	if s.exploring == nil {
		if !a.HasInterests() {
			interests := g.SearchAround(a, 3, func(t terrain.Tile) bool {
				return !a.HasExplored(t)
			})
			a.AddInterests(interests)
		}
		s.exploring = a.GetNextInterest()
	} else {
		if a.Position().Eq(s.exploring.Position()) {
			done := a.Explore(s.exploring)
			if done {
				return CreateState("idle")
			}
		} else {
			return CreateState("move", struct {
				Destination movement.Positionnable
				Next        State
			}{Destination: s.exploring, Next: s})
		}
	}

	return nil

}

func (s *exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
