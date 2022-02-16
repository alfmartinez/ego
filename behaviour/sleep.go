package behaviour

type sleep struct {
	behaviour
}

func (behaviour *sleep) GetName() string {
	return behaviour.Name
}

func (behaviour *sleep) Evaluate() {

}

func (behaviour *sleep) IsOver() bool {
	return true
}

func (behaviour *sleep) Reset() {}

func (behaviour *sleep) Next() string {
	return "idle"
}
