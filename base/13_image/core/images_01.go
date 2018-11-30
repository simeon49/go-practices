package core

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{byte(x) * 10, byte(y) * 10, 255, 255}
}

func Images01() {
	m := Image{}
	pic.ShowImage(m)
}
