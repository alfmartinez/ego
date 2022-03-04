package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Ego Client")
	w.SetPadded(false)
	w.SetContent(widget.NewLabel("Hello world !"))
	w.SetFullScreen(true)
	w.ShowAndRun()
}
