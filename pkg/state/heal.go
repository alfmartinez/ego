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
		tile := g.FindClosest(patient, func(t terrain.Tile) bool {
			if t == nil {
				return false
			}
			return t.HasResource("health")
		})
		if tile != nil {
			s.target = tile
		}
	}
	if patient.IsAt(s.target) {
		patient.Provide(motivator.CreateNeed("health", 0), 1, 10)
		s.target.Consume("health")
	} else {
		return CreateState("move", struct {
			Destination movement.Positionnable
			Next        State
		}{Destination: s.target, Next: s})
	}
	return CreateState("idle")
}

func (s *healState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
