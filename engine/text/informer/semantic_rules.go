package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"strings"
)

var semRules = []SemanticRule{
	CreateSemanticRule(
		"create a new kind",
		func(s *grammar.Statement) bool {
			return s.KindDefinition != nil
		},
		func(s *grammar.Statement, r Semantix) {
			if r.Debug() {
				fmt.Printf("Creating %s as kind of %s. \n", s.KindDefinition.Name.Get(), s.KindDefinition.Parent.Get())
			}
			parentKey := strings.ToLower(s.KindDefinition.Parent.Get())
			newKindKey := strings.ToLower(s.KindDefinition.Name.Get())
			parent, ok := kinds[parentKey]
			if !ok {
				panic(fmt.Errorf("cannot inherit from unknown kind %s", parentKey))
			}
			newKind := &kind{
				parent: parent,
				properties: map[string]string{
					"name":        newKindKey,
					"description": "",
				},
			}
			kinds[newKindKey] = newKind
		},
	),
	CreateSemanticRule(
		"display title on start",
		func(s *grammar.Statement) bool {
			return s.Title != ""
		},
		func(s *grammar.Statement, r Semantix) {
			rule := CreateWhenSayRule(START_PHASE, s.Title)
			r.AddStoryRule(rule)
		},
	),
	CreateSemanticRule(
		"create object",
		func(s *grammar.Statement) bool {
			return s.Sentence != nil && s.Sentence.DP != nil && s.Sentence.DP.Designator != nil && s.Sentence.VP.Verb == "is"
		},
		func(s *grammar.Statement, r Semantix) {
			kindKey := s.Sentence.VP.DP.Designator.Get()
			object := CreateObject(kindKey)
			object.SetName(s.Sentence.DP.Designator.Get())
			if s.Sentence.Description != "" {
				object.SetDescription(s.Sentence.Description)
			}
			r.AddObject(object)
		},
	),
	CreateSemanticRule(
		"create room at the other side of connector",
		func(s *grammar.Statement) bool {
			return s.Direction != nil
		},
		func(s *grammar.Statement, r Semantix) {
			dir := s.Direction
			target := CreateObject("room")
			target.SetName(dir.Target.Get())
			if dir.Description != "" {
				target.SetDescription(dir.Description)
			}
			r.AddObject(target)
		},
	),
	CreateSemanticRule(
		"create connections between rooms",
		func(s *grammar.Statement) bool {
			return s.Direction != nil
		},
		func(s *grammar.Statement, r Semantix) {
			origin := r.GetObject(s.Direction.Origin.Get())
			target := r.GetObject(s.Direction.Target.Get())
			direct := s.Direction.Direction.Direct()
			inverse := s.Direction.Direction.Reverse()
			rule1 := CreateConnectorRule(origin, target, direct)
			rule2 := CreateConnectorRule(target, origin, inverse)
			r.AddStoryRule(rule1)
			r.AddStoryRule(rule2)

		},
	),
	CreateSemanticRule(
		"set tests",
		func(s *grammar.Statement) bool {
			return s.Test != nil
		},
		func(s *grammar.Statement, r Semantix) {
			r.AddTest(s.Test.Commands)
		},
	),
	CreateSemanticRule(
		"player inventory",
		func(s *grammar.Statement) bool {
			return s.Sentence != nil && s.Sentence.DP.Designator.Get() == "player" && s.Sentence.VP.Verb == "carries"
		},
		func(s *grammar.Statement, r Semantix) {
			ids := make([]string, 0)
			ids = append(ids, s.Sentence.VP.DP.Designator.Get())
			for _, i := range s.Sentence.VP.DP.List.Elements {
				ids = append(ids, i.Get())
			}
			ids = append(ids, s.Sentence.VP.DP.List.Last.Get())
			var objects []Object
			for _, itemName := range ids {
				if itemName != "" {
					item := CreateObject("thing")
					item.SetName(itemName)
					r.AddObject(item)
					objects = append(objects, item)
				}
			}
			rule := CreatePlayerInventoryRule(objects)
			r.AddStoryRule(rule)
		},
	),
	CreateSemanticRule(
		"set description",
		func(s *grammar.Statement) bool {
			return s.Description != nil
		},
		func(s *grammar.Statement, r Semantix) {
			name := s.Description.Target.Get()
			object := r.GetObject(name)
			if object == nil {
				panic(fmt.Errorf("object should have been created : %s", name))
			}
			object.SetDescription(s.Description.Description)
		},
	),
}
