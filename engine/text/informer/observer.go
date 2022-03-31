package informer

type Publisher interface {
	Subscribe(chan Message)
	Publish(Message)
}

func CreatePublisher() Publisher {
	return &publisher{}
}

type publisher struct {
	observers []chan Message
}

func (p *publisher) Subscribe(c chan Message) {
	p.observers = append(p.observers, c)
}

func (p *publisher) Publish(msg Message) {
	for _, c := range p.observers {
		c <- msg
	}
}
