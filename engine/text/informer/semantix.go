package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"strings"
)

type Semantix interface {
	BuildStory(*grammar.Grammar)
	GetStory() Story
	AddRulebook(string)
	AddStoryRule(StoryRule)
	AddObject(Object)
	AddTest([]string)
	GetObject(string) Object
	SetLastRoom(Object)
	LastRoom() Object
	LastThing() Object
	IsDirection(string) bool
	Debug() bool
}

func CreateRuleSemantix(debug bool) Semantix {
	locationSet = false
	return &semantix{
		debug:     debug,
		semRules:  semRules,
		rulebooks: CreateRulebooks(),
		index:     make(map[string]Object),
	}
}

type semantix struct {
	debug      bool
	semRules   []SemanticRule
	index      map[string]Object
	storyRules []StoryRule
	tests      []string
	lastRoom   Object
	lastThing  Object
	rulebooks  Rulebooks
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

func (r *semantix) IsDirection(value string) bool {
	key := strings.ToLower(value)
	o := r.GetObject(key)
	return o != nil && o.IsKind("direction")
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
			panic(fmt.Errorf("%v\n", statement))
		}
	}
}

func (s *semantix) GetStory() Story {
	return CreateRuleStory(s.rulebooks, s.index, s.tests)
}

func (s *semantix) AddStoryRule(r StoryRule) {
	s.rulebooks.Register(r)
}

func (s *semantix) AddRulebook(book string) {
	s.rulebooks.AddRulebook(book)
}

func (s *semantix) AddObject(o Object) {
	if o.IsKind("thing") {
		s.lastThing = o
	}
	key := o.Get("name")
	if _, ok := s.index[key]; ok {
		panic(fmt.Errorf("index already has key %q", key))
	}
	s.index[key] = o
}

func (s *semantix) GetObject(name string) Object {
	return s.index[name]
}

func (s *semantix) AddTest(t []string) {
	s.tests = append(s.tests, t...)
}
