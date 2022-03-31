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
