package state

import (
	"ego/pkg/memory"
	"ego/pkg/movement"
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
	baseState
	exploring movement.Positionnable
}

func (s *exploreState) Label() string {
	return "exploring"
}

func (s *exploreState) Update(sm interface{}, g terrain.Terrain) State {
	a, _ := sm.(Explorer)
	if s.exploring == nil {
		if !a.HasInterests() {
			interests := g.FindClosest(a, 9, func(t terrain.Tile) bool {
				return !a.HasExplored(t)
			})
			positions := make([]movement.Positionnable, 0, len(interests))
			for _, x := range interests {
				positions = append(positions, movement.Loc(x.Position()))
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
			return CreateState("move", MoveArguments{Destination: s.exploring, Next: s})
		}
	}

	return nil

}
