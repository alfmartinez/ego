package display

import (
	"image"
	"testing"
)

func TestDisplayable(t *testing.T) {
	t.Run("CreateDisplayable returns a displayable struct", func(t *testing.T) {
		o := CreateDisplayable("foo", 1, image.Pt(0, 0))
		if _, ok := o.(displayable); !ok {
			t.Errorf("should have displayable, got %+v", o)
		}
	})

	t.Run("displayable has path", func(t *testing.T) {
		o := CreateDisplayable("foo", 1, image.Pt(0, 0))
		if o.Path() != "foo" {
			t.Errorf("Wrong path, got %+v", o.Path())
		}
	})

	t.Run("displayable has size", func(t *testing.T) {
		o := CreateDisplayable("foo", 1, image.Pt(0, 0))
		if o.Size() != 1 {
			t.Errorf("Wrong size, got %+v", o.Size())
		}
	})
	t.Run("displayable has position", func(t *testing.T) {
		o := CreateDisplayable("foo", 1, image.Pt(0, 0))
		if o.Position() != image.Pt(0, 0) {
			t.Errorf("Wrong path, got %+v", o.Position())
		}
	})
}
