package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (i Image) At(x, y int) color.Color {
	if (x>255) {
		x = 255 - x
	}
	if (y>255) {
		y = 255 - y
	}
	return color.RGBA{255, uint8(x), uint8(y), 255}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 512, 512)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

