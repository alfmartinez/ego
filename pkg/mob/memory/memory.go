package memory

import (
	"ego/pkg/mob/movement"
	"image"
)

type Memory interface {
	HasInterests() bool
	AddInterests([]movement.Positionnable)
	HasExplored(movement.Positionnable) bool
	Explore(movement.Positionnable) bool
	GetNextInterest() movement.Positionnable
}

type memory struct {
	places    map[image.Point]PlaceMemory
	interests []movement.Positionnable
}

func CreateMemory() Memory {
	places := make(map[image.Point]PlaceMemory)
	var interests []movement.Positionnable
	return &memory{places, interests}
}

func (m *memory) HasInterests() bool {
	return len(m.interests) > 0
}

func (m *memory) AddInterests(interests []movement.Positionnable) {
	m.interests = append(m.interests, interests...)
}

func (m *memory) GetNextInterest() movement.Positionnable {
	a := m.interests
	var i movement.Positionnable
	i, m.interests = a[0], a[1:]
	return i

}

func (m *memory) HasExplored(location movement.Positionnable) bool {
	p, ok := m.places[location.Position()]
	return ok && p.IsExplored()
}

func (m *memory) Explore(location movement.Positionnable) bool {
	place, ok := m.places[location.Position()]
	if !ok {
		place = CreatePlaceMemory()
		m.places[location.Position()] = place
	}
	place.Explore()
	return place.IsExplored()
}
