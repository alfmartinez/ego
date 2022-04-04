package informer

type ActivityType int

const (
	SAY_ACTIVITY ActivityType = iota
	ROOM_CHANGE_ACTIVITY
)

type Activity func() RuleResult

func CreateSay(value string) Activity {
	return func() RuleResult {
		Say(value)
		return RULE_UNDECIDED
	}
}

func Launch(action Action, value string) Activity {
	return func() RuleResult {
		msg := Try{
			Action:   action,
			Argument: value,
		}
		return Publish(msg)
	}
}

func Sequence(activities []Activity) Activity {
	return func() RuleResult {
		for _, act := range activities {
			if result := act(); result != RULE_UNDECIDED {
				return result
			}
		}
		return RULE_UNDECIDED
	}
}
