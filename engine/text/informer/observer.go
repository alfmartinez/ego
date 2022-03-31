package informer

type Publisher interface {
	Subscribe(func(Message))
	Publish(Message)
}

func CreatePublisher() Publisher {
	return &publisher{}
}

type publisher struct {
	observers []func(Message)
}

func (p *publisher) Subscribe(listener func(Message)) {
	p.observers = append(p.observers, listener)
}

func (p *publisher) Publish(msg Message) {
	for _, c := range p.observers {
		c(msg)
	}
}
