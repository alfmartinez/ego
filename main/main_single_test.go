package main

import (
	"testing"
)

func TestLaunchSingle(t *testing.T) {
	game, error := Launch("single.yml")
	if game == nil {
		t.Errorf("Launch should return true if configuration file parsed : %s", error.Error())
	}
	if error != nil {
		t.Errorf("Launch should return no error code for missing, got %s", error.Error())
	}
	game.Init()
	if game.Mobs[0].Status != "idle" {
		t.Errorf("Character should be idle after init, was %s", game.Mobs[0].Status)
	}
	game.Tick()
	if game.Mobs[0].Status != "wait" {
		t.Errorf("Character should wait after first tick, was %s", game.Mobs[0].Status)
	}
	game.Tick()
	if game.Mobs[0].Status != "wait" {
		t.Errorf("Character should wait after second tick, was %s", game.Mobs[0].Status)
	}
	game.Tick()
	if game.Mobs[0].Status != "idle" {
		t.Errorf("Character should wait after third tick, was %s", game.Mobs[0].Status)
	}
	game.Tick()
	if game.Mobs[0].Status != "wait" {
		t.Errorf("Character should wait after fourth tick, was %s", game.Mobs[0].Status)
	}
}
