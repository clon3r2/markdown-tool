package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	cfg := config{}
	edit, preview := cfg.makeUi()
	win := a.NewWindow("MarkDown")

	win.SetContent(container.NewHSplit(edit, preview))
	win.Resize(fyne.Size{
		Width:  800,
		Height: 500,
	})
	win.CenterOnScreen()
	win.ShowAndRun()
}
