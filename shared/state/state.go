// Package shared State implements specific states for a game
package state

import (
	"ego/engine/input"
	"ego/engine/movement"
	"ego/engine/state"
	"time"
)

func Register() {
	state.RegisterStatesClosure("morris", morrisStates)
}

type Morris interface {
	Frame(x, y int)
	MoveDirection(movement.Direction, time.Duration)
}

func morrisStates(a any) state.States {
	var m Morris = a.(Morris)
	var inputHandler = input.FromContext()
	var direction movement.Direction
	var frame int
	var impulse, jumping bool

	return state.States{
		"default": func(dt time.Duration) string {
			m.Frame(0, 0)
			switch {
			case inputHandler.IsPressed(input.UP):
				direction = movement.MOVE_UP
				jumping = true
				return "move"
			case inputHandler.IsPressed(input.RIGHT):
				direction = movement.MOVE_RIGHT
				return "move"
			case inputHandler.IsPressed(input.LEFT):
				direction = movement.MOVE_LEFT
				return "move"
			}

			return ""
		},
		"move": func(dt time.Duration) string {
			m.Frame(frame, 0)
			frame = (frame + 1) % 20
			if !impulse && frame > 5 {
				m.MoveDirection(direction, time.Second)
				impulse = true
			}
			if frame == 15 && !jumping {
				m.MoveDirection(direction, -time.Second)
			}
			if frame == 0 {
				impulse = false
				jumping = false
				return "default"
			}
			return ""
		},
	}
}
