package informer

type StoryRule interface {
	Matches(Story) bool
	Execute(Story) bool
}

type storyRule struct {
	match func(Story) bool
	exec  func(Story) bool
}

func (r *storyRule) Matches(s Story) bool {
	return r.match(s)
}

func (r *storyRule) Execute(s Story) bool {
	return r.exec(s)
}

func CreateConnectorRule(o Object, t Object, direction string) StoryRule {
	return &storyRule{
		match: func(s Story) bool {
			return s.CurrentRoom() == o && s.Command().Direction.Value == direction
		},
		exec: func(s Story) bool {
			s.SetCurrentRoom(t)
			return true
		},
	}
}

func CreateWhenSayRule(when Phase, say string) StoryRule {
	return &storyRule{
		match: func(s Story) bool {
			return s.Phase() == when
		},
		exec: func(s Story) bool {
			s.Say(say)
			return true
		},
	}
}

func CreatePlayerInventoryRule(items []Object) StoryRule {
	return &storyRule{
		match: func(s Story) bool {
			return s.Phase() == START_PHASE
		},
		exec: func(s Story) bool {
			s.AddToInventory(items)
			return true
		},
	}
}
