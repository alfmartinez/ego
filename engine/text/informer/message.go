package informer

import (
	"github.com/alfmartinez/ego/engine/text/grammar"
)

type Action string

type Message struct {
	Story    Story
	Phase    Phase
	Action   Action
	Argument string
	Command  *grammar.Command
}
