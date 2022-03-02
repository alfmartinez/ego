package state

import (
	"ego/pkg/object/movement"
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
	exploring movement.Positionnable
}

func (s *exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Enter() {

}

func (s *exploreState) Update(a StateMachine, g terrain.Terrain) State {
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
				Next        string
			}{Destination: s.exploring, Next: "explore"})
		}
	}
	return nil

}

func (s *exploreState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
