package physics

import "fmt"

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
}
