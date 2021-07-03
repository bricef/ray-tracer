package color

import (
	utils "github.com/bricef/ray-tracer/utils"
)

type Color struct {
	R float64
	G float64
	B float64
}

func New(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c Color) Sub(o Color) Color {
	return Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c Color) Scale(s float64) Color {
	return Color{c.R * s, c.G * s, c.B * s}
}

func (c Color) Mult(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c Color) Equal(o Color) bool {
	return utils.AlmostEqual(c.R, o.R) && utils.AlmostEqual(c.G, o.G) && utils.AlmostEqual(c.B, o.B)
}
