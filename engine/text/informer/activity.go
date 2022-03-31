package informer

type ActivityType int

const (
	SAY_ACTIVITY ActivityType = iota
	ROOM_CHANGE_ACTIVITY
)

type Activity func(s Story) bool

func Say(value string) Activity {
	return func(s Story) bool {
		s.Say(value)
		return true
	}
}

func Launch(action Action, value string) Activity {
	return func(s Story) bool {
		s.Publish(Message{
			Action:   action,
			Argument: value,
		})
		return true
	}
}

func Sequence(activities []Activity) Activity {
	return func(s Story) bool {
		for _, act := range activities {
			if !act(s) {
				return false
			}
		}
		return true
	}
}
