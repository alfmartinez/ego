package informer

type StoryRule interface {
	Name() string
	Matches(Story) bool
	Execute(Story) bool
}

type storyRule struct {
	name  string
	match func(Story) bool
	exec  func(Story) bool
}

func (r *storyRule) Name() string {
	return r.name
}

func (r *storyRule) Matches(s Story) bool {
	return r.match(s)
}

func (r *storyRule) Execute(s Story) bool {
	return r.exec(s)
}

func CreateConnectorRule(o Object, t Object, direction string) StoryRule {
	return &storyRule{
		name: "move between rooms rule",
		match: func(s Story) bool {
			return s.Phase() == UPDATE_PHASE && s.CurrentRoom() == o && s.Command().Direction.Value == direction
		},
		exec: func(s Story) bool {
			s.SetCurrentRoom(t)
			return true
		},
	}
}

func CreateWhenSayRule(when Phase, say string) StoryRule {
	return &storyRule{
		name: "when phase say rule",
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
		name: "put items in inventory rule",
		match: func(s Story) bool {
			return s.Phase() == START_PHASE
		},
		exec: func(s Story) bool {
			s.AddToInventory(items)
			return true
		},
	}
}

var defaultStoryRules = []StoryRule{
	&storyRule{
		name: "room display description and name on first visit",
		match: func(s Story) bool {
			return s.Phase() == RENDER_PHASE
		},
		exec: func(s Story) bool {
			// FIX
			return false
		},
	},
}