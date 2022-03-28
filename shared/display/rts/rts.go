package rts

import (
	"fmt"
	"github.com/alfmartinez/ego/engine/context"
	"github.com/alfmartinez/ego/engine/display"
	"github.com/alfmartinez/ego/engine/layer"
	"github.com/alfmartinez/ego/engine/loader"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"sort"

	"github.com/spf13/viper"
)

func Register() {
	display.RegisterDisplay("rts", func() display.Display {
		var displayConfig RtsData
		viper := context.GetContext().Get("cfg").(*viper.Viper)
		err := viper.UnmarshalKey("renderer.display", &displayConfig)
		if err != nil {
			panic(err)
		}
		loader := loader.CreateSpriteLoader("pre_load")
		loader.Init()
		return &rts{
			loader:   loader,
			size:     displayConfig.Size,
			viewport: displayConfig.ViewPort,
		}
	})
}

type RtsData struct {
	Size     image.Point
	ViewPort image.Point
}

type rts struct {
	loader   loader.Loader
	size     image.Point
	viewport image.Point
	vpOrigin image.Point
	layers   layer.LayerMap
}

func (d *rts) Init() {
	d.vpOrigin = image.Point{0, 0}
	d.layers = layer.CreateLayerMap()
}

func (d *rts) Render() image.Image {
	buffer := d.renderLayers()

	//	if p, ok := buffer.(display.CropableImage); ok {
	//		vpLimit := d.vpOrigin.Add(d.viewport)
	//		cropRect := image.Rectangle{d.vpOrigin, vpLimit}
	//		cropImg := p.SubImage(cropRect)
	//		buffer = cropImg.(draw.Image)
	//	}

	return buffer
}

func (d *rts) renderLayers() image.Image {
	layerKeys := make([]layer.Layer, 0)
	for k, _ := range d.layers {
		layerKeys = append(layerKeys, k)
	}

	sort.Slice(layerKeys, func(i, j int) bool {
		return layerKeys[i] < layerKeys[j]
	})

	buffer := createBlankBuffer(d.size.X, d.size.Y)
	for idx := range layerKeys {
		elements := d.layers[layerKeys[idx]]
		for _, element := range elements {
			s := element.(display.Displayable)
			src := d.loader.GetSprite(s.Path(), s.Size())
			pos := s.Position()
			srcPoint := pos
			if src == nil {
				panic(fmt.Errorf("sprite not found %s", s.Path()))
			}
			r := image.Rectangle{srcPoint, srcPoint.Add(src.Bounds().Size())}
			draw.Draw(buffer, r, src, buffer.Bounds().Min, draw.Over)
		}
	}

	d.layers = layer.CreateLayerMap()
	return buffer
}

func (d *rts) AddObject(i interface{}) {
	s := i.(display.Displayable)
	d.layers.Add(s.Layer(), s)
}

func createBlankBuffer(w, h int) draw.Image {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0}, draw.Src)
	return m
}
