package input

import "ego/engine/context"

func FromContext() InputHandler {
	return context.GetContext().Get("input").(InputHandler)
}
