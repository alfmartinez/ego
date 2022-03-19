package observer

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	NotifyAll(Event)
}

type subject struct {
	observers map[Observer]bool
}

func CreateSubject() Subject {
	o := make(map[Observer]bool)
	return &subject{observers: o}
}

func (s *subject) Register(o Observer) {
	s.observers[o] = true
}

func (s *subject) Unregister(o Observer) {
	delete(s.observers, o)
}

func (s *subject) NotifyAll(e Event) {
	for o := range s.observers {
		o.OnNotify(e)
	}
}
