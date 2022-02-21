package display

import (
	"ego/pkg/renderable"
	"image"
	"image/color"
	"image/draw"
)

type rts struct {
	buffer image.Image
}

func (d *rts) Init() {
	m := image.NewRGBA(image.Rect(0, 0, 450, 450))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0}, draw.Src)
	d.buffer = m
}

func (d *rts) Render() image.Image {
	return d.buffer
}

func (d *rts) AddObject(s renderable.Renderable) {

}
