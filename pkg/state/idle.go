package state

import (
	"ego/pkg/input"
	"ego/pkg/movement"
	"time"
)

func init() {
	RegisterStateFactory(StateIdle, func(data []interface{}) State {
		return &idleState{StateIdle}
	})
}

type idleState struct {
	StateType
}

func (s *idleState) Update(a Updatable, dt time.Duration) State {
	a.Frame(0, 0)

	inputHandler := input.FromContext()
	if inputHandler.IsPressed(input.JUMP) {
		return CreateState(StateMove, movement.MOVE_NONE)
	}
	if inputHandler.IsPressed(input.RIGHT) {
		return CreateState(StateMove, movement.MOVE_RIGHT)
	}
	if inputHandler.IsPressed(input.LEFT) {
		return CreateState(StateMove, movement.MOVE_LEFT)
	}

	return nil

}
