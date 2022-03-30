package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"io"
)

type Phase int

const (
	PRE_START_PHASE Phase = iota
	START_PHASE
	START_TURN_PHASE
	PRE_UPDATE_PHASE
	UPDATE_PHASE
	POST_UPDATE_PHASE
	RENDER_PHASE
	POST_TURN_PHASE
	TURN_ENDED
)

type Story interface {
	Start()
	Test()
	Phase() Phase
	AdvancePhase()
	CurrentRoom() Object
	SetCurrentRoom(Object)
	Command() *grammar.Command
	Say(string)
	SetWriter(io.Writer)
	AddToInventory([]Object)
	AddItemToRoom(item Object, place Object)
}

func CreateRuleStory(storyRules []StoryRule, objects []Object, tests []string) Story {
	return &story{
		phase:    PRE_START_PHASE,
		objects:  objects,
		rules:    storyRules,
		cmdChan:  make(chan *grammar.Command),
		tests:    tests,
		location: make(map[Object]Object),
		contains: make(map[Object][]Object),
	}
}

type story struct {
	phase       Phase
	objects     []Object
	rules       []StoryRule
	tests       []string
	cmdChan     chan *grammar.Command
	currentRoom Object
	command     *grammar.Command
	writer      io.Writer
	inventory   []Object
	location    map[Object]Object
	contains    map[Object][]Object
}

func (s *story) AddToInventory(objects []Object) {
	s.inventory = append(s.inventory, objects...)
}

func (s *story) AddItemToRoom(item Object, place Object) {
	s.location[item] = place
	s.contains[place] = append(s.contains[place], item)
}

func (s *story) Phase() Phase {
	return s.phase
}

func (s *story) AdvancePhase() {
	s.phase = s.phase + 1
	if s.phase == TURN_ENDED {
		s.phase = START_TURN_PHASE
	}
}

func (s *story) SetCurrentRoom(o Object) {
	s.currentRoom = o
}

func (s *story) CurrentRoom() Object {
	return s.currentRoom
}

func (s *story) Command() *grammar.Command {
	return s.command
}

func (s *story) Start() {
	s.ApplyStoryRules()
	s.AdvancePhase()
	s.ApplyStoryRules()
	go func() {
		for cmd := range s.cmdChan {
			s.command = cmd
		}
	}()
}

func (s *story) Test() {
	s.Start()
	for _, cmd := range s.tests {
		command := grammar.ParseCommand(cmd)
		s.cmdChan <- command
	}
	close(s.cmdChan)
}

func (s *story) SetWriter(writer io.Writer) {
	s.writer = writer
}

func (s *story) Say(say string) {
	fmt.Fprintln(s.writer, say)
}

func (s *story) ApplyStoryRules() {
	for _, rule := range s.rules {
		if rule.Matches(s) {
			rule.Execute(s)
		}
	}
}
