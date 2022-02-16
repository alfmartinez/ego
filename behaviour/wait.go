package behaviour

type wait struct {
	behaviour
}

func (behaviour *wait) GetName() string {
	return behaviour.Name
}

func (behaviour *wait) Evaluate() string {
	if behaviour.behaviour.data.duration == 0 {
		behaviour.SetDuration(1)
	} else {
		if behaviour.decrementDuration() == 0 {
			return "idle"
		}
	}
	return "wait"
}
