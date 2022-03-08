package terrain

import "testing"

func TestCreateResources(t *testing.T) {
	resources := CreateResources()
	if resources.HasResource(Health) {
		t.Errorf("New Resources should have no resource")
	}
	resources.AddResource(Health, 1)
	resources.AddResource(Food, 2)

	if !resources.HasResource(Health) {
		t.Error("Resources should have Health")
	}

	if !resources.HasResource(Food) {
		t.Error("Resources should have Food")
	}

	resources.Consume(Health)
	resources.Consume(Food)

	if resources.HasResource(Health) {
		t.Error("Resources should have no more foo")
	}

	if !resources.HasResource(Food) {
		t.Error("Resources should still have bar")
	}

	resources.Consume(Food)

	if resources.HasResource(Food) {
		t.Error("Resources should have no more bar")
	}

}
