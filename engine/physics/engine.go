package physics

import (
	"github.com/alfmartinez/ego/engine/configuration"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/physics/module"
	"github.com/alfmartinez/ego/engine/slices"
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

	slices.Apply(moduleList, func(name string) {
		module := module.CreateModule(name)
		modules = append(modules, module)
	})

	return &phyiscsEngine{modules}
}

type phyiscsEngine struct {
	modules []module.Module
}

func (e *phyiscsEngine) Init() {
	slices.Apply(e.modules, func(x module.Module) {
		x.Init()
	})
}

func (e *phyiscsEngine) Add(i interface{}) {
	slices.Apply(e.modules, func(x module.Module) {
		x.Add(i)
	})
}

func (e *phyiscsEngine) Advance(dt time.Duration) {
	slices.Apply(e.modules, func(x module.Module) {
		x.Advance(dt.Seconds())
	})
}
