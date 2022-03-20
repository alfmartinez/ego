package states

import (
	"ego/engine/input"
	"ego/engine/observer"
	"ego/engine/renderer"
	"ego/engine/state"
	"ego/shared/input/prompt"
	"fmt"
	"time"
)

func States(a any) state.States {
	render := renderer.FromContext().Render
	return state.States{
		"default": checkInput(func(getInput func() string) string {
			render("\n--------------\n")
			render("You are inside an office building.\nThere's a door behind you, leading outside.\n")
			render("> ")
			input := getInput()

			var next string
			switch {
			case input == "go outside":
				next = "outside"
			default:
				displayError(input)
			}
			return next
		}),
		"outside": checkInput(func(getInput func() string) string {
			render("\n--------------\n")
			render("You are outisde the building.\n")
			render("> ")
			input := getInput()
			if input == "enter building" {
				return "default"
			}
			return ""
		}),
	}
}

type InputFunc func(func() string) string
type DurationFunc func(time.Duration) string

func checkInput(f InputFunc) DurationFunc {
	inputHandler := input.FromContext().(prompt.TextHandler)
	subject := observer.FromContext()
	return func(time.Duration) string {
		return f(func() string {
			input := inputHandler.GetText()
			if input == "exit" {
				subject.NotifyAll(observer.CreateEvent(observer.EXIT))
			}
			return input
		})
	}
}

func displayError(input string) {
	render := renderer.FromContext().Render
	render(fmt.Sprintf("What do you mean exactly by \"%s\" ?", input))
}
