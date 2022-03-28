package states

import (
	"github.com/alfmartinez/ego/engine/input"
	"github.com/alfmartinez/ego/engine/observer"
	"github.com/alfmartinez/ego/engine/renderer"
	"github.com/alfmartinez/ego/engine/state"
	"github.com/alfmartinez/ego/engine/template"
	"github.com/alfmartinez/ego/shared/input/prompt"
	"github.com/inancgumus/screen"
	"time"
)

type TextWriter struct {
	render func(any)
}

func (w *TextWriter) Write(p []byte) (n int, err error) {
	w.render(string(p))
	return len(p), nil
}

type defaultState struct {
	DoorUnlocked bool
	DoorOpen     bool
	Items        []string
}

func States(a any) state.States {
	writer := &TextWriter{renderer.FromContext().Render}
	tmpl := template.FromContext()
	status := struct {
		Inventory []string
		Feedback  string
		Error     string
		Default   defaultState
	}{
		Default: defaultState{
			Items: []string{"key", "orange"},
		},
	}

	renderTemplate := func(name string) {
		screen.Clear()
		screen.MoveTopLeft()
		tmpl.ExecuteTemplate(writer, name, status)
		status.Feedback = ""
		status.Error = ""
	}

	checkInput := func(f InputFunc) DurationFunc {
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

	return state.States{
		"default": checkInput(func(getInput func() string) string {
			return "office"
		}),
		"office": checkInput(func(getInput func() string) string {
			renderTemplate("office.tmpl")
			input := getInput()
			var next string
			switch {
			case input == "sortir":
				next = "outside"
			case input == "go outside":
				if status.Default.DoorUnlocked {
					next = "outside"
				} else {
					status.Feedback = "The door does not open."
				}
			case input == "open door":
				status.Default.DoorOpen = true
			case input == "open drawer":
				status.Default.Items = append(status.Default.Items, "key")
			default:
				status.Error = input
			}
			return next
		}),
		"outside": checkInput(func(getInput func() string) string {
			renderTemplate("outside.tmpl")
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
