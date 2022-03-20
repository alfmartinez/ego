package states

import (
	"ego/engine/input"
	"ego/engine/observer"
	"ego/engine/renderer"
	"ego/engine/state"
	"ego/engine/template"
	"ego/shared/input/prompt"
	"github.com/inancgumus/screen"
	"strings"
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
			Items: []string{"key", "stapler", "orange"},
		},
	}

	renderTemplate := func(name string) {
		screen.Clear()
		screen.MoveTopLeft()
		tmpl.ExecuteTemplate(writer, "default", status)
		status.Feedback = ""
		status.Error = ""
	}

	checkInput := func(f InputFunc) DurationFunc {
		inputHandler := input.FromContext().(prompt.TextHandler)
		render := renderer.FromContext().Render
		subject := observer.FromContext()
		return func(time.Duration) string {
			return f(func() []string {
				render("> ")
				input := inputHandler.GetText()
				if input == "exit" {
					subject.NotifyAll(observer.CreateEvent(observer.EXIT))
				}
				return strings.Split(input, " ")
			})
		}
	}

	return state.States{
		"default": checkInput(func(getInput func() []string) string {
			renderTemplate("default")
			input := getInput()
			var verb, subject string = input[0], input[1]
			var next string
			switch {
			case verb == "go" && subject == "outside":
				if status.Default.DoorUnlocked {
					next = "outside"
				} else {
					status.Feedback = "The door does not open."
				}
			case verb == "open" && subject == "door":
				status.Default.DoorOpen = true
			default:
				status.Error = strings.Join(input, " ")
			}
			return next
		}),
		"outside": checkInput(func(getInput func() []string) string {
			renderTemplate("outside")
			input := getInput()
			var verb, subject string = input[0], input[1]
			var next string
			switch {
			case verb == "go" && subject == "inside":
				next = "default"
			default:
				status.Error = strings.Join(input, " ")
			}
			return next
		}),
	}
}

type InputFunc func(func() []string) string
type DurationFunc func(time.Duration) string
