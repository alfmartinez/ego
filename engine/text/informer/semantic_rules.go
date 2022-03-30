package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
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
				e := FindPropertyByValue(value)
				if e == nil {
					panic(fmt.Errorf("can't find property with value %q", value))
				}
				kind.Set(e.Name(), value+":"+def.Certainty)
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
		"create object / set property",
		func(s *grammar.Statement) bool {
			return s.Sentence != nil && s.Sentence.DP != nil && s.Sentence.DP.Designator != nil && s.Sentence.VP.Verb == "is"
		},
		func(s *grammar.Statement, r Semantix) {
			designator := s.Sentence.VP.DP.Designator
			elements := designator.Elements
			var kindKey string
			var properties = make(map[string]ValueKind)
			if _, ok := kinds[designator.Get()]; !ok {
				for _, tag := range elements {
					if _, ok := kinds[tag]; ok {
						kindKey = tag

					}
					if e := FindPropertyByValue(tag); e != nil {
						properties[tag] = e
					}
				}
			} else {
				if _, ok := kinds[designator.Get()]; ok {
					kindKey = designator.Get()
				}
			}
			var object Object
			if kindKey == "" {
				object = r.GetObject(s.Sentence.DP.Designator.Get())

			} else {
				object = CreateObject(kindKey)
				object.Set("name", s.Sentence.DP.Designator.Get())
				r.AddObject(object)
			}
			for value, property := range properties {
				object.Set(property.Name(), value)
			}
			if s.Sentence.Description != "" {
				object.Set("description", s.Sentence.Description)
			}
			if object.IsKind("room") {
				r.SetLastRoom(object)
			}
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
	CreateSemanticRule(
		"create room relative to current",
		func(s *grammar.Statement) bool {
			return s.RelativeRoom != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.RelativeRoom
			room := CreateObject(def.Kind.Get())
			r.AddObject(room)
			if def.Name != nil {
				room.Set("printed name", def.Name.GetCase())
				room.Set("name", def.Name.Get())
			}
			for _, property := range def.With {
				room.Set(property.Property.Get(), property.Value)
			}
			origin := r.LastRoom()
			direct := def.Direction.Direct()
			inverse := def.Direction.Reverse()
			rule1 := CreateConnectorRule(origin, room, direct)
			rule2 := CreateConnectorRule(room, origin, inverse)
			r.AddStoryRule(rule1)
			r.AddStoryRule(rule2)
		},
	),
	CreateSemanticRule(
		"Create Thing In Place",
		func(s *grammar.Statement) bool {
			return s.CreateInPlace != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.CreateInPlace
			elements := def.Thing.Elements
			var kindKey string
			var properties = make(map[string]ValueKind)
			for _, value := range elements {
				p := FindPropertyByValue(value)
				if p != nil {
					properties[value] = p
				} else {
					kindKey = value
				}
			}
			if kindKey == "" {
				panic("Missing kind key")
			}
			var o Object
			if _, ok := kinds[kindKey]; !ok {
				o = CreateObject("thing")
				o.Set("name", def.Thing.Get())
			} else {
				o = CreateObject(kindKey)
				for value, p := range properties {
					o.Set(p.Name(), value)
				}
			}
			r.AddObject(o)
			dest := r.GetObject(def.Place.Get())
			if dest == nil {
				panic(fmt.Errorf("missing place %s", def.Place.Get()))
			}
			rule := CreateAddItemToRoomRule(dest, o)
			r.AddStoryRule(rule)
		},
	),
	CreateSemanticRule(
		"Add property to last added thing",
		func(s *grammar.Statement) bool {
			return s.QuickProperty != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.QuickProperty
			o := r.LastThing()
			for _, value := range def.Values {
				p := FindPropertyByValue(value)
				o.Set(p.Name(), value)
			}
		},
	),
}
