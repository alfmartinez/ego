package behaviour

type sleep struct {
	behaviour
}

func (behaviour *sleep) GetName() string {
	return behaviour.Name
}

func (behaviour *sleep) Evaluate() string {
	return "idle"
}
