package state

import (
	"fmt"
	"time"
)

type States map[string]func(dt time.Duration) string

type StatesClosure func(m any) States

var statesFactories = make(map[string]StatesClosure)

func RegisterStatesClosure(name string, states StatesClosure) {
	statesFactories[name] = states
}

func CreateStates(name string, m any) States {
	f, ok := statesFactories[name]
	if !ok {
		panic(fmt.Errorf("Unknown states factory %s", name))
	}
	return f(m)
}