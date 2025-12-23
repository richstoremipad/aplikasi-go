package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
				case lifecycle.CrossOff:
					glctx = nil
				}
			case size.Event:
				// Menangani perubahan ukuran layar
			case paint.Event:
				if glctx == nil || e.External {
					continue
				}
				// Mewarnai layar menjadi biru (RGB)
				glctx.ClearColor(0, 0, 1, 1)
				glctx.Clear(gl.COLOR_BUFFER_BIT)
				a.Publish()
				a.Send(paint.Event{}) // Terus menggambar
			}
		}
	})
}
