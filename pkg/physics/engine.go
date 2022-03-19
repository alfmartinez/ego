package physics

import (
	"ego/pkg/configuration"
	"ego/pkg/context"
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
	var modules []Module
	cfg.UnmarshalKey("physics.modules", moduleList)
	for _, name := range moduleList {
		module := CreateModule(name)
		modules = append(modules, module)
	}
	return &phyiscsEngine{modules}
}

type phyiscsEngine struct {
	modules []Module
}

func (e *phyiscsEngine) Init() {

}

func (e *phyiscsEngine) Add(i interface{}) {

}

func (e *phyiscsEngine) Advance(dt time.Duration) {

}
