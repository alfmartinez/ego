package module

import (
	"github.com/alfmartinez/ego/engine/physics/mobile"
	"github.com/alfmartinez/ego/engine/slices"
)

func init() {
	RegisterModuleFactory("gravity", func() Module {
		return &gravity{}
	})
}

type gravity struct {
	mobiles []mobile.Mobile
}

// Init implements Module
func (m *gravity) Init() {
	m.mobiles = make([]mobile.Mobile, 0)
}

// Add implements Module
func (m *gravity) Add(i any) {
	if o, ok := i.(mobile.Mobile); ok {
		m.mobiles = append(m.mobiles, o)
	}
}

// Advance implements Module
func (m *gravity) Advance(dt float64) {
	slices.Apply(m.mobiles, func(m mobile.Mobile) {
		matrix := m.GetMatrix()
		matrix.A.Y = 10
		m.SetMatrix(matrix)
	})
}
