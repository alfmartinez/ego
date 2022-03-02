package motivator

type NeedLevel interface {
	Need
	Value() int
	Update()
	AddIncrement(LevelIncrement)
}

type needLevel struct {
	Need
	value      int
	increments []LevelIncrement
}

func CreateNeedLevel(need Need, value int) NeedLevel {
	increments := make([]LevelIncrement, 0)
	return &needLevel{need, value, increments}
}

func (n *needLevel) Value() int {
	return n.value
}

func (n *needLevel) Update() {
	n.applyIncrements()
	n.removeExpiredIncrements()
	n.enforceBoundaries()
}

func (n *needLevel) enforceBoundaries() {
	switch {
	case n.value > 100:
		n.value = 100
	case n.value < 0:
		n.value = 0

	}
}

func (n *needLevel) applyIncrements() {
	for _, inc := range n.increments {
		n.value += inc.Apply()
	}
}

func (n *needLevel) removeExpiredIncrements() {
	var remains = make([]LevelIncrement, 0)
	for _, inc := range n.increments {
		if !inc.IsOver() {
			remains = append(remains, inc)
		}
	}
	n.increments = remains
}

func (n *needLevel) AddIncrement(increment LevelIncrement) {
	n.increments = append(n.increments, increment)
}
