package game

import (
	"testing"
)

func TestBackendGame(t *testing.T) {
	game := CreateGame("backend")
	t.Run("Game is backend", func(t *testing.T) {
		_, ok := game.(*backendGame)
		if !ok {
			t.Error("Should have game of backend")
		}
	})

}
