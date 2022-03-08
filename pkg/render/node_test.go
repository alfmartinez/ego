package render

import (
	"ego/pkg/display"
	"image"
	"testing"
)

func TestRenderNode(t *testing.T) {
	t.Run("CreateRenderNode", func(t *testing.T) {
		o := CreateRenderNode()
		if _, ok := o.(*renderNode); !ok {
			t.Errorf("should return renderNode pointer, got %+v", o)
		}
		if o.Display() != nil {
			t.Errorf("render Node should not be initialized with display, got %+v", o.Display())
		}
	})

	t.Run("RenderNode should Apply function to children", func(t *testing.T) {
		d := display.CreateDisplayable("foo", 1, image.Pt(0, 0))

		o := CreateRenderNode()
		f := o.CreateFolder("foo")
		g := f.CreateFolder("bar")
		c := f.CreateObject(d)
		var nodes []RenderNode = make([]RenderNode, 0)
		o.Apply(func(rn RenderNode) {
			nodes = append(nodes, rn)
		})
		if nodes[0] != o {
			t.Errorf("function should have been called on Root RenderNode")
		}

		if nodes[1] != f {
			t.Errorf("function should have been called on Child RenderNode")
		}

		if nodes[2] != g {
			t.Errorf("function should have been called on Child RenderNode")
		}

		if nodes[3] != c {
			t.Errorf("function should have been called on Child RenderNode")
		}
	})
}
