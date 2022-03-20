package module

import (
	"ego/engine/physics/matrix"
	"ego/engine/physics/mobile"
	"image"
	"testing"
)

type mockMobile struct {
	mockGetMatrix func() matrix.M
	mockSetMatrix func(matrix.M)
	mockIsHit     func(image.Rectangle) bool
}

func (m *mockMobile) GetMatrix() matrix.M          { return m.mockGetMatrix() }
func (m *mockMobile) SetMatrix(mat matrix.M)       { m.mockSetMatrix(mat) }
func (m *mockMobile) IsHit(r image.Rectangle) bool { return m.mockIsHit(r) }

func Test_gravity_Init(t *testing.T) {
	type fields struct {
		mobiles []mobile.Mobile
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"blank",
			fields{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat := matrix.M{}
			m := &gravity{
				mobiles: tt.fields.mobiles,
			}
			m.Init()
			m.Add(&mockMobile{
				func() matrix.M { return mat },
				func(m matrix.M) { mat = m },
				func(r image.Rectangle) bool { return false },
			})
			m.Advance(1)
			if mat.A.Y != 10 {
				t.Errorf("Gravity should apply")
			}
		})
	}
}
