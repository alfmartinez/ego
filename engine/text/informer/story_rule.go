package informer

import (
	"fmt"
)

type StoryRule interface {
	SetName(string) StoryRule
	Name() string
	OnNotify(Message)
}

type storyRule struct {
	name  string
	match func(Message) bool
	exec  func(Story) bool
}

func (r *storyRule) OnNotify(msg Message) {
	s := msg.Story
	if s.Debug() {
		//fmt.Printf("Checking if %s matches\n", r.name)
	}
	if r.match(msg) {
		if s.Debug() {
			fmt.Printf("Applying rule %s\n", r.name)
		}
		r.exec(s)
	}
}

func (r *storyRule) SetName(name string) StoryRule {
	r.name = name
	return r
}

func (r *storyRule) Name() string {
	return r.name
}

func CreateConnectorRule(o Object, t Object, direction string) StoryRule {
	return &storyRule{
		match: func(msg Message) bool {
			s := msg.Story
			cmd := s.Command()
			return cmd != nil && s.Phase() == UPDATE_PHASE && s.CurrentRoom() == o && cmd.Direction.Value == direction
		},
		exec: func(s Story) bool {
			s.SetCurrentRoom(t)
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

func CreateActivityRule(o Object, f Activity) StoryRule {
	if !o.IsKind("action") {
		panic(fmt.Errorf("Cannot create activity rule with non-action %q", o.Get("name")))
	}
	return &storyRule{
		match: func(msg Message) bool {
			s := msg.Story
			return s.IsSame(o.Get("name"), string(msg.Action))
		},
		exec: f,
	}
}
