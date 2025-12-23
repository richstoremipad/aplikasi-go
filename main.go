package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Aplikasi Saya")

	myWindow.SetContent(widget.NewLabel("Halo! Ini aplikasi berukuran lebih kecil."))

	myWindow.ShowAndRun()
}
