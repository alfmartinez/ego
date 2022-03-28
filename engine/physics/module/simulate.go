package module

import (
	"github.com/alfmartinez/ego/engine/physics/mobile"
	"github.com/alfmartinez/ego/engine/slices"
)

func init() {
	RegisterModuleFactory("simulate", func() Module {
		return &simulate{}
	})
}

type simulate struct {
	mobiles []mobile.Mobile
}

// Init implements Module
func (m *simulate) Init() {
	m.mobiles = make([]mobile.Mobile, 0)
}

// Add implements Module
func (m *simulate) Add(i any) {
	if o, ok := i.(mobile.Mobile); ok {
		m.mobiles = append(m.mobiles, o)
	}
}

// Advance implements Module
func (m *simulate) Advance(dt float64) {
	slices.Apply(m.mobiles, func(m mobile.Mobile) {
		matrix := m.GetMatrix()
		s := matrix.S.Add(matrix.A.Mul(dt))
		p := matrix.P.Add(matrix.S.Mul(dt))
		matrix.S = s
		matrix.P = p
		m.SetMatrix(matrix)
	})
}
