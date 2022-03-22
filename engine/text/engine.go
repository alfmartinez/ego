package text

import (
	"ego/engine/state"
)

type Engine interface {
	AddText(path string)
	Transcribe() state.States
}

func New() Engine {
	return &textEngine{}
}

type textEngine struct{}

func (*textEngine) AddText(path string) {

}
func (*textEngine) Transcribe() state.States {
	return state.States{}
}
