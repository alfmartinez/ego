package physics

import (
	"ego/engine/configuration"
	"ego/engine/context"
	"ego/engine/physics/module"
	"time"
)

const (
	NUM_STEPS    = 100
	SCALE_FACTOR = 1
)

func FromContext() Engine {
	return context.GetContext().Get("physics").(Engine)
}

func CreatePhysicsEngine() Engine {
	cfg := configuration.FromContext()
	var moduleList []string
	var modules []module.Module

	cfg.UnmarshalKey("physics.modules", &moduleList)

	for _, name := range moduleList {
		module := module.CreateModule(name)
		modules = append(modules, module)
	}
	return &phyiscsEngine{modules}
}

type phyiscsEngine struct {
	modules module.Modules
}

func (e *phyiscsEngine) Init() {
	e.modules.Apply(func(x module.Module) {
		x.Init()
	})
}

func (e *phyiscsEngine) Add(i interface{}) {
	e.modules.Apply(func(x module.Module) {
		x.Add(i)
	})
}

func (e *phyiscsEngine) Advance(dt time.Duration) {
	var results = make([]interface{}, 0)
	e.modules.Apply(func(x module.Module) {
		results = x.Advance(dt, results)
	})
}
