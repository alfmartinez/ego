package motivator

import (
	"testing"
)

func TestNeedLevel(t *testing.T) {
	need := CreateNeed("foo", 1)
	level := CreateNeedLevel(need, 100)
	increment := CreateLevelIncrement(-1, 1)
	level.AddIncrement(increment)
	if level.Name() != need.Name() {
		t.Error("Need Level should have same name as need")
	}
	if level.Priority() != need.Priority() {
		t.Error("Need Level should have same priority as need")
	}
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
