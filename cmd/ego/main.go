package main

import (
	_ "ego/internal/renderer/glfw"
	"ego/pkg/game"
	"fmt"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	viper.SetConfigName("game")
	viper.AddConfigPath("./assets/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	game := game.CreateGame("viper")
	game.Start()
}
