package state

import (
	"ego/pkg/renderer"
	"ego/pkg/terrain"
	"ego/pkg/utils"
)

type idleState struct {
}

func (s idleState) Label() string {
	return "preparing"
}

func (s idleState) Enter() {
}

func (s idleState) Update(a *StateMachine, g terrain.Terrain) State {
	a.UpdateInterests(a.Position(), func(pos utils.Position) bool {
		return g.GetTile(pos) != nil
	})
	return CreateState("explore")
}

func (s idleState) Render(r renderer.Renderer, m renderer.Renderable) {
	r.Render(m)
}
