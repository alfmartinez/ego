package movement

import (
	"image"
	"testing"
)

func TestLocation(t *testing.T) {
	location := Loc(image.Pt(1, 1))
	sameLocation := Loc(image.Pt(1, 1))
	otherLocation := Loc(image.Pt(0, 1))
	if !location.IsAt(location) {
		t.Error("Location should be at self location")
	}
	if !location.IsAt(sameLocation) {
		t.Error("Location should be at same location")
	}
	if location.IsAt(otherLocation) {
		t.Error("Location should not be at other location")
	}
}
