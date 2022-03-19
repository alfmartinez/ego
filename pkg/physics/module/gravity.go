package module

import "time"

func init() {
	RegisterModuleFactory("gravity", func() Module {
		return &gravity{}
	})
}

type gravity struct{}

// Init implements Module
func (*gravity) Init() {

}

// Add implements Module
func (*gravity) Add(interface{}) {

}

// Advance implements Module
func (*gravity) Advance(dt time.Duration, m []interface{}) []interface{} {
	return m
}
