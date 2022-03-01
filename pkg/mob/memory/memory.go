package memory

import (
	"ego/pkg/mob/movement"
	"image"
)

type Memory interface {
	HasInterests() bool
	AddInterests([]movement.Positionnable)
	HasExplored(movement.Positionnable) bool
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

func (m *memory) HasExplored(location movement.Positionnable) bool {
	p, ok := m.places[location.Position()]
	return ok && p.IsExplored()
}
