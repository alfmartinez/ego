package renderer

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type FyneRenderer struct {
	app    fyne.App
	canvas fyne.Canvas
	window fyne.Window
}

func (r *FyneRenderer) Init() {
	r.app = app.New()
	r.window = r.app.NewWindow("Ego")
	r.canvas = r.window.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	r.canvas.SetContent(rect)

}

func (r *FyneRenderer) Start() {
	r.window.Resize(fyne.NewSize(400, 400))
	r.window.ShowAndRun()
}

func (r *FyneRenderer) Render(s Renderable) {

}
