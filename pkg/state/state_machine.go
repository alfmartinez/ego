package state

type StateMachine interface {
	Update()
}

type stateMachine struct {
	current State
	next    State
}

func CreateStateMachine() StateMachine {
	return &stateMachine{}
}

func (m *stateMachine) Update() {
	if m.current == nil {
		m.next = CreateState("idle")
	}
	if m.next != nil {
		m.current = m.next
		m.next = nil
	}
	m.next = m.current.Update(m)
}
