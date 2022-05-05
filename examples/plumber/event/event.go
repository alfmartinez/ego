package event

type Event interface {
	Type() EventType
}

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
