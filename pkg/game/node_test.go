package game

import "testing"

var mockUpdateMethod func()

type mockObject struct{}

func (o mockObject) Update() {
	mockUpdateMethod()
}

func TestNode(t *testing.T) {
	t.Run("CreateSceneNode", func(t *testing.T) {
		n := CreateSceneNode()
		if n.Value() != nil {
			t.Error("New Node should have no value set")
		}
	})

	t.Run("Update updates node value", func(t *testing.T) {
		var called bool = false
		n := CreateSceneNode()
		o := mockObject{}
		mockUpdateMethod = func() {
			called = true
		}
		n.SetValue(o)
		n.Update()
		if !called {
			t.Error("Object should have been updated")
		}
	})

	t.Run("Apply is applied to node and children", func(t *testing.T) {
		var count int = 0
		n := CreateSceneNode()
		mockUpdateMethod = func() {
			count++
		}
		o0, o1, o2 := mockObject{}, mockObject{}, mockObject{}
		n.SetValue(o0)
		n.AddObject(o1)
		n.AddObject(o2)
		n.Apply(func(sn SceneNode) {
			sn.Update()
		})
		if count != 3 {
			t.Error("Update should have been called on all 3 nodes")
		}
	})
}
