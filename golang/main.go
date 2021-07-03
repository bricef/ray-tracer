package main

import (
	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/color"
)

func main() {
	c := canvas.NewCanvas(100, 100)
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			c.Set(x, y, color.NewColor(.1, .1, .1))
		}
	}
	c.Set(0, 0, color.NewColor(1, 1, 1))

	c.WritePNG("test.png")
}
