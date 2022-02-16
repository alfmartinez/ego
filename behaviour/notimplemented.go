package behaviour

type notimplemented struct {
	behaviour
}

func (behaviour *notimplemented) GetName() string {
	return behaviour.Name
}

func (behaviour *notimplemented) Evaluate() {}

func (behaviour *notimplemented) Reset() {}

func (behaviour *notimplemented) IsOver() bool {
	return false
}

func (behaviour *notimplemented) Next() string {
	return "idle"
}
