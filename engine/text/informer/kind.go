package informer

import (
	"golang.org/x/exp/maps"
	"strings"
)

type (
	ObjectKind interface {
		Set(string, string)
		Get(string) string
		Clone() ObjectKind
		Defaults() map[string]string
		IsKind(string) bool
	}

	ValueKind interface {
		SetValues([]string)
		Values() []string
		Name() string
	}

	objectKind struct {
		parent     ObjectKind
		properties map[string]string
	}

	valueKind struct {
		name   string
		values []string
	}
)

func CreateValue(valueKey string) ValueKind {
	return &valueKind{
		name: valueKey,
	}
}

func (k *objectKind) Clone() ObjectKind {
	return &objectKind{
		parent:     k,
		properties: maps.Clone(k.properties),
	}
}

func (k *objectKind) IsKind(kind string) bool {
	if k.Get("name") == kind {
		return true
	}
	return k.parent != nil && k.parent.IsKind(kind)
}

func (k *objectKind) Set(property, value string) {
	k.properties[property] = value
}

func (k *objectKind) Get(property string) string {
	return k.properties[property]
}

func (k *objectKind) Defaults() map[string]string {
	var defaults = make(map[string]string)
	for key, value := range k.properties {
		parts := strings.Split(value, ":")
		var pValue string
		if len(parts) > 1 {
			pValue = parts[0]
			if parts[1] == "always" {
				defaults[key+"!"] = parts[0]
			}
		} else {
			pValue = value
		}
		defaults[key] = pValue
	}
	return defaults
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
