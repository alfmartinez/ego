package game

import (
	"testing"
)

type fakeMob struct {
	Updated bool
}

func (f *fakeMob) Update() {
	f.Updated = true
}

func TestCreateSampleGame(t *testing.T) {
	t.Error("Unimplemented")
}
