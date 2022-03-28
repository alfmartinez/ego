package memory

import "github.com/alfmartinez/ego/engine/movement"

type Interested interface {
	HasInterests() bool
	AddInterests([]movement.Location)
	GetNextInterest() (interest movement.Location, found bool)
}

type interests struct {
	interests []movement.Location
}

func CreateInterests() Interested {
	var i []movement.Location
	return &interests{i}
}

func (m *interests) HasInterests() bool {
	return len(m.interests) > 0
}

func (m *interests) AddInterests(interests []movement.Location) {
	m.interests = append(m.interests, interests...)
}

func (m *interests) GetNextInterest() (interest movement.Location, found bool) {
	a := m.interests
	var i movement.Location
	if len(m.interests) == 0 {
		return movement.Location{}, false
	}
	if len(m.interests) > 1 {
		i, m.interests = a[0], a[1:]
	} else {
		i = a[0]
		m.interests = nil
	}

	return i, true

}
