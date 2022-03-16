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
	viper.SetConfigName("config")
	viper.AddConfigPath("./assets/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	imports := viper.GetStringSlice("imports")
	for _, x := range imports {
		v := viper.New()
		v.SetConfigName(x)
		v.AddConfigPath("./assets/config/")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		err = viper.MergeConfigMap(v.AllSettings())
		if err != nil {
			panic(fmt.Errorf("cannot merge config %s", x))
		}
	}

	game := game.CreateGame("viper")
	game.Start()
}
