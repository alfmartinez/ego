package game

import "testing"

type fakeGame struct{}

func (f *fakeGame) Start() {}

func TestCreateGameUsesRegisterFactory(t *testing.T) {
	RegisterGameFactory("test", func() (Game, error) {
		return &fakeGame{}, nil
	})
	game, err := CreateGame("test")
	if err != nil {
		t.Errorf("CreateGame should not return error : %+v", err)
	}
	_, ok := game.(*fakeGame)
	if !ok {
		t.Errorf("CreateGame should have returned FakeGame")
	}
}
