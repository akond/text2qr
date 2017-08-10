package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"
	"github.com/skip2/go-qrcode"
	"log"
	"image"
	"os"
)

func appMain(driver gxui.Driver) {
	if len(os.Args) != 2 {
		log.Fatalln("Exactly one parameter is expected")
	}
	text := os.Args [1]
	size := 256
	a, err := qrcode.New(text, qrcode.Medium)

	if err != nil {
		log.Fatalln("Cannot create qr code")
	}
	b := a.Image(size)
	m := image.NewRGBA(image.Rect(0, 0, size, size))
	for row := 0; row < b.Bounds().Max.Y; row++ {
		for col := 0; col < b.Bounds().Max.X; col++ {
			m.Set(col, row, b.At(col, row))
		}
	}

	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(size, size, "QR code")
	texture := driver.CreateTexture(m, 1.0)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
