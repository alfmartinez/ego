package rts

import (
	"ego/pkg/context"
	"ego/pkg/display"
	"ego/pkg/layer"
	"ego/pkg/loader"
	"errors"
	"image"
	"image/color"
	"image/draw"

	_ "image/png"

	"github.com/spf13/viper"
)

func init() {
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
	buffer   draw.Image
	vpOrigin image.Point
	layers   layer.LayerMap
}

func (d *rts) Init() {
	d.buffer = createBlankBuffer(d.size.X, d.size.Y)
	d.vpOrigin = image.Point{0, 0}
	d.layers = layer.CreateLayerMap()
}

func (d *rts) Render() image.Image {
	buffer := d.renderLayers()

	if p, ok := buffer.(display.CropableImage); ok {
		vpLimit := d.vpOrigin.Add(d.viewport)
		cropRect := image.Rectangle{d.vpOrigin, vpLimit}
		cropImg := p.SubImage(cropRect)
		buffer = cropImg.(draw.Image)
	}

	return buffer
}

func (d *rts) renderLayers() image.Image {
	buffer := createBlankBuffer(d.size.X, d.size.Y)
	for _, elements := range d.layers {
		for _, element := range elements {
			s := element.(display.Displayable)
			src := d.loader.GetSprite(s.Path(), s.Size())
			pos := s.Position()
			srcPoint := pos
			if src == nil {
				panic(errors.New("Sprite not found " + s.Path()))
			}
			r := image.Rectangle{srcPoint, srcPoint.Add(src.Bounds().Size())}
			draw.Draw(buffer, r, src, buffer.Bounds().Min, draw.Over)
		}
	}
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
