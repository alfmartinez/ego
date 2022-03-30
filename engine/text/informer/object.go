package informer

import (
	"fmt"
)

type (
	Object interface {
		ObjectKind
	}

	object struct {
		ObjectKind
	}
)

func CreateObject(kindKey string) Object {
	kind, ok := kinds[kindKey]
	if !ok {
		panic(fmt.Errorf("unknown kind %s", kindKey))
	}
	return &object{
		kind.(ObjectKind),
	}
}

func CreateValue(valueKey string) ValueKind {
	return &valueKind{
		name: valueKey,
	}
}
