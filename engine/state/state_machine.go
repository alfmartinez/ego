package state

import "time"

type StateMachine interface {
	DoUpdate(Updatable, time.Duration)
	SetState(string)
	SetStates(States)
}

type stateMachine struct {
	states  States
	current string
	next    string
}

func CreateStateMachine() StateMachine {
	return &stateMachine{}
}

func (m *stateMachine) SetStates(states States) {
	m.states = states
}

func (m *stateMachine) DoUpdate(self Updatable, dt time.Duration) {
	if m.current == "" {
		m.next = "default"
	}
	if m.next != "" {
		m.current = m.next
		m.next = ""
	}
	m.next = m.states[m.current](dt)
}

func (m *stateMachine) SetState(t string) {
	if m.current == "" || t != m.current {
		m.current = t
	}
}
