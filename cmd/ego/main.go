package main

import (
	_ "ego/internal/renderer/glfw"
	"ego/pkg/game"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	err := readConfig()
	if err != nil {
		panic(err)
	}

	game := game.CreateGame("viper")
	game.Start()
}

func readConfig() error {
	var err error
	viper.SetConfigName("config")
	viper.AddConfigPath("./assets/config/")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	imports := viper.GetStringSlice("imports")
	err = importConfig(imports)
	if err != nil {
		return err
	}
	return nil
}

func importConfig(imports []string) error {
	for _, x := range imports {
		v := viper.New()
		v.SetConfigName(x)
		v.AddConfigPath("./assets/config/")
		err := v.ReadInConfig()
		if err != nil {
			return err
		}

		err = viper.MergeConfigMap(v.AllSettings())
		if err != nil {
			return err
		}

		imports := v.GetStringSlice("imports")
		err = importConfig(imports)
		if err != nil {
			return err
		}
	}
	return nil
}
