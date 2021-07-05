package main

import (
	"math"

	"github.com/bricef/ray-tracer/canvas"
	"github.com/bricef/ray-tracer/color"
	. "github.com/bricef/ray-tracer/raytracer"
)

func main() {

	c := canvas.New(100, 100)

	pos := Point(50, 50, 0)

	for hour := 1; hour <= 12; hour++ {
		angle := (float64(hour) / 12.0) * (2.0 * math.Pi)
		dot := Transform().RotateZ(angle).Apply(Point(40, 0, 0)).Add(pos)
		c.Set(int(dot.X), int(dot.Y), color.New(1, 1, 1))

	}

	c.WritePNG("test.png")

}
