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
	d.buffer = createBlankBuffer(450, 450)
}

func (d *rts) Render() image.Image {
	buffer := d.buffer
	d.buffer = createBlankBuffer(450, 450)
	return buffer
}

func (d *rts) AddObject(s renderable.Renderable) {

}

func createBlankBuffer(w, h int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{0, 0}, draw.Src)
	return m
}
