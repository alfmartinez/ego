package informer

type Rulebooks interface {
	Register(StoryRule)
	Publish(Message)
}

func CreatePublisher() Rulebooks {
	return &publisher{}
}

type publisher struct {
	observers []StoryRule
}

func (p *publisher) Register(rule StoryRule) {
	p.observers = append(p.observers, rule)
}

func (p *publisher) Publish(msg Message) {
	for _, c := range p.observers {
		if ok := c.Try(msg); !ok {
			break
		}
	}
}
