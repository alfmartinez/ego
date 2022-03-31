package informer

import (
	"fmt"
)

type StoryRule interface {
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
		fmt.Printf("Checking if %s matches\n", r.name)
	}
	if r.match(msg) {
		if s.Debug() {
			fmt.Printf("Applying rule %s\n", r.name)
		}
		r.exec(s)
	}
}

func (r *storyRule) Name() string {
	return r.name
}

func CreateConnectorRule(name string, o Object, t Object, direction string) StoryRule {
	return &storyRule{
		name: "move between rooms rule",
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

func CreateWhenRule(name string, when Phase, f Activity) StoryRule {
	return &storyRule{
		name: "when phase say rule",
		match: func(msg Message) bool {
			return msg.Phase == when
		},
		exec: f,
	}
}

func CreateActivityRule(name string, o Object, f Activity) StoryRule {
	return &storyRule{
		name: "when activity say rule",
		match: func(msg Message) bool {
			action := msg.Action
			return o.Get("name") == string(action)
		},
		exec: f,
	}
}
