package informer

type Rulebooks interface {
	Register(StoryRule)
	Publish(Message)
	AddRulebook(string)
}

func CreateRulebooks() Rulebooks {
	return &rulebooks{}
}

type rulebooks struct {
	observers []StoryRule
	rulebooks []string
}

func (p *rulebooks) AddRulebook(book string) {
	p.rulebooks = append(p.rulebooks, book)
}

func (p *rulebooks) Register(rule StoryRule) {
	p.observers = append(p.observers, rule)
}

func (p *rulebooks) Publish(msg Message) {

	for _, c := range p.observers {
		if _, ok := c.Try(msg); !ok {
			break
		}
	}
}
