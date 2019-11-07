package main

import (
	"golang.org/x/tour/pic"
    "fmt"
    "image"
    "image/color"
)

type Image struct{}

func (i Image) At(x, y int) color.Color {
    return color.RGBA{byte(x),byte(y),255,255}
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 100, 100)
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func main() {
    m := image.NewRGBA(image.Rect(0,0,100,100))
    fmt.Println(m.Bounds())
    fmt.Println(m.At(100, 100).RGBA())

    image := Image{}
    pic.ShowImage(image)
}