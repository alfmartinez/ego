package input

import "github.com/alfmartinez/ego/engine/context"

func FromContext() InputHandler {
	return context.GetContext().Get("input").(InputHandler)
}
