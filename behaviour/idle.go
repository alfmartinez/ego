package behaviour

// IDLE
type idle struct {
	behaviour
}

func (behaviour *idle) GetName() string {
	return behaviour.Name
}

func (behaviour *idle) Evaluate() {}

func (behaviour *idle) IsOver() bool {
	return true
}

func (behaviour *idle) Reset() {}

func (behaviour *idle) Next() string {
	return "wait"
}
