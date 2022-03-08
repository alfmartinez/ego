package state

import (
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("heal", func(data []interface{}) State {
		return &healState{}
	})
}

type Patient interface {
	movement.Positionnable
	motivator.NeedsCollection
}

type healState struct {
	target terrain.Tile
}

func (s *healState) Update(a Updatable) State {
	g := terrain.GetTerrain()
	patient := a.(Patient)
	if s.target == nil {
		tiles := g.FindClosest(patient, 1, func(t terrain.Tile) bool {
			return t.HasResource(terrain.Health)
		})
		tile := tiles[0]
		if tile != nil {
			s.target = tile
		}
	}
	if patient.IsAt(s.target) {
		patient.Provide(motivator.Health, 1, 10)
		s.target.Consume(terrain.Health)
	} else {
		return CreateState("move", MoveArguments{Destination: s.target, Next: s})
	}
	return CreateState("idle")
}
