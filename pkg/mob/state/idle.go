package state

import (
	"ego/pkg/renderable"
	"ego/pkg/renderer"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("idle", func(data []interface{}) State {
		return &idleState{}
	})
}

type idleState struct {
}

func (s *idleState) Label() string {
	return "preparing"
}

func (s *idleState) Enter() {
}

func (s *idleState) Update(a StateMachine, g terrain.Terrain) State {
	topNeed := a.TopNeed()
	switch topNeed {
	case "curiosity":
		return CreateState("explore")
	case "social":
		return CreateState("interact")
	case "rest":
		return CreateState("rest")
	case "health":
		return CreateState("heal")
	}
	return nil

}

func (s *idleState) Render(r renderer.Renderer, m renderable.Renderable) {
	r.Render(m)
}
