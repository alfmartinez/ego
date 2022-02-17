package state

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"log"
)

type idleState struct {
}

func (s idleState) Enter() {
	log.Print("Entering Idle State")
}

func (s idleState) Update(a *StateMachine, g terrain.Terrain) State {
	return CreateState("explore")
}

func (s idleState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
