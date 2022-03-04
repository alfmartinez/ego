package main

import (
	"ego/pkg/game"
	"testing"
)

func fakeGameFactory(f func() game.Game) func() game.Game {
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
	gameFactory := fakeGameFactory(func() game.Game {
		return mockGame
	})
	game.RegisterGameFactory("file", gameFactory)
	main()
	if !mockGame.launched {
		t.Error("main should have launched game")
	}
}
