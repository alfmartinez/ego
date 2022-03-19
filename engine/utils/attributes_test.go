package utils

import "testing"

func TestAttributes(t *testing.T) {
	t.Run("CreateAttributes should return attributes", func(t *testing.T) {
		b := CreateAttributes()
		if _, ok := b.(*attributes); !ok {
			t.Errorf("Wrong type returned, got %+v", b)
		}
	})
	t.Run("Attributes has empty values initially", func(t *testing.T) {
		b := CreateAttributes()
		if v := b.Attribute("name"); v != "" {
			t.Errorf("Value retrieved wrong, got %+v", v)
		}
	})
	t.Run("Attributes can store and retrieve string values", func(t *testing.T) {
		b := CreateAttributes()
		b.SetAttribute("name", "foo")
		if v := b.Attribute("name"); v != "foo" {
			t.Errorf("Value retrieved wrong, got %+v", v)
		}
	})
}
