package state

type StateMachine interface {
	DoUpdate(Updatable)
	SetState(StateType)
}

type stateMachine struct {
	currentType StateType
	current     State
	next        State
}

func CreateStateMachine() StateMachine {
	return &stateMachine{}
}

func (m *stateMachine) DoUpdate(self Updatable) {
	if m.current == nil {
		m.next = CreateState(StateIdle)
	}
	if m.next != nil {
		m.current = m.next
		m.next = nil
	}
	m.next = m.current.Update(self)
}

func (m *stateMachine) SetState(t StateType) {
	if m.current == nil || t != m.current.Type() {
		m.current = CreateState(t)
	}
}
