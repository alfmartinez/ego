package module

import "time"

func init() {
	RegisterModuleFactory("collider", func() Module {
		return &collider{}
	})
}

type collider struct {
}

// Init implements Module
func (*collider) Init() {

}

// Add implements Module
func (*collider) Add(interface{}) {

}

// Advance implements Module
func (*collider) Advance(dt time.Duration, m []interface{}) []interface{} {
	return m
}
