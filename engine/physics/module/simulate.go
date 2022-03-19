package module

import "time"

func init() {
	RegisterModuleFactory("simulate", func() Module {
		return &simulate{}
	})
}

type simulate struct {
}

// Init implements Module
func (*simulate) Init() {

}

// Add implements Module
func (*simulate) Add(interface{}) {

}

// Advance implements Module
func (*simulate) Advance(dt time.Duration) {

}
