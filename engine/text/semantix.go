package text

import (
	"fmt"
	"strings"
)

type Semantix interface {
	BuildStory(*Grammar) Story
}

func CreateSemantix() Semantix {
	return &semantix{
		rooms: make(map[string]Room),
	}
}

type semantix struct {
	title    string
	rooms    map[string]Room
	lastRoom Room
	startAt  string
	tests    []string
}

func (s *semantix) BuildStory(ast *Grammar) Story {
	for _, statement := range ast.World.Statements {
		s.analyze(statement)
	}
	return &story{
		title:   s.title,
		rooms:   s.rooms,
		current: s.startAt,
		tests:   s.tests,
	}
}

func (s *semantix) analyze(t *Statement) {
	switch {
	case t.Title != "":
		s.analyzeTitle(t.Title)
	case t.Sentence != nil:
		s.analyzeSentence(t.Sentence)
	case t.Test != nil:
		s.analyzeTest(t.Test)
	case t.Direction != nil:
		s.analyzeDirection(t.Direction)
	}
}

func (s *semantix) analyzeTitle(value string) {
	if s.title != "" {
		panic(fmt.Errorf("title already set"))
	}
	s.title = value
}

func (s *semantix) analyzeSentence(value *Sentence) {
	switch value.VP.DP.Designator.Elements[0] {
	case "room":
		s.createRoom(value)
	}

}

func (s *semantix) analyzeTest(t *Test) {
	s.tests = append(s.tests, t.Commands...)
}

func (s *semantix) analyzeDirection(t *Connector) {
	direction := strings.ToLower(t.Direction.Value)
	originKey := strings.Join(t.Origin.Elements, " ")
	targetKey := strings.Join(t.Target.Elements, " ")
	var origin, target Room
	var ok bool
	origin, ok = s.rooms[originKey]
	if !ok {
		origin = CreateRoomFromName(originKey)
		s.rooms[originKey] = origin
	}
	target, ok = s.rooms[targetKey]
	if !ok {
		target = CreateRoomFromName(targetKey)
		s.rooms[targetKey] = target
	}
	origin.AddDirection(direction, targetKey)
	if t.Description != "" {
		target.SetDescription(t.Description)
	}
}

func (s *semantix) createRoom(value *Sentence) {
	room := CreateRoomFromSentence(value)
	name := room.Name()
	s.rooms[name] = room
	if s.startAt == "" {
		s.startAt = name
	}

}
