package main

import (
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
)

func main() {
	c := canvas.NewImageCanvas(100, 100)
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			c.Set(x, y, color.New(.1, .1, .1))
		}
	}
	c.Set(0, 0, color.New(1, 1, 1))

	c.WritePNG("test.png")
}
