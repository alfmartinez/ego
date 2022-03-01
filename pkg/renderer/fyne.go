package renderer

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/renderable"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func init() {
	RegisterRendererFactory("fyne", func(config configuration.Renderer) Renderer {
		display := display.CreateDisplay(config.Display)
		return &fyneRenderer{display: display}
	})
}

const (
	cellSize = 4
)

type fyneRenderer struct {
	display display.Display
	window  fyne.Window
}

func (r *fyneRenderer) IsAsync() bool {
	return true
}

func (r *fyneRenderer) Init() {
	app := app.New()
	icon, err := fyne.LoadResourceFromPath("Icon.png")
	if err != nil {
		panic(err)
	}
	app.SetIcon(icon)
	r.window = app.NewWindow("Ego")
	r.window.SetPadded(false)
	r.display.Init()
}

func (r *fyneRenderer) Start(exit chan bool) {
	//r.window.SetFullScreen(true)
	size := r.display.GetSize()
	r.window.Resize(fyne.NewSize(float32(size.Width), float32(size.Height)))
	r.window.SetOnClosed(func() {
		exit <- true
	})
	r.window.ShowAndRun()
}

func (r *fyneRenderer) Render(s renderable.Renderable) {
	r.display.AddObject(s)
}

func (r *fyneRenderer) Refresh() {
	image := canvas.NewImageFromImage(r.display.Render())
	r.window.Canvas().SetContent(image)
	r.window.Content().Refresh()
}
