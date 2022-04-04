package informer

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type RuleResult int

const (
	RULE_SUCCESS RuleResult = iota
	RULE_FAILURE
	RULE_UNDECIDED
)

type StoryRule interface {
	SetName(string) StoryRule
	Name() string
	Try(Message) RuleResult
	IsInBook(string) bool
}

type storyRule struct {
	name  string
	books []string
	match func(Message) bool
	exec  func() RuleResult
}

func (r *storyRule) IsInBook(book string) bool {
	if len(r.books) == 0 {
		panic(fmt.Errorf("Rule has no book %v", r))
	}
	return slices.Contains(r.books, book)
}

func (r *storyRule) Try(message Message) RuleResult {
	msg, ok := message.(Try)
	if !ok {
		return RULE_UNDECIDED
	}
	if Debug() {
		//fmt.Printf("Checking if %s matches\n", r.name)
	}
	if r.match(msg) {
		if Debug() {
			fmt.Printf("Applying rule %s\n", r.name)
		}
		return r.exec()
	}
	return RULE_UNDECIDED
}

func (r *storyRule) SetName(name string) StoryRule {
	r.name = name
	return r
}

func (r *storyRule) Name() string {
	return r.name
}

func CreateWhenRule(book string, f Activity) StoryRule {
	return &storyRule{
		books: []string{book},
		match: func(msg Message) bool {
			return true
		},
		exec: f,
	}
}

func CreateActivityRule(o Object, f Activity, alias string) StoryRule {
	if !o.IsKind("action") {
		panic(fmt.Errorf("Cannot create activity rule with non-action %q", o.Get("name")))
	}
	return &storyRule{
		books: []string{fmt.Sprintf("carry out %s", o.Get("name"))},
		match: func(message Message) bool {
			msg := message.(Try)
			matches := IsSame(o.Get("name"), string(msg.Action))
			if matches {
				argument := GetObject(msg.Argument)
				if argument == nil {
					panic(fmt.Errorf("index does not know %q", msg.Argument))
				}
				SetAlias(alias, argument)
			}
			return matches
		},
		exec: f,
	}
}
