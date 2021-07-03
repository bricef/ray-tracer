package canvas

import (
	"fmt"
	"image"
	imageColor "image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/bricef/ray-tracer/color"
)

type Canvas struct {
	Width  int
	Height int
	pixels [][]color.Color
}

func NewCanvas(width, height int) Canvas {
	pixels := make([][]color.Color, width)
	for i := range pixels {
		pixels[i] = make([]color.Color, height)
	}
	return Canvas{
		width,
		height,
		pixels,
	}

}

func (c Canvas) Set(x int, y int, value color.Color) (Canvas, error) {
	if x >= c.Width || y >= c.Height {
		return Canvas{}, fmt.Errorf("out of bounds. Pixel %v,%v doesn't exist on canvas sized %v,%v", x, y, c.Width, c.Height)
	}
	c.pixels[x][y] = value

	return c, nil
}

func (c Canvas) Get(x, y int) (color.Color, error) {
	if x >= c.Width || y >= c.Height {
		return color.Color{}, fmt.Errorf("out of bounds. Pixel %v,%v doesn't exist on canvas sized %v,%v", x, y, c.Width, c.Height)
	}
	return c.pixels[x][y], nil
}

func (c Canvas) Image() image.Image {
	img := image.NewNRGBA64(image.Rect(0, 0, c.Width-1, c.Height-1))
	for x, column := range c.pixels {
		for y, pixel := range column {
			img.Set(x, y, imageColor.RGBA64{
				uint16(pixel.R * math.MaxUint16),
				uint16(pixel.G * math.MaxUint16),
				uint16(pixel.B * math.MaxUint16),
				math.MaxUint16,
			})
		}

	}
	return img
}

func (c Canvas) WritePNG(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(f, c.Image())
	if err != nil {
		log.Fatal(err)
	}
}
