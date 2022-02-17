package renderer

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type FyneRenderer struct {
	canvas fyne.Canvas
	window fyne.Window
}

func (r *FyneRenderer) Init() {
	myApp := app.New()
	r.window = myApp.NewWindow("Ego")
	r.canvas = r.window.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	r.canvas.SetContent(rect)
	log.Print("Created App and Canvas")
}

func (r *FyneRenderer) Start() {
	log.Print("Launch App Loop")
	r.window.Resize(fyne.NewSize(400, 400))
	r.window.ShowAndRun()
}

func (r *FyneRenderer) Render(s Renderable) {

}
