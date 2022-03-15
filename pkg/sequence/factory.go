package sequence

import "ego/pkg/command"

type SequenceType int

const (
	Evaluate SequenceType = iota
	Seek
)

type SequenceFactory func(...interface{}) command.Command

var factories = make(map[SequenceType]SequenceFactory)

func RegisterSequenceFactory(t SequenceType, f SequenceFactory) {
	factories[t] = f
}

func CreateSequence(t SequenceType) SequenceFactory {
	return factories[t]
}
