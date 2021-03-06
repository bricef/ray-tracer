package canvas

import (
	"fmt"
	"image"
	imageColor "image/color"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/bricef/ray-tracer/pkg/color"
	"github.com/bricef/ray-tracer/pkg/utils"
)

type Pixel struct {
	X int
	Y int
}

type ImageCanvas struct {
	width  int
	height int
	pixels [][]color.Color
}

type Canvas interface {
	Width() int
	Height() int
	Set(int, int, color.Color) error
	Get(x, y int) (color.Color, error)
	Pixels() *PixelIterator
	PixelChannel() <-chan Pixel
}

type PixelIterator struct {
	Canvas Canvas
	Cx     int
	Cy     int
}

func (i *PixelIterator) Get() (int, int) {
	x, y := i.Cx, i.Cy
	if i.Cx >= i.Canvas.Width()-1 {
		i.Cx = 0
		i.Cy += 1
	} else {
		i.Cx += 1
	}
	return x, y
}

func (i *PixelIterator) More() bool {
	more := i.Cx <= i.Canvas.Width()-1 && i.Cy <= i.Canvas.Height()-1
	return more
}

func (c *ImageCanvas) Pixels() *PixelIterator {
	return &PixelIterator{c, 0, 0}
}

func (c *ImageCanvas) PixelChannel() <-chan Pixel {
	out := make(chan Pixel)
	go func() {
		for x := 0; x < c.width; x++ {
			for y := 0; y < c.height; y++ {
				out <- Pixel{X: x, Y: y}
			}
		}
		close(out)
	}()
	return out
}

func (c *ImageCanvas) Height() int {
	return c.height
}

func (c *ImageCanvas) Width() int {
	return c.width
}

func NewImageCanvas(width, height int) *ImageCanvas {
	pixels := make([][]color.Color, width)
	for i := range pixels {
		pixels[i] = make([]color.Color, height)
	}
	return &ImageCanvas{
		width,
		height,
		pixels,
	}

}

func (c *ImageCanvas) Set(x int, y int, value color.Color) error {
	// fmt.Printf("Pixel %v,%v: %v\n", x, y, value)
	if x >= c.width || y >= c.height {
		return fmt.Errorf("out of bounds. Pixel %v,%v doesn't exist on canvas sized %v,%v", x, y, c.width, c.height)
	}
	c.pixels[x][y] = value
	// fmt.Printf("[%v,%v]=%v\n", x, y, value)
	return nil
}

func (c *ImageCanvas) Get(x, y int) (color.Color, error) {
	if x >= c.width || y >= c.height {
		return color.Color{}, fmt.Errorf("out of bounds. Pixel %v,%v doesn't exist on canvas sized %v,%v", x, y, c.width, c.height)
	}
	return c.pixels[x][y], nil
}

func (c *ImageCanvas) Image() image.Image {
	img := image.NewNRGBA64(image.Rect(0, 0, c.width, c.height))
	for x, column := range c.pixels {
		for y, pixel := range column {
			pixel = pixel.Cutoff()
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

func (c *ImageCanvas) WritePNG(filename string) {
	utils.EnsureDir(filepath.Dir(filename))
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(f, c.Image())
	if err != nil {
		log.Fatal(err)
	}
}
