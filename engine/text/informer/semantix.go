package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
)

type Semantix interface {
	BuildStory(*grammar.Grammar)
	GetStory() Story
	AddStoryRule(StoryRule)
	AddObject(Object)
	AddTest([]string)
	GetObject(string) Object
	SetLastRoom(Object)
	LastRoom() Object
	LastThing() Object
	Debug() bool
}

func CreateRuleSemantix(debug bool) Semantix {
	return &semantix{
		debug:      debug,
		semRules:   semRules,
		storyRules: defaultStoryRules,
	}
}

type semantix struct {
	debug      bool
	semRules   []SemanticRule
	objects    []Object
	storyRules []StoryRule
	tests      []string
	lastRoom   Object
	lastThing  Object
}

func (r *semantix) Debug() bool {
	return r.debug
}

func (r *semantix) SetLastRoom(o Object) {
	r.lastRoom = o
}

func (r *semantix) LastRoom() Object {
	return r.lastRoom
}

func (r *semantix) LastThing() Object {
	return r.lastThing
}

func (r *semantix) BuildStory(g *grammar.Grammar) {
	for _, statement := range g.World.Statements {

		var matched bool
		for _, rule := range r.semRules {
			if rule.Matches(statement) {
				rule.Execute(statement, r)
				matched = true
			}
		}
		if !matched {
			panic(fmt.Errorf("%+v\n", statement))
		}
	}
}

func (s *semantix) GetStory() Story {
	return CreateRuleStory(s.storyRules, s.objects, s.tests)
}

func (s *semantix) AddStoryRule(r StoryRule) {
	s.storyRules = append(s.storyRules, r)
}

func (s *semantix) AddObject(o Object) {
	if o.IsKind("thing") {
		s.lastThing = o
	}
	s.objects = append(s.objects, o)
}

func (s *semantix) GetObject(name string) Object {
	for idx, o := range s.objects {
		if o.Get("name") == name {
			return s.objects[idx]
		}
	}
	return nil
}

func (s *semantix) AddTest(t []string) {
	s.tests = append(s.tests, t...)
}
