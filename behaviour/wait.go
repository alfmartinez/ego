package behaviour

type wait struct {
	behaviour
}

func (behaviour *wait) GetName() string {
	return behaviour.Name
}

func (behaviour *wait) Evaluate() {
	if behaviour.isTimerRunning() {
		behaviour.decrementTimer()
	} else {
		behaviour.startTimer(1)
	}
}

func (behaviour *wait) IsOver() bool {
	return behaviour.isTimerOver()
}

func (behaviour *wait) Reset() {
	behaviour.resetTimer()
}

func (behaviour *wait) Next() string {
	return "sleep"
}
