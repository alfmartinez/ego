package informer

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
	if r.match(msg) {
		s := msg.Story
		r.exec(s)
	}
}

func (r *storyRule) Name() string {
	return r.name
}

func CreateConnectorRule(o Object, t Object, direction string) StoryRule {
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

func CreateWhenRule(when Phase, f Activity) StoryRule {
	return &storyRule{
		name: "when phase say rule",
		match: func(msg Message) bool {
			return msg.Phase == when
		},
		exec: f,
	}
}

func CreateActivityRule(o Object, f Activity) StoryRule {
	return &storyRule{
		name: "when activity say rule",
		match: func(msg Message) bool {
			return false
		},
		exec: f,
	}
}

func CreatePlayerInventoryRule(items []Object) StoryRule {
	return &storyRule{
		name: "put items in inventory rule",
		match: func(msg Message) bool {
			return msg.Phase == START_PHASE
		},
		exec: func(s Story) bool {
			s.AddToInventory(items)
			return true
		},
	}
}

func CreateAddItemToRoomRule(dest Object, o Object) StoryRule {
	return &storyRule{
		name: "add item to place rule",
		match: func(msg Message) bool {
			return msg.Phase == START_PHASE
		},
		exec: func(s Story) bool {
			s.AddItemToRoom(o, dest)
			return true
		},
	}
}
