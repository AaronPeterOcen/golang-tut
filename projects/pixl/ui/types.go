package ui

import (
	"fyne.io/fyne/v2"
	"zerotomastery.io-golang/src/projects/pixl/apptype"
	"zerotomastery.io-golang/src/projects/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State 		*apptype.State
	Swatches	[]*swatch.Swatch
}