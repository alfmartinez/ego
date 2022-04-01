package informer

type ActivityType int

const (
	SAY_ACTIVITY ActivityType = iota
	ROOM_CHANGE_ACTIVITY
)

type Activity func(s Story) (success, cont bool)

func Say(value string) Activity {
	return func(s Story) (success, cont bool) {
		s.Say(value)
		return true, true
	}
}

func Launch(action Action, value string) Activity {
	return func(s Story) (success, cont bool) {
		s.Publish(Message{
			Action:   action,
			Argument: value,
		})
		return true, true
	}
}

func Sequence(activities []Activity) Activity {
	return func(s Story) (success, cont bool) {
		for _, act := range activities {
			if success, _ := act(s); !success {
				return false, false
			}
		}
		return true, true
	}
}
