package observer

type EventType uint
type ObserverFunc func(Observer)

const (
	UPDATE EventType = iota
	RENDER
	PHYSICS
	EXIT
)

type Event interface {
	Type() EventType
	Data() []interface{}
}

type event struct {
	eventType EventType
	data      []interface{}
}

func CreateEvent(t EventType, data ...interface{}) Event {
	return event{t, data}
}

func (e event) Type() EventType {
	return e.eventType
}

func (e event) Data() []interface{} {
	return e.data
}
