package display

import (
	"ego/pkg/configuration"
	"ego/pkg/renderable"
	"ego/pkg/sprite/loader"
	"errors"
	"image"
	"image/color"
	"image/draw"

	_ "image/png"
)

func init() {
	RegisterDisplay("rts", func(config configuration.Display) Display {
		loader := loader.CreateSpriteLoader("on_demand")
		return &rts{
			loader: loader,
			config: config,
		}
	})
}

type rts struct {
	loader   loader.Loader
	buffer   draw.Image
	config   configuration.Display
	vpOrigin image.Point
}

func (d *rts) Init() {
	d.buffer = createBlankBuffer(d.config.Size.Width, d.config.Size.Height)
	d.vpOrigin = image.Point{0, 0}
}

func (d *rts) GetSize() configuration.Size {
	return d.config.ViewPort
}

func (d *rts) Render() image.Image {
	buffer := d.buffer
	d.buffer = createBlankBuffer(d.config.Size.Width, d.config.Size.Height)

	if p, ok := buffer.(CropableImage); ok {
		vpLimit := d.vpOrigin.Add(image.Point{d.config.ViewPort.Width, d.config.ViewPort.Height})
		cropRect := image.Rectangle{d.vpOrigin, vpLimit}
		cropImg := p.SubImage(cropRect)
		return cropImg
	}

	return buffer
}

func (d *rts) AddObject(s renderable.Renderable) {
	src := d.loader.GetSprite(s.Path(), s.Size())
	pos := s.Position()
	srcPoint := pos
	if src == nil {
		panic(errors.New("Sprite not found " + s.Path()))
	}
	r := image.Rectangle{srcPoint, srcPoint.Add(src.Bounds().Size())}
	draw.Draw(d.buffer, r, src, d.buffer.Bounds().Min, draw.Over)
}

func createBlankBuffer(w, h int) draw.Image {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0}, draw.Src)
	return m
}
