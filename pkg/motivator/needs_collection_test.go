package motivator

import "testing"

func TestNeedsCollection(t *testing.T) {
	collection := CreateNeedsCollection()
	collection.AddNeed(Health, 100)
	collection.AddNeed(Rest, 100)
	collection.AddNeed(Beauty, 100)
	collection.Provide(Health, -1, -1)
	collection.Provide(Rest, -1, 60)
	collection.Provide(Beauty, -2, 50)

	top := collection.TopNeed()
	if top != None {
		t.Errorf("Should have need foo, got %v", top)
	}

	var countNone = 0
	for collection.TopNeed() == None {
		collection.UpdateNeeds()
		countNone++
	}

	if countNone != 26 {
		t.Errorf("Count %d", countNone)
	}

	var countBaz = 0
	for collection.TopNeed() == Beauty {
		collection.UpdateNeeds()
		countBaz++
	}
	if countBaz != 25 {
		t.Errorf("Count %d", countBaz)
	}

	var countBar = 0
	collection.Provide(Health, 2, 10)
	for collection.TopNeed() == Rest {
		collection.UpdateNeeds()
		countBar++
	}
	if countBar != 20 {
		t.Errorf("Count should be 20, got %d", countBar)
	}

}
