package game

import "testing"

func TestScene(t *testing.T) {
	t.Run("CreateScene gives Scene with Root Node", func(t *testing.T) {
		s := CreateScene()
		root := s.Root()
		if _, ok := s.(*scene); !ok {
			t.Error("should return scene object")
		}
		if _, ok := root.(*sceneNode); !ok {
			t.Error("should return root scene node")
		}
	})

	t.Run("Update should update root node", func(t *testing.T) {
		var called bool = false
		s := CreateScene()
		s.Root().SetValue(&mockObject{})
		mockUpdateMethod = func() {
			called = true
		}
		s.Update()
		if !called {
			t.Error("Object Update method should have been called")
		}
	})
	t.Run("Render returns render tree", func(t *testing.T) {
		s := CreateScene()
		s.Root().SetValue(&mockObject{})
		tree := s.Render()
		children := tree.Root().Children()
		if len(children) != 1 {
			t.Errorf("%+v", children)
		}
	})

}
