package state

import (
	"ego/pkg/motivator"
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
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

func (s *healState) Label() string {
	return "healing"
}

func (s *healState) Enter() {}

func (s *healState) Update(a interface{}, g terrain.Terrain) State {
	patient := a.(Patient)
	if s.target == nil {
		tiles := g.FindClosest(patient, 1, func(t terrain.Tile) bool {
			return t.HasResource("health")
		})
		tile := tiles[0]
		if tile != nil {
			s.target = tile
		}
	}
	if patient.IsAt(s.target) {
		patient.Provide(motivator.CreateNeed("health", 0), 1, 10)
		s.target.Consume("health")
	} else {
		return CreateState("move", MoveArguments{Destination: s.target, Next: s})
	}
	return CreateState("idle")
}

func (s *healState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
