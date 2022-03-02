package motivator

import "testing"

func TestNeedsCollection(t *testing.T) {
	collection := CreateNeedsCollection()
	need := CreateNeed("foo", 0)
	otherNeed := CreateNeed("bar", 1)
	anotherNeed := CreateNeed("baz", 0)
	collection.AddNeed(need, 100)
	collection.AddNeed(otherNeed, 100)
	collection.AddNeed(anotherNeed, 100)
	collection.Provide(need, -1, -1)
	collection.Provide(otherNeed, -1, 60)
	collection.Provide(anotherNeed, -2, 50)

	top := collection.TopNeed()
	if top != "none" {
		t.Errorf("Should have need foo, got %s", top)
	}

	var countNone = 0
	for collection.TopNeed() == "none" {
		collection.UpdateNeeds()
		countNone++
	}

	if countNone != 26 {
		t.Errorf("Count %d", countNone)
	}

	var countBaz = 0
	for collection.TopNeed() == "baz" {
		collection.UpdateNeeds()
		countBaz++
	}
	if countBaz != 25 {
		t.Errorf("Count %d", countBaz)
	}

	var countBar = 0
	collection.Provide(need, 2, 10)
	for collection.TopNeed() == "bar" {
		collection.UpdateNeeds()
		countBar++
	}
	if countBar != 20 {
		t.Errorf("Count should be 20, got %d", countBar)
	}

}
