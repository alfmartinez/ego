package render

import "testing"

func TestTree(t *testing.T) {
	t.Run("CreateRenderTree", func(t *testing.T) {
		o := CreateRenderTree()
		if _, ok := o.(renderTree); !ok {
			t.Error("should return renderTree struct")
		}
	})

	t.Run("RenderTree has RenderNode Root", func(t *testing.T) {
		o := CreateRenderTree()
		root := o.Root()
		if _, ok := root.(*renderNode); !ok {
			t.Error("should have renderNode Root")
		}
	})

	t.Run("RenderTree can apply func to root", func(t *testing.T) {
		o := CreateRenderTree()
		var actual RenderNode = nil
		o.Apply(func(rn RenderNode) {
			actual = rn
		})
		if actual == nil {
			t.Error("Apply should have executed on root")
		}
	})
}
