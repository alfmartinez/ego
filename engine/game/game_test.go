package game

import (
	"ego/engine/context"
	"ego/engine/display"
	"ego/engine/renderer"
	"image"
	"testing"

	"github.com/spf13/viper"
)

type fakeGame struct{}

func (f *fakeGame) Start() {}
func (f *fakeGame) Stop()  {}

type fakeDisplay struct{}

func (f *fakeDisplay) AddObject(any) {}
func (f *fakeDisplay) Init()         {}

func (f *fakeDisplay) Render() image.Image {
	return image.NewNRGBA(image.Rect(0, 0, 0, 0))
}

func TestCreateGame(t *testing.T) {
	RegisterGameFactory("test", func() Game {
		return &fakeGame{}
	})

	t.Run("uses Registered Factories", func(t *testing.T) {
		game := CreateGame("test")

		_, ok := game.(*fakeGame)
		if !ok {
			t.Errorf("CreateGame should have returned FakeGame")
		}
	})

	t.Run("uses viper factory", func(t *testing.T) {
		context.CreateAndRegisterContext("grrrr")
		renderer.RegisterRendererFactory("foo", func() renderer.Renderer {
			return &fakeRenderer{}
		})
		display.RegisterDisplay("bar", func() display.Display {
			return &fakeDisplay{}
		})
		v := viper.New()

		type mobData struct {
		}
		context.Set("cfg", v)
		v.Set("renderer.type", "foo")
		v.Set("renderer.display.type", "bar")
		v.Set("mobs", map[string]mobData{
			"foo": {},
			"bar": {},
		})
		game := CreateGame("viper")
		if _, ok := game.(*sampleGame); !ok {
			t.Error("viper factory should return sample game")
		}

	})
}
