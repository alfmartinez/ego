package fyne

import (
	"ego/pkg/configuration"
	"ego/pkg/display"
	"ego/pkg/render"
	"ego/pkg/renderer"
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

//go:embed data/Icon.png
var icon []byte

func init() {
	renderer.RegisterRendererFactory("fyne", func(config configuration.Renderer) renderer.Renderer {
		display := display.CreateDisplay(config.Display)
		return &fyneRenderer{display: display}
	})
}

type fyneRenderer struct {
	display display.Display
	window  fyne.Window
	exit    chan bool
}

func (r *fyneRenderer) IsAsync() bool {
	return true
}

func (r *fyneRenderer) Init() {
	app := app.New()
	var iconRes = &fyne.StaticResource{
		StaticName:    "Icon.png",
		StaticContent: icon,
	}
	app.SetIcon(iconRes)
	r.window = app.NewWindow("Ego")
	r.window.SetPadded(false)
	r.display.Init()
}

func (r *fyneRenderer) Start(exit chan bool) {
	//r.window.SetFullScreen(true)
	r.exit = exit
	size := r.display.GetSize()
	r.window.Resize(fyne.NewSize(float32(size.Width), float32(size.Height)))
	r.window.SetOnClosed(r.Close)
	r.window.ShowAndRun()
}

func (r *fyneRenderer) Close() {
	r.exit <- true
	close(r.exit)
}

func (r *fyneRenderer) Render(tree render.RenderTree) {
	tree.Apply(func(node render.RenderNode) {
		s := node.Display()
		r.display.AddObject(s)
	})

}

func (r *fyneRenderer) Refresh() {
	image := canvas.NewImageFromImage(r.display.Render())
	r.window.Canvas().SetContent(image)
	r.window.Content().Refresh()
}