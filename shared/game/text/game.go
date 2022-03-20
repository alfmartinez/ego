package text

import (
	"ego/shared/input/prompt"
	"fmt"
)

type textGame struct {
	inputHandler prompt.TextHandler
}

func (g *textGame) Start() {
	g.scanner()
}

func (g *textGame) Stop() {

}

func (g *textGame) scanner() {
	var stopLoop bool
	for !stopLoop {
		input := g.inputHandler.GetText()
		fmt.Printf("Got : %s\n", input)
		if input == "exit" {
			stopLoop = true
		}
	}
}
