package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
)

type Semantix interface {
	BuildStory(*grammar.Grammar) Story
	AddStoryRule(StoryRule)
	AddObject(Object)
	AddTest([]string)
	GetObject(string) Object
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
}

func (r *semantix) Debug() bool {
	return r.debug
}

func (r *semantix) BuildStory(g *grammar.Grammar) Story {
	//r.initStory()
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
	return CreateRuleStory(r.storyRules, r.objects, r.tests)
}

func (s *semantix) AddStoryRule(r StoryRule) {
	s.storyRules = append(s.storyRules, r)
}

func (s *semantix) AddObject(o Object) {
	s.objects = append(s.objects, o)
}

func (s *semantix) GetObject(name string) Object {
	for _, o := range s.objects {
		if o.Name() == name {
			return o
		}
	}
	return nil
}

func (s *semantix) AddTest(t []string) {
	s.tests = append(s.tests, t...)
}
