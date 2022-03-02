package motivator

import "testing"

func TestNeed(t *testing.T) {
	need := CreateNeed("foo", 1)
	if need.Name() != "foo" {
		t.Errorf("need should have name 'foo', got %s", need.Name())
	}
	if need.Priority() != 1 {
		t.Errorf("need should have priority 1, got %d", need.Priority())
	}
}
