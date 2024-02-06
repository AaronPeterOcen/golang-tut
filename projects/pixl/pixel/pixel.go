package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"zerotomastery.io-golang/src/projects/pixl/apptype"
	"zerotomastery.io-golang/src/projects/pixl/swatch"
	"zerotomastery.io-golang/src/projects/pixl/ui"
)

func main()  {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixl")

	state := apptype.State {
		BrushColor: color.White,
		SwatchSelected: 0,
	}

	appInit := ui.AppInit {
		PixlWindow: pixelWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}
	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}