package states

import (
	"ego/engine/input"
	"ego/engine/observer"
	"ego/engine/renderer"
	"ego/engine/state"
	"ego/engine/template"
	"ego/shared/input/prompt"
	"time"
)

type TextWriter struct {
	render func(any)
}

func (w *TextWriter) Write(p []byte) (n int, err error) {
	w.render(string(p))
	return len(p), nil
}

func States(a any) state.States {
	writer := &TextWriter{renderer.FromContext().Render}
	tmpl := template.FromContext()
	status := struct {
		Feedback string
		Error    string
		DoorOpen bool
	}{}

	renderTemplate := func(name string) {
		tmpl.ExecuteTemplate(writer, "default", status)
		status.Feedback = ""
		status.Error = ""
	}

	return state.States{
		"default": checkInput(func(getInput func() string) string {
			renderTemplate("default")
			input := getInput()

			var next string
			switch {
			case input == "go outside":
				if status.DoorOpen {
					next = "outside"
				} else {
					status.Feedback = "The door does not open."
				}
			case input == "open door":
				status.DoorOpen = true
			default:
				status.Error = input
			}
			return next
		}),
		"outside": checkInput(func(getInput func() string) string {
			renderTemplate("outside")
			input := getInput()
			var next string
			switch {
			case input == "go inside":
				next = "default"
			default:
				status.Error = input
			}
			return next
		}),
	}
}

type InputFunc func(func() string) string
type DurationFunc func(time.Duration) string

func checkInput(f InputFunc) DurationFunc {
	inputHandler := input.FromContext().(prompt.TextHandler)
	render := renderer.FromContext().Render
	subject := observer.FromContext()
	return func(time.Duration) string {
		return f(func() string {
			render("> ")
			input := inputHandler.GetText()
			if input == "exit" {
				subject.NotifyAll(observer.CreateEvent(observer.EXIT))
			}
			return input
		})
	}
}
