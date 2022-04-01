package informer

import (
	"github.com/alfmartinez/ego/engine/slices"
)

type Rulebooks interface {
	Register(StoryRule)
	Publish(Message)
	AddRulebook(string)
}

func CreateRulebooks() Rulebooks {
	return &rulebooks{
		books: []string{},
	}
}

type rulebooks struct {
	observers []StoryRule
	books     []string
}

func (p *rulebooks) AddRulebook(book string) {
	p.books = append(p.books, book)
}

func (p *rulebooks) Register(rule StoryRule) {
	p.observers = append(p.observers, rule)
}

func (p *rulebooks) Publish(msg Message) {
	for _, book := range p.books {
		rules := slices.Filter(p.observers, func(rule StoryRule) bool {
			return rule.IsInBook(book)
		})
		for _, c := range rules {
			if _, ok := c.Try(msg); !ok {
				break
			}
		}
	}

}
