package color

import (
	"fmt"
	"math"

	"github.com/bricef/ray-tracer/pkg/utils"
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

var (
	Black = Color{0, 0, 0}
	White = Color{1, 1, 1}
	Red   = Color{1, 0, 0}
	Green = Color{0, 1, 0}
	Blue  = Color{0, 0, 1}
)

func (c Color) Cutoff() Color {
	return Color{
		math.Min(c.R, 1.0),
		math.Min(c.G, 1.0),
		math.Min(c.B, 1.0),
	}
}

func (c Color) String() string {
	return fmt.Sprintf("Color(%v,%v,%v)", c.R, c.G, c.B)
	// return fmt.Sprintf("Color(%v,%v,%v)", int(c.R*255), int(c.G*255), int(c.B*255))
}
