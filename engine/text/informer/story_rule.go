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
	exec  func(Story) RuleResult
}

func (r *storyRule) IsInBook(book string) bool {
	if len(r.books) == 0 {
		panic(fmt.Errorf("Rule has no book %v", r))
	}
	return slices.Contains(r.books, book)
}

func (r *storyRule) Try(msg Message) RuleResult {
	s := msg.Story
	if s.Debug() {
		//fmt.Printf("Checking if %s matches\n", r.name)
	}
	if r.match(msg) {
		if s.Debug() {
			fmt.Printf("Applying rule %s\n", r.name)
		}
		return r.exec(s)
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

func CreateConnectorRule(o Object, t Object, direction Object) StoryRule {
	return &storyRule{
		books: []string{"actions"},
		match: func(msg Message) bool {
			s := msg.Story
			return msg.Action == "going" && s.CurrentRoom() == o && s.IsSame(msg.Argument, direction.Get("name"))
		},
		exec: func(s Story) RuleResult {
			s.SetCurrentRoom(t)
			s.Publish(Message{
				Action:   "enter",
				Argument: "location",
			})
			return RULE_SUCCESS
		},
	}
}

func CreateWhenRule(book string, when Phase, f Activity) StoryRule {
	return &storyRule{
		books: []string{book},
		match: func(msg Message) bool {
			return msg.Phase == when
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
		match: func(msg Message) bool {
			s := msg.Story
			matches := s.IsSame(o.Get("name"), string(msg.Action))
			if matches {
				argument := s.GetObject(msg.Argument)
				if argument == nil {
					panic(fmt.Errorf("index does not know %q", msg.Argument))
				}
				s.SetAlias(alias, argument)
			}
			return matches
		},
		exec: f,
	}
}
