package game

import (
	"testing"
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

	t.Run("uses file factory", func(t *testing.T) {
		game := CreateGame("file")
		if _, ok := game.(*sampleGame); !ok {
			t.Error("file factory should return sample game")
		}

	})
}
