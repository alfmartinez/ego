package informer

import (
	"fmt"
)

type StoryRule interface {
	SetName(string) StoryRule
	Name() string
	OnNotify(Message) bool
}

type storyRule struct {
	name  string
	books []string
	match func(Message) bool
	exec  func(Story) bool
}

func (r *storyRule) OnNotify(msg Message) bool {
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
	return true
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
		match: func(msg Message) bool {
			s := msg.Story
			return msg.Action == "going" && s.CurrentRoom() == o && s.IsSame(msg.Argument, direction.Get("name"))
		},
		exec: func(s Story) bool {
			s.SetCurrentRoom(t)
			s.Publish(Message{
				Action:   "enter",
				Argument: "location",
			})
			return true
		},
	}
}

func CreateWhenRule(when Phase, f Activity) StoryRule {
	return &storyRule{
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
