package module

import (
	"fmt"
	"time"
)

type moduleFactory = func() Module

var moduleFactories = make(map[string]moduleFactory)

func RegisterModuleFactory(name string, f moduleFactory) {
	moduleFactories[name] = f
}

func CreateModule(name string) Module {
	factoryFunc, ok := moduleFactories[name]
	if !ok {
		panic(fmt.Errorf("missing module factory %s", name))
	}
	return factoryFunc()
}

type Module interface {
	Init()
	Add(interface{})
	Advance(time.Duration, []interface{}) []interface{}
}

type Modules []Module

func (m Modules) Map(f func(x Module) interface{}) []interface{} {
	results := make([]interface{}, 0)
	for _, m := range m {
		results = append(results, f(m))
	}
	return results
}

func (a Modules) Apply(f func(m Module)) {
	for _, m := range a {
		f(m)
	}
}
