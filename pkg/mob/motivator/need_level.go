package motivator

type NeedLevel interface {
	Update()
	Provide(increment int, duration int)
}

type needLevel struct {
	need      Need
	value     int
	increment int
	duration  int
}

func CreateNeedLevel(need Need, value int) NeedLevel {
	return &needLevel{need, value, 0, 0}
}

func (n *needLevel) Update() {
	if n.duration > 0 {
		n.value += n.increment
		n.duration--
	} else {
		n.value--
	}

	switch {
	case n.value > 100:
		n.value = 100
	case n.value < 0:
		n.value = 0

	}
}

func (n *needLevel) Provide(increment int, duration int) {
	n.increment = increment
	n.duration = duration
}
