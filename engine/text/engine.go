package text

import (
	"ego/engine/state"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"os"
)

type Engine interface {
	AddText(path string)
	Transcribe() state.States
}

func New() Engine {
	return &textEngine{}
}

type textEngine struct {
	file string
}

func (e *textEngine) AddText(path string) {
	e.file = path
}
func (e *textEngine) Transcribe() state.States {
	parser := participle.MustBuild(&Story{}, participle.Lexer(def)) //, participle.Elide("Whitespace"))

	f, _ := os.Open(e.file)

	ast := &Story{}
	err := parser.Parse(e.file, f, ast)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", ast)

	return state.States{}
}
