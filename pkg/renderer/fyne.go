package renderer

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type FyneRenderer struct {
	window   fyne.Window
	imgCache *image.RGBA
	buffer   [101][101]color.Color
	flush    int
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

	raster := canvas.NewRaster(r.draw)
	r.window.Canvas().SetContent(raster)
}

func (r *FyneRenderer) Start() {
	//r.window.SetFullScreen(true)
	r.window.Resize(fyne.NewSize(450, 450))
	r.window.ShowAndRun()
}

func (r *FyneRenderer) Render(s Renderable) {
	colors := map[string]color.Color{
		"Charlie": color.RGBA{0x00, 0x00, 0xff, 0xff},
		"Arthur":  color.RGBA{0xff, 0x00, 0x00, 0xff},
		"Bob":     color.RGBA{0x00, 0xff, 0x00, 0xff},
	}

	r.buffer[s.Position().Y][s.Position().X] = colors[s.Name()]
}

func (r *FyneRenderer) Refresh() {
	r.window.Content().Refresh()
}

const (
	cellSize = 4
)

func (r *FyneRenderer) cellForCoord(x, y int, density float32) (int, int) {
	return int(float32(x) / float32(cellSize) / density),
		int(float32(y) / float32(cellSize) / density)
}

func (r *FyneRenderer) pixelDensity() float32 {
	c := fyne.CurrentApp().Driver().CanvasForObject(r.window.Content())
	if c == nil {
		return 1.0
	}

	pixW, _ := c.PixelCoordinateForPosition(fyne.NewPos(cellSize, cellSize))
	return float32(pixW) / float32(cellSize)
}

func (r *FyneRenderer) draw(x int, y int) image.Image {
	density := r.pixelDensity()
	pixW, pixH := r.cellForCoord(x, y, density)
	img := r.imgCache
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, pixW, pixH))
	}

	for y := 0; y < pixH; y++ {
		for x := 0; x < pixW; x++ {
			if x < 101 && y < 101 && r.buffer[y][x] != nil {
				pixColor := r.buffer[y][x]
				img.Set(x+1, y+1, pixColor)
			} else {
				img.Set(x+1, y+1, color.Black)
			}
		}
	}

	return img

}
