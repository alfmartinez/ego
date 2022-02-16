package behaviour

type sleep struct {
	behaviour
}

func (behaviour *sleep) GetName() string {
	return behaviour.Name
}

func (behaviour *sleep) Evaluate() {
	if behaviour.isTimerRunning() {
		behaviour.decrementTimer()
	} else {
		behaviour.startTimer(10)
	}
}

func (behaviour *sleep) IsOver() bool {
	return behaviour.isTimerOver()
}

func (behaviour *sleep) Reset() {
	behaviour.resetTimer()
}

func (behaviour *sleep) Next() string {
	return "idle"
}
