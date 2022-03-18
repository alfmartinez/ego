package renderer

import (
	"testing"
)

type fakeRenderer struct{}

func (f *fakeRenderer) Init()              {}
func (f *fakeRenderer) IsAsync() bool      { return false }
func (f *fakeRenderer) Start()             {}
func (f *fakeRenderer) Render(interface{}) {}
func (f *fakeRenderer) Refresh()           {}
func (f *fakeRenderer) Close()             {}

func TestCreateRenderer(t *testing.T) {
	RegisterRendererFactory("fake", func() Renderer {
		return &fakeRenderer{}
	})
	t.Run("should create registered renderer", func(t *testing.T) {
		o := CreateRenderer("fake")
		if _, ok := o.(*fakeRenderer); !ok {
			t.Errorf("Should return fakeRenderer, got %+v", o)
		}
	})

	t.Run("should panic if not registered", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Should have panicked")
			}
		}()
		CreateRenderer("foo")

	})
}
