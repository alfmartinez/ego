package state

import (
	"ego/pkg/mob/movement"
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

type moveState struct {
	destination movement.Positionnable
	next        string
}

func (s *moveState) Label() string {
	return "moving"
}

func (s *moveState) Enter() {
}

func (s *moveState) Update(a StateMachine, g terrain.Terrain) State {
	done := a.MoveTowards(s.destination.Position())
	if done {
		return CreateState(s.next)
	}
	return nil
}

func (s *moveState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
