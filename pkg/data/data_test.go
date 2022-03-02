package data

import "testing"

func TestData(t *testing.T) {
	data := CreateData("foo")
	if data.Name() != "foo" {
		t.Errorf("Expected 'foo', got %s", data.Name())
	}
}
