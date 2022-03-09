package state

import "ego/pkg/motivator"

func init() {
	RegisterStateFactory("idle", func(data []interface{}) State {
		return &idleState{0}
	})
}

type Idler interface {
	TopNeed() motivator.Need
	Frame(x, y int)
}

type idleState struct {
	frame int
}

func (s *idleState) Update(sm Updatable) State {
	a := sm.(Idler)
	a.Frame(s.frame, 0)
	s.frame = (s.frame + 1) % 20
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
