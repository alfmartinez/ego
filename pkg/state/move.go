package state

import (
	"ego/pkg/movement"
	"time"
)

func init() {
	RegisterStateFactory(StateMove, func(data []interface{}) State {
		if len(data) > 0 {
			return &moveState{StateMove, 0, data[0].(movement.Direction), false}
		} else {
			return &moveState{StateMove, 0, 0, false}
		}

	})
}

type moveState struct {
	StateType
	frame     int
	direction movement.Direction
	impulse   bool
}

type Mover interface {
	Updatable
	MoveDirection(movement.Direction, time.Duration)
}

func (s *moveState) Update(a Updatable, dt time.Duration) State {
	mover := a.(Mover)
	a.Frame(s.frame, 0)
	s.frame = (s.frame + 1) % 20
	if !s.impulse {
		mover.MoveDirection(s.direction, dt)
		s.impulse = true
	} else {
		mover.MoveDirection(movement.MOVE_NONE, dt)
	}
	if s.frame == 0 {
		return CreateState(StateIdle)
	}
	return nil
}
