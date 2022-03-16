package game

import (
	"testing"

	"github.com/spf13/viper"
)

type fakeGame struct{}

func (f *fakeGame) Start() {}
func (f *fakeGame) Stop()  {}

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
		viper.Set("renderer.type", "glfw")
		viper.Set("renderer.display.type", "rts")
		game := CreateGame("viper")
		if _, ok := game.(*sampleGame); !ok {
			t.Error("viper factory should return sample game")
		}

	})
}
