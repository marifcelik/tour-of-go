package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width  uint
	height uint
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(i.width), int(i.height))
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func images() {
	fmt.Println("\n-- images --")

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	// image exercise
	i := Image{80, 80}
	pic.ShowImage(i)
}
