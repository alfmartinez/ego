package informer

type ActivityType int

const (
	SAY_ACTIVITY ActivityType = iota
	ROOM_CHANGE_ACTIVITY
)

type Activity func(s Story) RuleResult

func Say(value string) Activity {
	return func(s Story) RuleResult {
		s.Say(value)
		return RULE_UNDECIDED
	}
}

func Launch(action Action, value string) Activity {
	return func(s Story) RuleResult {
		s.Publish(Message{
			Action:   action,
			Argument: value,
		})
		return RULE_UNDECIDED
	}
}

func Sequence(activities []Activity) Activity {
	return func(s Story) RuleResult {
		for _, act := range activities {
			if result := act(s); result != RULE_UNDECIDED {
				return result
			}
		}
		return RULE_UNDECIDED
	}
}
