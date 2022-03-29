package text

import (
	"strings"
)

type Room interface {
	Name() string
	Description() string
}

func CreateRoom(s *Sentence) Room {
	return &room{
		name:        strings.Join(s.DP.Designator.Elements, " "),
		description: s.Description,
	}
}

type room struct {
	name        string
	description string
}

func (r *room) Name() string {
	return r.name
}

func (r *room) Description() string {
	return r.description
}
