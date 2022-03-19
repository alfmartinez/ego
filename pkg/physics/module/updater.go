package module

import "time"

func init() {
	RegisterModuleFactory("updater", func() Module {
		return &updater{}
	})
}

type updater struct {
}

// Init implements Module
func (*updater) Init() {

}

// Add implements Module
func (*updater) Add(interface{}) {

}

// Advance implements Module
func (*updater) Advance(dt time.Duration, m []interface{}) []interface{} {
	return m
}
