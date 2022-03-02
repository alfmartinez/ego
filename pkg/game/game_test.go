package game

import "testing"

type fakeGame struct{}

func (f *fakeGame) Start() {}
func (f *fakeGame) Stop()  {}

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

func TestCreateGameFromFile(t *testing.T) {
	_, err := CreateGame("file")
	if err != nil {
		t.Errorf("Create game should return game without error, got %+v", err)
	}
}
