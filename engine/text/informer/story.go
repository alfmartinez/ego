package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"io"
	"strings"
)

type Phase int

const (
	NONE Phase = iota
	START_PHASE
	START_TURN_PHASE
	UPDATE_PHASE
	RENDER_PHASE
	TURN_ENDED
)

var phaseLexic = map[string]Phase{
	"play begins": START_PHASE,
	"turn starts": START_TURN_PHASE,
}

func GetPhase(value string) Phase {
	return phaseLexic[value]
}

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
	GetObject(string) Object
}

func CreateRuleStory(publisher Publisher, objects []Object, tests []string) Story {
	index := make(map[string]Object)
	for _, o := range objects {
		index[o.Get("name")] = o
		for _, alias := range o.Aliases() {
			index[alias] = o
		}
	}
	return &story{
		phase:     NONE,
		index:     index,
		publisher: publisher,
		cmdChan:   make(chan *grammar.Command),
		tests:     tests,
		location:  make(map[Object]Object),
		contains:  make(map[Object][]Object),
	}
}

type story struct {
	phase     Phase
	index     map[string]Object
	publisher Publisher
	tests     []string
	cmdChan   chan *grammar.Command
	command   *grammar.Command
	writer    io.Writer
	inventory []Object
	location  map[Object]Object
	contains  map[Object][]Object
	test      bool
	cmdText   string
	stop      bool
}

func (s *story) GetObject(key string) Object {
	return s.index[key]
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
	s.publisher.Publish(Message{
		Story: s,
		Phase: s.phase,
	})
}

func (s *story) SetCurrentRoom(o Object) {
	s.index["location"] = o
}

func (s *story) CurrentRoom() Object {
	return s.index["location"]
}

func (s *story) Command() *grammar.Command {
	return s.command
}

func (s *story) Start() {
	go s.startLoop()
}

func (s *story) startLoop() {
	var more = true
	s.AdvancePhase() // START PLAY
	for more {
		s.AdvancePhase() // START TURN
		s.command, more = <-s.cmdChan
		s.AdvancePhase() // UPDATE
		s.AdvancePhase() // RENDER
	}
}

func (s *story) Test() {
	s.test = true
	go func() {
		for _, cmd := range s.tests {
			s.cmdText = cmd
			command := grammar.ParseCommand(cmd)
			s.cmdChan <- command
		}
		close(s.cmdChan)
	}()

	s.Start()
}

func (s *story) SetWriter(writer io.Writer) {
	s.writer = writer
}

func (s *story) Say(say string) {
	replacer := s.buildReplacer()
	replacer.WriteString(s.writer, say)
}

func (s *story) buildReplacer() *strings.Replacer {
	var oldNew = []string{}
	for key, o := range s.index {
		if !o.Has("printed name") {
			panic(fmt.Errorf("Object %s has no printed name", o))
		}
		oldNew = append(oldNew, "["+key+"]", o.Get("printed name"))
	}
	return strings.NewReplacer(oldNew...)
}
