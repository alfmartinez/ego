package state

import (
	"ego/pkg/motivator"
)

func init() {
	RegisterStateFactory("idle", func(data []interface{}) State {
		return &idleState{}
	})
}

type Idler interface {
	motivator.NeedsCollection
}

type idleState struct{}

func (s *idleState) Update(sm Updatable) State {
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
