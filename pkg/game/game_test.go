package game

import "testing"

type fakeGame struct{}

func (f *fakeGame) Start() {}
func (f *fakeGame) Stop()  {}

func TestCreateGameUsesRegisterFactory(t *testing.T) {
	RegisterGameFactory("test", func() Game {
		return &fakeGame{}
	})
	game := CreateGame("test")

	_, ok := game.(*fakeGame)
	if !ok {
		t.Errorf("CreateGame should have returned FakeGame")
	}
}
