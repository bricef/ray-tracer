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

func Hex(v uint32) Color {
	r := uint16(v & 0xff0000 >> 16)
	g := uint16(v & 0x00ff00 >> 8)
	b := uint16(v & 0x0000ff)
	return Bytes(r, g, b)
}

func Bytes(rb, gb, bb uint16) Color {
	max := float64(255)
	r := float64(rb) / max
	g := float64(gb) / max
	b := float64(bb) / max
	return New(r, g, b)
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

func (c Color) EqualToTolerance(o Color, tolerance float64) bool {
	return utils.EqualToTolerance(c.R, o.R, tolerance) && utils.EqualToTolerance(c.G, o.G, tolerance) && utils.EqualToTolerance(c.B, o.B, tolerance)
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
