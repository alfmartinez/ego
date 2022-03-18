package state

import (
	"ego/pkg/movement"
)

func init() {
	RegisterStateFactory(StateMove, func(data []interface{}) State {
		if len(data) > 0 {
			return &moveState{StateMove, 0, data[0].(movement.Direction)}
		} else {
			return &moveState{StateMove, 0, 0}
		}

	})
}

type moveState struct {
	StateType
	frame     int
	direction movement.Direction
}

type Mover interface {
	Updatable
	MoveDirection(movement.Direction)
}

func (s *moveState) Update(a Updatable) State {
	mover := a.(Mover)
	a.Frame(s.frame, 0)
	s.frame = (s.frame + 1) % 20
	if s.frame > 5 && s.frame < 15 {
		mover.MoveDirection(s.direction)
	}
	if s.frame == 0 {
		return CreateState(StateIdle)
	}
	return nil
}
