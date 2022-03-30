package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"golang.org/x/exp/slices"
	"strings"
)

var semRules = []SemanticRule{
	CreateSemanticRule(
		"add certainty property to kind",
		func(s *grammar.Statement) bool {
			return s.Certainty != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.Certainty
			kindKey := strings.ToLower(def.Name.Get())
			kind := kinds[kindKey].(ObjectKind)
			for _, value := range def.Values {
				var found bool = false
				for _, e := range values {
					if slices.Contains(e.Values(), value) {
						found = true
						kind.Set(e.Name(), value+":"+def.Certainty)
					}
				}
				if !found {
					panic(fmt.Errorf("can't find property with value %q", value))
				}
			}
		},
	),
	CreateSemanticRule(
		"create a new value kind",
		func(s *grammar.Statement) bool {
			return s.ValueDefinition != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.ValueDefinition
			if r.Debug() {
				fmt.Printf("Creating %s as kind of value. \n", def.Name.Get())
			}
			valueKey := strings.ToLower(def.Name.Get())
			newValue := CreateValue(valueKey)
			values[valueKey] = newValue
			newValue.SetValues(def.With)
		},
	),
	CreateSemanticRule(
		"create a new object kind",
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
			newKind := parent.Clone()
			newKind.Set("name", newKindKey)
			kinds[newKindKey] = newKind
			for _, property := range s.KindDefinition.With {
				newKind.Set(property.Property.Get(), property.Value)
			}
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
			kindKey := strings.ToLower(s.Sentence.VP.DP.Designator.Get())
			object := CreateObject(kindKey)
			object.Set("name", s.Sentence.DP.Designator.Get())
			if s.Sentence.Description != "" {
				object.Set("description", s.Sentence.Description)
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
			target.Set("name", dir.Target.Get())
			if dir.Description != "" {
				target.Set("description", dir.Description)
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
				key := i.Get()
				if key != "" {
					ids = append(ids, i.Get())
				}
			}
			ids = append(ids, s.Sentence.VP.DP.List.Last.Get())

			var objects []Object
			for _, itemName := range ids {
				item := CreateObject("thing")
				item.Set("name", itemName)
				r.AddObject(item)
				objects = append(objects, item)
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
			o := r.GetObject(name)
			if o == nil {
				panic(fmt.Errorf("object should have been created : %s", name))
			}
			o.Set("description", s.Description.Description)
		},
	),
}
