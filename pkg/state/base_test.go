package state

import "testing"

func TestBase(t *testing.T) {
	s := &baseState{}

	s.Update(nil, nil)
}
