package observer

type EventType uint
type ObserverFunc func(Observer)

const (
	UPDATE EventType = iota
	RENDER
)

type Event interface {
	Type() EventType
}

type event struct {
	eventType EventType
}

func CreateEvent(t EventType) Event {
	return event{t}
}

func (e event) Type() EventType {
	return e.eventType
}
