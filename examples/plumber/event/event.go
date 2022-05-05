package event

type Event interface {
	Type() EventType
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=EventType
type EventType int

const (
	StartEvent EventType = iota
	StopEvent
)

func CreateEvent(eType EventType) Event {
	return &event{
		eType: eType,
	}
}

type event struct {
	eType EventType
}

// Type implements Event
func (e *event) Type() EventType {
	return e.eType
}
