package informer

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/text/grammar"
	"strings"
)

var locationSet bool

var semRules = []SemanticRule{
	CreateSemanticRule(
		"Instanciate object",
		func(s *grammar.Statement) bool {
			return s.Instanciate != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.Instanciate
			o := CreateObject(def.Kind.Get(), def.Name.Get(), def.Name.GetCase())
			r.AddObject(o)
			if def.With != nil {
				o.Set(def.With.Property.Get(), def.With.Value)
			}
		},
	),
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
			rule := CreateWhenRule(START_PHASE, Say(s.Title+"\n\n")).SetName("display title on start rule")
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
				name := s.Sentence.DP.Designator
				object = CreateObject(kindKey, name.Get(), name.GetCase())
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
			target := CreateObject("room", dir.Target.Get(), dir.Target.GetCase())
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
			direct := r.GetObject(s.Direction.Direction.Get())
			inverse := r.GetObject(direct.Get("opposite"))
			rule1 := CreateConnectorRule(origin, target, direct).SetName("connect direct rule")
			rule2 := CreateConnectorRule(target, origin, inverse).SetName("connect inverse rule")
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
				item := CreateObject("thing", itemName, itemName)
				r.AddObject(item)
				objects = append(objects, item)
			}
			rule := CreateWhenRule(START_PHASE, func(s Story) (success, cont bool) {
				s.AddToInventory(objects)
				return true, false
			}).SetName("add items to inventory rule")
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
				o = CreateObject("thing", def.Thing.Get(), def.Thing.GetCase())
				o.Set("name", def.Thing.Get())
			} else {
				o = CreateObject(kindKey, kindKey, kindKey)
				for value, p := range properties {
					o.Set(p.Name(), value)
				}
			}
			r.AddObject(o)
			dest := r.GetObject(def.Place.Get())
			if dest == nil {
				panic(fmt.Errorf("missing place %s", def.Place.Get()))
			}
			rule := CreateWhenRule(START_PHASE,
				func(s Story) (success, cont bool) {
					s.AddItemToRoom(o, dest)
					return true, true
				}).SetName("Add Item to Room")
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
	CreateSemanticRule(
		"Add either or property to kind",
		func(s *grammar.Statement) bool {
			return s.EitherProperty != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.EitherProperty
			parts := []string{def.Kind.Get()}
			parts = append(parts, def.Values...)
			valueKey := strings.Join(parts, "|")
			valueKind := CreateValue(valueKey)
			values[valueKey] = valueKind
			valueKind.SetValues(def.Values)
		},
	),
	CreateSemanticRule(
		"Set text property of kind",
		func(s *grammar.Statement) bool {
			return s.TextPropertyKind != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.TextPropertyKind
			kind := kinds[def.Target.Get()]
			kind.Set(def.PropertyName.Get(), def.Value+":"+def.Certainty)
		},
	),
	CreateSemanticRule(
		"Set text property of object",
		func(s *grammar.Statement) bool {
			return s.TextPropertyObject != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.TextPropertyObject
			o := r.GetObject(def.Target.Get())
			o.Set(def.PropertyName.Get(), def.Value)
		},
	),
	CreateSemanticRule(
		"create text property",
		func(s *grammar.Statement) bool {
			return s.Property != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.Property
			pName := def.Name.Get()
			var property ValueKind = &valueKind{
				name:   pName,
				values: []string{},
			}
			values[pName] = property
			kind := kinds[def.Kind.Get()]
			kind.Set(pName, "")
		},
	),
	CreateSemanticRule(
		"create when activity",
		func(s *grammar.Statement) bool {
			return s.WhenDeClaration != nil
		},
		func(s *grammar.Statement, r Semantix) {
			var ruleFactory func(Activity) StoryRule
			def := s.WhenDeClaration
			phase := GetPhase(def.Condition.Rule.Get())
			if phase == 0 {
				// Not a phase...
				// May be an action ?
				o := r.GetObject(def.Condition.Rule.Elements[0])
				if o == nil {
					panic(fmt.Errorf("Dont know what %q is", def.Condition.Rule.Elements[0]))
				}
				if o.IsKind("action") {
					fmt.Printf("Action detected ! %s \n", o.Get("name"))
					alias := def.Condition.Rule.Elements[1]
					ruleFactory = func(act Activity) StoryRule {
						return CreateActivityRule(o, act, alias)
					}
				} else {
					panic(fmt.Errorf("Dont know what to do with condition %s", def.Condition.Rule.Get()))
				}

			} else {
				ruleFactory = func(act Activity) StoryRule {
					return CreateWhenRule(phase, act)
				}
			}

			rule := ruleFactory(CreateStoryActivity(def.Activity)).SetName(fmt.Sprintf("%s", def.Activity))
			r.AddStoryRule(rule)
		},
	),
	CreateSemanticRule(
		"Define Alias with Understand",
		func(s *grammar.Statement) bool {
			return s.Understand != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.Understand
			o := r.GetObject(def.Target.Get())
			o.AddAlias(def.Alias)
		},
	),
	CreateSemanticRule(
		"Define Action",
		func(s *grammar.Statement) bool {
			return s.ActionDefinition != nil
		},
		func(s *grammar.Statement, r Semantix) {
			def := s.ActionDefinition
			o := CreateObject("action", def.Name.Get(), def.Name.GetCase())
			if def.Target != nil {
				o.Set("applies to", def.Target.Get())
			}
			r.AddObject(o)
		},
	),
}

func CreateStoryActivity(activity *grammar.Activity) Activity {
	var act Activity
	switch {
	case activity.Say != "":
		act = Say(activity.Say)
	case activity.Launch != nil:
		launch := activity.Launch
		action := Action(launch.Action)
		argument := launch.Argument.Get()
		act = Launch(action, argument)
	case len(activity.Activities) > 0:
		var activities []Activity
		for _, a := range activity.Activities {
			activities = append(activities, CreateStoryActivity(a))
		}
		act = Sequence(activities)
	default:
		panic(fmt.Errorf("Don't know %+v", activity))
	}
	return act
}
