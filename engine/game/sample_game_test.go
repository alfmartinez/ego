package game

import (
	"ego/engine/input"
	"ego/engine/observer"
	"testing"
	"time"
)

var mockStartMethod func()
var mockInitMethod func()
var mockCloseMethod func()
var mockRefreshMethod func()
var mockRenderMethod func(any)

type mockRenderer struct{}

func (r *mockRenderer) Start()       { mockStartMethod() }
func (r *mockRenderer) Init()        { mockInitMethod() }
func (r *mockRenderer) Close()       { mockCloseMethod() }
func (r *mockRenderer) Refresh()     { mockRefreshMethod() }
func (r *mockRenderer) Render(a any) { mockRenderMethod(a) }

func TestSampleGame(t *testing.T) {
	r := &mockRenderer{}
	t.Run("Start", func(t *testing.T) {
		var initCalled, startCalled, renderCalled, refreshCalled, closeCalled bool
		s := observer.CreateSubject()
		g := CreateSampleGame(s, r)
		mockInitMethod = func() {
			initCalled = true
		}
		mockStartMethod = func() {
			startCalled = true
		}
		mockRenderMethod = func(any) {
			renderCalled = true
		}
		mockRefreshMethod = func() {
			refreshCalled = true
		}
		mockCloseMethod = func() {
			closeCalled = true
		}

		go func() {
			time.Sleep(time.Second / 30)
			input.FromContext().Handle(input.Event{
				Action: input.PRESSED,
				Key:    input.ESCAPE,
			})
		}()
		g.Start()
		if !initCalled {
			t.Error("Init method should have been called")
		}
		if !startCalled {
			t.Error("Start method should have been called")
		}
		if renderCalled {
			t.Error("Render method should not have been called")
		}
		if !refreshCalled {
			t.Error("Refresh method should have been called")
		}
		if !closeCalled {
			t.Error("Close method should have been called")
		}
	})

}
