package memory

import "ego/pkg/movement"

type Interested interface {
	HasInterests() bool
	AddInterests([]movement.Positionnable)
	GetNextInterest() movement.Positionnable
}

type interests struct {
	interests []movement.Positionnable
}

func CreateInterests() Interested {
	var i []movement.Positionnable
	return &interests{i}
}

func (m *interests) HasInterests() bool {
	return len(m.interests) > 0
}

func (m *interests) AddInterests(interests []movement.Positionnable) {
	m.interests = append(m.interests, interests...)
}

func (m *interests) GetNextInterest() movement.Positionnable {
	a := m.interests
	var i movement.Positionnable
	if len(m.interests) == 0 {
		return nil
	}
	if len(m.interests) > 1 {
		i, m.interests = a[0], a[1:]
	} else {
		i = a[0]
		m.interests = nil
	}

	return i

}
