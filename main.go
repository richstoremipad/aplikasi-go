package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Membuat aplikasi baru
	myApp := app.New()
	myWindow := myApp.NewWindow("Aplikasi Go Saya")

	// Komponen: Label (Teks)
	label := widget.NewLabel("Halo! Masukkan nama Anda di bawah:")

	// Komponen: Input Teks
	input := widget.NewEntry()
	input.SetPlaceHolder("Ketik nama di sini...")

	// Komponen: Tombol
	button := widget.NewButton("Klik Saya", func() {
		if input.Text == "" {
			label.SetText("Silakan masukkan nama terlebih dahulu!")
		} else {
			label.SetText("Selamat datang, " + input.Text + "!")
		}
	})

	// Menyusun komponen secara vertikal
	content := container.NewVBox(
		label,
		input,
		button,
	)

	// Menampilkan konten dan menjalankan aplikasi
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
