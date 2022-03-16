package display

import (
	"ego/pkg/loader"
	"errors"
	"image"
	"image/color"
	"image/draw"

	_ "image/png"

	"github.com/spf13/viper"
)

func init() {
	RegisterDisplay("rts", func() Display {
		var displayConfig RtsData
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
}

func (d *rts) Init() {

	d.buffer = createBlankBuffer(d.size.X, d.size.Y)
	d.vpOrigin = image.Point{0, 0}
}

func (d *rts) Render() image.Image {
	buffer := d.buffer
	d.buffer = createBlankBuffer(d.size.X, d.size.Y)

	if p, ok := buffer.(CropableImage); ok {
		vpLimit := d.vpOrigin.Add(d.viewport)
		cropRect := image.Rectangle{d.vpOrigin, vpLimit}
		cropImg := p.SubImage(cropRect)
		buffer = cropImg.(draw.Image)
	}

	return buffer
}

func (d *rts) AddObject(s Displayable) {
	if s != nil {
		src := d.loader.GetSprite(s.Path(), s.Size())
		pos := s.Position()
		srcPoint := pos
		if src == nil {
			panic(errors.New("Sprite not found " + s.Path()))
		}
		r := image.Rectangle{srcPoint, srcPoint.Add(src.Bounds().Size())}
		draw.Draw(d.buffer, r, src, d.buffer.Bounds().Min, draw.Over)
	}
}

func createBlankBuffer(w, h int) draw.Image {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0}, draw.Src)
	return m
}
