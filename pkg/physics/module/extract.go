package module

import (
	"ego/pkg/physics/mobile"
	"time"
)

func init() {
	RegisterModuleFactory("extract", func() Module {
		return &extract{}
	})
}

// Extract module extract PSA matrix from mobile
type extract struct {
	mobiles mobile.Mobiles
}

// Init implements Module
func (e *extract) Init() {
	e.mobiles = make(mobile.Mobiles, 0)
}

// Add implements Module
func (e *extract) Add(i interface{}) {
	if m, ok := i.(mobile.Mobile); ok {
		e.mobiles = append(e.mobiles, m)
	}
}

// Advance implements Module
func (e *extract) Advance(time.Duration, []interface{}) []interface{} {
	return e.mobiles.Map(func(m mobile.Mobile) interface{} {
		return m.GetMatrix()
	})
}
