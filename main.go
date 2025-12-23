package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello Fyne")

	myWindow.SetContent(widget.NewLabel("Halo dari Go + Fyne!"))
	myWindow.ShowAndRun()
}
