package state

import (
	"ego/pkg/motivator"
	"ego/pkg/terrain"
)

func init() {
	RegisterStateFactory("idle", func(data []interface{}) State {
		return &idleState{}
	})
}

type Idler interface {
	motivator.NeedsCollection
}

type idleState struct {
	baseState
}

func (s *idleState) Label() string {
	return "preparing"
}

func (s *idleState) Update(sm interface{}, g terrain.Terrain) State {
	a := sm.(Idler)
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
