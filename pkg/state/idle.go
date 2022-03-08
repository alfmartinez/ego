package state

import "ego/pkg/motivator"

func init() {
	RegisterStateFactory("idle", func(data []interface{}) State {
		return &idleState{}
	})
}

type Idler interface {
	TopNeed() motivator.Need
}

type idleState struct{}

func (s *idleState) Update(sm Updatable) State {
	a := sm.(Idler)
	topNeed := a.TopNeed()
	switch topNeed {
	case motivator.Learn:
		return CreateState("explore")
	case motivator.Recreation:
		return CreateState("interact")
	case motivator.Rest:
		return CreateState("rest")
	case motivator.Health:
		return CreateState("heal")
	}
	return nil

}
