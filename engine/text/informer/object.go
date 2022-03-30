package informer

import (
	"fmt"
)

type (
	Object interface {
		Kind
	}

	object struct {
		Kind
	}
)

func CreateObject(kindKey string) Object {
	kind, ok := kinds[kindKey]
	if !ok {
		panic(fmt.Errorf("unknown kind %s", kindKey))
	}
	return &object{
		kind,
	}
}
