package informer

import (
	"fmt"
)

type (
	Object interface {
		Set(string, string)
		Get(string) string
		IsKind(string) bool
	}

	object struct {
		kind       ObjectKind
		properties map[string]string
	}
)

func CreateObject(kindKey string) Object {
	kind, ok := kinds[kindKey]
	if !ok {
		panic(fmt.Errorf("unknown kind %s", kindKey))
	}
	defaults := kind.Defaults()
	return &object{
		kind:       kind.Clone(),
		properties: defaults,
	}
}

func (o *object) Set(name string, value string) {
	if eValue, enforce := o.properties[name+"!"]; enforce && eValue != value {
		panic(fmt.Errorf("cannot set property %s, always enforced with value %s", name, eValue))
	}
	o.properties[name] = value
}

func (o *object) Get(name string) string {
	return o.properties[name]
}

func (o *object) IsKind(kind string) bool {
	return o.kind.IsKind(kind)
}

func CreateValue(valueKey string) ValueKind {
	return &valueKind{
		name: valueKey,
	}
}
