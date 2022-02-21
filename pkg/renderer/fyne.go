package renderer

import (
	"ego/pkg/display"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const (
	cellSize = 4
)

type FyneRenderer struct {
	display display.Display
	window  fyne.Window
}

func (r *FyneRenderer) IsAsync() bool {
	return true
}

func (r *FyneRenderer) Init() {
	app := app.New()
	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		panic(err)
	}
	app.SetIcon(icon)
	r.window = app.NewWindow("Ego")
}

func (r *FyneRenderer) Start(exit chan bool) {
	//r.window.SetFullScreen(true)
	r.window.Resize(fyne.NewSize(450, 450))
	r.window.SetOnClosed(func() {
		exit <- true
	})
	r.window.ShowAndRun()
}

func (r *FyneRenderer) Render(s Renderable) {

}

func (r *FyneRenderer) Refresh() {
	image := canvas.NewImageFromImage(r.display.Render())
	r.window.Canvas().SetContent(image)
	r.window.Content().Refresh()
}
