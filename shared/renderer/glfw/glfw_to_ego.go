package glfw

import (
	"ego/engine/input"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func ConvertKey(k glfw.Key) input.Key {
	var key input.Key = input.UNKNOWN_KEY
	switch k {
	case glfw.KeyEscape:
		key = input.ESCAPE
	case glfw.KeyLeft:
		key = input.LEFT
	case glfw.KeyRight:
		key = input.RIGHT
	case glfw.KeyUp:
		key = input.UP
	case glfw.KeyDown:
		key = input.DOWN
	case glfw.KeySpace:
		key = input.JUMP
	}
	return key
}

func ConvertAction(a glfw.Action) input.Action {
	var action input.Action = input.UNKNOWN_ACTION
	switch a {
	case glfw.Press:
		action = input.PRESSED
	case glfw.Release:
		action = input.RELEASED
	case glfw.Repeat:
		action = input.REPEATED
	}
	return action
}
