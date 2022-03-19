package physics

import (
	"ego/pkg/configuration"
	"ego/pkg/context"
	"ego/pkg/physics/module"
	"fmt"
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
	err := cfg.UnmarshalKey("physics.modules", &moduleList)
	if err != nil {
		panic(err)
	}
	for idx, name := range moduleList {
		fmt.Printf("%d : %s\n", idx, name)
		module := module.CreateModule(name)
		modules = append(modules, module)
	}
	return &phyiscsEngine{modules}
}

type phyiscsEngine struct {
	modules []module.Module
}

func (e *phyiscsEngine) Init() {

}

func (e *phyiscsEngine) Add(i interface{}) {

}

func (e *phyiscsEngine) Advance(dt time.Duration) {

}
