package main

import (
	"ego/pkg/game"
	"errors"
	"testing"
)

func fakeGameFactory(f func() (game.Game, error)) func() (game.Game, error) {
	return f
}

type fakeGame struct {
	launched bool
}

func (f *fakeGame) Start() {
	f.launched = true
}

func (f *fakeGame) Stop() {}

func TestMainCreatesGameAndLaunchesIt(t *testing.T) {
	mockGame := &fakeGame{false}
	gameFactory := fakeGameFactory(func() (game.Game, error) {
		return mockGame, nil
	})
	game.RegisterGameFactory("file", gameFactory)
	main()
	if !mockGame.launched {
		t.Error("main should have launched game")
	}
}

func TestMainPanicsIfGameCannotBeCreated(t *testing.T) {
	mockGame := &fakeGame{false}
	gameFactory := fakeGameFactory(func() (game.Game, error) {
		return mockGame, errors.New("Oupsie")
	})
	game.RegisterGameFactory("file", gameFactory)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	main()
	if mockGame.launched {
		t.Error("main should not launch game")
	}
}
