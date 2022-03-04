package state

import (
	"ego/pkg/memory"
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
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
			interests := g.FindClosest(a, 9, func(t terrain.Tile) bool {
				return !a.HasExplored(t)
			})
			positions := make([]movement.Positionnable, len(interests))
			for _, x := range interests {
				log.Printf(" --- %+v", x)
				if x != nil {
					if p, ok := x.(movement.Positionnable); ok {
						positions = append(positions, p)
					}
				}
			}
			a.AddInterests(positions)
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
