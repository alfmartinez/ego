package observer

import "testing"

type NotifyMethod func(Event)
type mockObserver struct {
	notify NotifyMethod
}

func (o *mockObserver) OnNotify(e Event) {
	o.notify(e)
}

func TestObserver(t *testing.T) {
	var called bool
	evt := CreateEvent(UPDATE, nil)
	s := CreateSubject()
	s.Register(&mockObserver{func(e Event) {
		called = true
		if e.Type() != evt.Type() {
			t.Errorf("Observer should be notified about %v, received %v", evt, e)
		}
	}})
	s.NotifyAll(evt)
	if !called {
		t.Error("Observer should have been notified")
	}
}
