package state

import (
	"ego/pkg/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("move", func(data []interface{}) State {
		args := data[0].(struct {
			Destination movement.Positionnable
			Next        string
		})
		return &moveState{args.Destination, args.Next}
	})
}

type Mover interface {
	movement.Movement
}

type moveState struct {
	destination movement.Positionnable
	next        string
}

func (s *moveState) Label() string {
	return "moving"
}

func (s *moveState) Enter() {
}

func (s *moveState) Update(sm interface{}, g terrain.Terrain) State {
	a := sm.(Mover)
	done := a.MoveTowards(s.destination)
	if done {
		return CreateState(s.next)
	}
	return nil
}

func (s *moveState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
