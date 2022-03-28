package text

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alfmartinez/ego/engine/state"
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
	parser := participle.MustBuild(&Grammar{}, participle.Lexer(def)) //, participle.Elide("Whitespace"))

	f, _ := os.Open(e.file)

	ast := &Grammar{}
	err := parser.Parse(e.file, f, ast)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", ast)

	return state.States{}
}
