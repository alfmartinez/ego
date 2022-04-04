package informer

import (
	"github.com/alfmartinez/ego/engine/slices"
)

type Rulebooks interface {
	Register(StoryRule)
	Publish(Message) RuleResult
	AddRulebook(RuleBook)
	GetRuleBook(string) RuleBook
}

var books Rulebooks

func CreateRulebooks() Rulebooks {
	return &rulebooks{
		books: make([]RuleBook, 0),
	}
}

type rulebooks struct {
	observers []StoryRule
	books     []RuleBook
}

func (p *rulebooks) AddRulebook(book RuleBook) {
	p.books = append(p.books, book)
}

func (p *rulebooks) GetRuleBook(name string) RuleBook {
	for _, book := range p.books {
		if book.Matches(name) {
			return book
		}
	}
	return nil
}

func (p *rulebooks) Register(rule StoryRule) {
	p.observers = append(p.observers, rule)
}

func (p *rulebooks) Publish(msg Message) RuleResult {

	// Rulebook object should be able to tell if its rules apply
	// - Kind based answer to queries about object of given kind
	// - Activity based should respond to activity
	// - Action based should respond to action
	// - Phase based should respond to turn sequence

	for _, book := range p.books {
		rules := slices.Filter(p.observers, func(rule StoryRule) bool {
			return rule.IsInBook(book.Name())
		})
		for _, c := range rules {
			if result := c.Try(msg); result != RULE_UNDECIDED {
				return result
			}
		}
	}
	return RULE_UNDECIDED
}
