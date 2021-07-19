package main

import (
	"math"

	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/color"
	m "github.com/bricef/ray-tracer/pkg/math"
)

func main() {

	c := canvas.NewImageCanvas(100, 100)

	pos := m.NewPoint(50, 50, 0)

	for hour := 1; hour <= 12; hour++ {
		angle := (float64(hour) / 12.0) * (2.0 * math.Pi)
		dot := m.NewTransform().RotateZ(angle).Apply(m.NewPoint(40, 0, 0)).Add(pos)
		c.Set(int(dot.X()), int(dot.Y()), color.New(1, 1, 1))

	}

	c.WritePNG("output/chapter4.png")

}
