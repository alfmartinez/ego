package behaviour

// IDLE
type idle struct {
	behaviour
}

func (behaviour *idle) GetName() string {
	return behaviour.Name
}

func (behaviour *idle) Evaluate() string {
	return "wait"
}
