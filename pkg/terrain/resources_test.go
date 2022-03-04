package terrain

import "testing"

func TestCreateResources(t *testing.T) {
	resources := CreateResources()
	if resources.HasResource("foo") {
		t.Errorf("New Resources should have no resource")
	}
	resources.AddResource("foo", 1)
	resources.AddResource("bar", 2)

	if !resources.HasResource("foo") {
		t.Error("Resources should have foo")
	}

	if !resources.HasResource("bar") {
		t.Error("Resources should have bar")
	}

	resources.Consume("foo")
	resources.Consume("bar")

	if resources.HasResource("foo") {
		t.Error("Resources should have no more foo")
	}

	if !resources.HasResource("bar") {
		t.Error("Resources should still have bar")
	}

	resources.Consume("bar")

	if resources.HasResource("bar") {
		t.Error("Resources should have no more bar")
	}

}
