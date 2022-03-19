package motivator

import (
	"testing"
)

func TestNeedLevel(t *testing.T) {
	need := Health
	level := CreateNeedLevel(need, 100)
	increment := CreateLevelIncrement(-1, 1)
	level.AddIncrement(increment)

	if level.Value() != 100 {
		t.Errorf("Need level should have level %d", level.Value())
	}
	level.Update()
	if level.Value() != 99 {
		t.Errorf("Should have decreased 1 pt, got %d", level.Value())
	}
	level.Update()
	if level.Value() != 99 {
		t.Errorf("Should not decrease more")
	}

	increment = CreateLevelIncrement(1, 3)
	level.AddIncrement(increment)
	for !increment.IsOver() {
		level.Update()
	}

	if level.Value() != 100 {
		t.Error("Error")
	}

	increment = CreateLevelIncrement(-1, -1)
	var count = 0
	level.AddIncrement(increment)
	for level.Value() > 0 {
		level.Update()
		count++
	}
	if count > 100 {
		t.Error("Error")
	}
	level.Update()
	if level.Value() != 0 {
		t.Error("Level should not fall below 0")
	}
}
