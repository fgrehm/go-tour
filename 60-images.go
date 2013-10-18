package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width  int
	Height int
	Color  int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(i.Color + x), uint8(i.Color + y), 255, 255}
}

func main() {
	m := Image{150, 250, 250}
	pic.ShowImage(m)
}
