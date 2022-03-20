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
	t.Run("Simple observer", func(t *testing.T) {
		var called bool
		evt := CreateEvent(UPDATE)
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
	})

	t.Run("Two observers", func(t *testing.T) {
		var called1, called2 int
		evt := CreateEvent(UPDATE)
		s := CreateSubject()
		o1 := &mockObserver{func(e Event) {
			called1++
			if e.Type() != evt.Type() {
				t.Errorf("Observer should be notified about %v, received %v", evt, e)
			}
		}}
		o2 := &mockObserver{func(e Event) {
			called2++
			if e.Type() != evt.Type() {
				t.Errorf("Observer should be notified about %v, received %v", evt, e)
			}
		}}
		s.Register(o1)
		s.Register(o2)
		s.NotifyAll(evt)
		s.Unregister(o2)
		s.NotifyAll(evt)
		if called1 != 2 {
			t.Errorf("Observer 1 should have been notified twice, got %d", called1)
		}
		if called2 != 1 {
			t.Errorf("Observer 2 should have been notified once, got %d", called2)
		}
	})

}
