package motivator

type LevelIncrement interface {
	Apply() int
	IsOver() bool
}

type levelIncrement struct {
	increment int
	duration  int
}

func CreateLevelIncrement(i int, d int) LevelIncrement {
	return &levelIncrement{i, d}
}

func (inc *levelIncrement) Apply() int {
	if inc.duration > 0 {
		inc.duration = inc.duration - 1
	}
	return inc.increment
}

func (inc *levelIncrement) IsOver() bool {
	return inc.duration == 0
}
