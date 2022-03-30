package informer

import (
	"golang.org/x/exp/maps"
)

type (
	ObjectKind interface {
		Set(string, string)
		Get(string) string
		Clone() ObjectKind
	}

	ValueKind interface {
		SetValues([]string)
		Values() []string
		Name() string
	}

	objectKind struct {
		parent     ObjectKind
		name       string
		properties map[string]string
	}

	valueKind struct {
		name   string
		values []string
	}
)

func (k *objectKind) Clone() ObjectKind {

	return &objectKind{
		parent:     k,
		properties: maps.Clone(k.properties),
	}
}

func (k *objectKind) Set(property, value string) {
	k.properties[property] = value
}

func (k *objectKind) Get(property string) string {
	return k.properties[property]
}

func (v *valueKind) SetValues(values []string) {
	v.values = values
}

func (v *valueKind) Values() []string {
	return v.values
}

func (v *valueKind) Name() string {
	return v.name
}
