package input

import "ego/pkg/context"

func FromContext() InputHandler {
	return context.GetContext().Get("input").(InputHandler)
}
