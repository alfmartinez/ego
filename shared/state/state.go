// Package shared State implements specific states for a game
package state

import (
	"github.com/alfmartinez/ego/engine/input"
	"github.com/alfmartinez/ego/engine/movement"
	"github.com/alfmartinez/ego/engine/state"
	"time"
)

func Register() {
	state.RegisterStatesClosure("morris", morrisStates)
	state.RegisterStatesClosure("bomber", bomberStates)
	state.RegisterStatesClosure("goon", goonStates)
}

type Morris interface {
	Frame(x, y int)
	MoveDirection(movement.Direction, time.Duration)
	movement.Positionnable
}

func morrisStates(a any) state.States {
	var m Morris = a.(Morris)
	var inputHandler = input.FromContext().(input.KeyHandler)
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

func bomberStates(a any) state.States {
	var m Morris = a.(Morris)
	var frame int
	var inputHandler = input.FromContext().(input.KeyHandler)
	var direction movement.Direction

	return state.States{
		"default": func(dt time.Duration) string {
			frame = 0
			m.Frame(0, 0)
			switch {
			case inputHandler.IsPressed(input.UP):
				direction = movement.MOVE_UP
				return "move"
			case inputHandler.IsPressed(input.RIGHT):
				direction = movement.MOVE_RIGHT
				return "move"
			case inputHandler.IsPressed(input.LEFT):
				direction = movement.MOVE_LEFT
				return "move"
			case inputHandler.IsPressed(input.DOWN):
				direction = movement.MOVE_DOWN
				return "move"
			}
			return ""
		},
		"move": func(dt time.Duration) string {
			if frame == 0 {
				m.MoveDirection(direction, time.Second)
			}
			frame = (frame + 1) % 15
			if frame == 0 {
				m.MoveDirection(direction, -time.Second)
				return "default"
			}
			return ""
		},
	}
}

func goonStates(a any) state.States {
	var m Morris = a.(Morris)
	var steps int
	var direction movement.Direction

	return state.States{
		"default": func(dt time.Duration) string {
			m.Frame(0, 0)
			switch direction {
			case movement.MOVE_NONE:
				direction = movement.MOVE_DOWN
			case movement.MOVE_DOWN:
				direction = movement.MOVE_LEFT
			case movement.MOVE_LEFT:
				direction = movement.MOVE_UP
			case movement.MOVE_UP:
				direction = movement.MOVE_RIGHT
			case movement.MOVE_RIGHT:
				direction = movement.MOVE_DOWN
			}
			return "patrol"
		},
		"patrol": func(dt time.Duration) string {
			if steps == 0 {
				m.MoveDirection(direction, time.Second)
			}
			steps = (steps + 1) % 30
			if steps == 0 {
				m.MoveDirection(direction, -time.Second)
				return "default"
			}
			return ""
		},
	}
}
