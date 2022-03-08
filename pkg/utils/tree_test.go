package utils

import "testing"

func TestTree(t *testing.T) {
	t.Run("CreateTree returns tree pointer", func(t *testing.T) {
		o := CreateTree()
		if _, ok := o.(*tree); !ok {
			t.Errorf("should have tree pointer, got %+v", o)
		}
	})

	t.Run("Tree has no children initially", func(t *testing.T) {
		o := CreateTree()
		children := o.Children()
		if len(children) != 0 {
			t.Error("Should have no children")
		}
	})

	t.Run("Tree can add a single child", func(t *testing.T) {
		o := CreateTree()
		c := CreateTree()
		o.AddChild(c)
		children := o.Children()
		if children[0] != c {
			t.Errorf("should have returned children, got %+v", children)
		}
	})

	t.Run("Tree can add a multiple children", func(t *testing.T) {
		o := CreateTree()
		c := []Tree{CreateTree(), CreateTree()}
		o.AddChildren(c)
		children := o.Children()
		if children[0] != c[0] || children[1] != c[1] {
			t.Errorf("should have returned children, got %+v", children)
		}
	})
}
