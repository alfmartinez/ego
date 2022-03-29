package text

import (
	"strings"
)

type Room interface {
	Name() string
	Description() string
	SetDescription(string)
	AddDirection(direction string, destination string)
	Execute(cmd *Command) string
}

func CreateRoomFromSentence(s *Sentence) Room {
	return &room{
		name:        strings.Join(s.DP.Designator.Elements, " "),
		description: s.Description,
		directions:  make(map[string]string),
	}
}

func CreateRoomFromName(name string) Room {
	return &room{
		name:       name,
		directions: make(map[string]string),
	}
}

type room struct {
	name        string
	description string
	directions  map[string]string
}

func (r *room) Name() string {
	return r.name
}

func (r *room) Description() string {
	return r.description
}

func (r *room) SetDescription(desc string) {
	r.description = desc
}

func (r *room) AddDirection(direction string, destination string) {
	r.directions[direction] = destination
}

func (r *room) Execute(cmd *Command) string {
	direction := cmd.Direction.Value
	if target, ok := r.directions[direction]; ok {
		return target
	}
	return ""
}
