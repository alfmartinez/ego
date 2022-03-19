package state

import (
	"ego/engine/movement"
	"time"
)

func init() {
	RegisterStateFactory(StateMove, func(data []interface{}) State {
		if len(data) > 0 {
			return &moveState{
				StateType: StateMove,
				direction: data[0].(movement.Direction),
			}
		} else {
			return &moveState{
				StateType: StateMove,
			}
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
	if !s.impulse && s.frame > 5 {
		mover.MoveDirection(s.direction, time.Second)
		s.impulse = true
	}
	if s.frame == 15 {
		mover.MoveDirection(s.direction, -time.Second)
	}
	if s.frame == 0 {
		return CreateState(StateIdle)
	}
	return nil
}
