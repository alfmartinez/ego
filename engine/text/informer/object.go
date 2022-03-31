package informer

import (
	"fmt"
)

type (
	Object interface {
		Set(string, string)
		Get(string) string
		IsKind(string) bool
		Has(string) bool
	}

	object struct {
		kind       ObjectKind
		properties map[string]string
	}
)

func CreateObject(kindKey string, name string, printed string) Object {
	kind, ok := kinds[kindKey]
	if !ok {
		panic(fmt.Errorf("unknown kind %s", kindKey))
	}
	defaults := kind.Defaults()
	o := &object{
		kind:       kind.Clone(),
		properties: defaults,
	}
	if name != "" {
		o.Set("name", name)
	}
	if printed != "" {
		o.Set("printed name", printed)
	}
	return o
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

func (o *object) Has(property string) bool {
	_, ok := o.properties[property]
	return ok
}

func CreateValue(valueKey string) ValueKind {
	return &valueKind{
		name: valueKey,
	}
}
