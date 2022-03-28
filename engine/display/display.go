package display

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/context"
	"image"

	"github.com/spf13/viper"
)

type Display interface {
	Init()
	Render() image.Image
	AddObject(interface{})
}

type CropableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

var displayFactories = make(map[string]func() Display)

func RegisterDisplay(name string, f func() Display) {
	displayFactories[name] = f
}

func CreateDisplay() Display {
	viper := context.GetContext().Get("cfg").(*viper.Viper)
	name := viper.GetString("renderer.display.type")
	factory, ok := displayFactories[name]
	if !ok {
		panic(fmt.Errorf("can't find display factory for type %s", name))
	}
	return factory()
}
